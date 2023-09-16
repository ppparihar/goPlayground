package main

import (
	"crypto/md5"
	"fmt"
	"sort"
)

type ConsistentHashing struct {
	replicas   int
	circle     map[uint32]string
	sortedKeys []uint32
}

func NewConsistentHashing(replicas int) *ConsistentHashing {
	return &ConsistentHashing{
		replicas:   replicas,
		circle:     make(map[uint32]string),
		sortedKeys: []uint32{},
	}
}

func (ch *ConsistentHashing) AddNode(node string) {
	for i := 0; i < ch.replicas; i++ {
		replicaKey := ch.getReplicaKey(node, i)
		ch.circle[replicaKey] = node
		ch.sortedKeys = append(ch.sortedKeys, replicaKey)
	}

	sort.Slice(ch.sortedKeys, func(i, j int) bool {
		return ch.sortedKeys[i] < ch.sortedKeys[j]
	})
}

func (ch *ConsistentHashing) RemoveNode(node string) {
	for i := 0; i < ch.replicas; i++ {
		replicaKey := ch.getReplicaKey(node, i)
		delete(ch.circle, replicaKey)
		for j, key := range ch.sortedKeys {
			if key == replicaKey {
				ch.sortedKeys = append(ch.sortedKeys[:j], ch.sortedKeys[j+1:]...)
				break
			}
		}
	}
}

func (ch *ConsistentHashing) GetNode(key string) string {
	if len(ch.circle) == 0 {
		return ""
	}

	hashKey := ch.hashKey(key)
	for _, nodeKey := range ch.sortedKeys {
		if hashKey <= nodeKey {
			return ch.circle[nodeKey]
		}
	}

	// Wrap around to the first node if the key is larger than all node keys
	return ch.circle[ch.sortedKeys[0]]
}

func (ch *ConsistentHashing) getReplicaKey(node string, replicaIndex int) uint32 {
	replicaStr := fmt.Sprintf("%s-%d", node, replicaIndex)
	return ch.hashKey(replicaStr)
}

func (ch *ConsistentHashing) hashKey(key string) uint32 {
	hash := md5.Sum([]byte(key))
	return uint32(hash[0])<<24 | uint32(hash[1])<<16 | uint32(hash[2])<<8 | uint32(hash[3])
}

func main() {
	nodes := []string{"Node A", "Node B", "Node C"}
	hashRing := NewConsistentHashing(3)
	for _, node := range nodes {
		hashRing.AddNode(node)
	}

	fmt.Println("Initial Nodes:")
	for _, node := range nodes {
		fmt.Printf("Key: %s, Node: %s\n", node, hashRing.GetNode(node))
	}

	newNode := "Node D"
	hashRing.AddNode(newNode)

	fmt.Println("\nAfter Adding Node D:")
	for _, node := range append(nodes, newNode) {
		fmt.Printf("Key: %s, Node: %s\n", node, hashRing.GetNode(node))
	}

	removedNode := "Node B"
	hashRing.RemoveNode(removedNode)

	fmt.Println("\nAfter Removing Node B:")
	for _, node := range append(nodes, newNode) {
		fmt.Printf("Key: %s, Node: %s\n", node, hashRing.GetNode(node))
	}
}
