package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync"
)

// BackendServer represents a backend server with its URL
type BackendServer struct {
	URL  string
	Host string
}

// LoadBalancer represents the round-robin load balancer
type LoadBalancer struct {
	backendServers []*BackendServer
	currentIndex   int
	mutex          sync.Mutex
}

// NewLoadBalancer creates a new load balancer with the provided backend servers
func NewLoadBalancer(backendServers []*BackendServer) *LoadBalancer {
	return &LoadBalancer{
		backendServers: backendServers,
		currentIndex:   0,
		mutex:          sync.Mutex{},
	}
}

// NextBackendServer returns the next backend server in a round-robin fashion
func (lb *LoadBalancer) NextBackendServer() *BackendServer {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	server := lb.backendServers[lb.currentIndex]
	lb.currentIndex = (lb.currentIndex + 1) % len(lb.backendServers)

	return server
}

// ReverseProxyHandler handles incoming requests and forwards them to backend servers
func (lb *LoadBalancer) ReverseProxyHandler(w http.ResponseWriter, r *http.Request) {
	backend := lb.NextBackendServer()
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   backend.Host,
	})
	proxy.ServeHTTP(w, r)
}

func main() {
	nodes := os.Getenv("nodes")

	//Define the backend servers
	backendServers := []*BackendServer{}

	if nodes != "" {
		nodeList := strings.Split(nodes, ",")
		for _, node := range nodeList {
			backendServers = append(backendServers, &BackendServer{
				URL: "http://" + node, Host: node,
			})
		}
	} else {
		log.Fatal("No nodes specified")
	}

	// Create the load balancer
	loadBalancer := NewLoadBalancer(backendServers)

	// Set up the reverse proxy handler
	http.HandleFunc("/", loadBalancer.ReverseProxyHandler)

	// Start the load balancer server
	log.Println("Starting load balancer server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Load balancer server error:", err)
	}
}
