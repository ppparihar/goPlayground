package main

import (
	"fmt"
	"time"
)

func main() {
	//goRoutineSample1()
	//goRoutineSample2()
	//goRoutineSample3()
	goRouteSample4()
}

func goRoutineSample1() {
	message := make(chan string)
	// Start a goroutine that sends "Hello" to the channel after 2 seconds
	go func() {
		time.Sleep(1 * time.Second)
		message <- "Hello"
	}()

	// Start another goroutine that sends "World" to the channel after 1 second
	go func() {
		time.Sleep(2 * time.Second)
		message <- "World"
	}()

	// Receive and print the messages from the channel
	fmt.Println(<-message) // World
	fmt.Println(<-message) // Hello

}
func goRoutineSample2() {
	number := make(chan int)
	go func() {
		number <- 10
	}()

	fmt.Println(<-number)
}

func goRoutineSample3() {
	number := make(chan int)
	go channelData(number)

	fmt.Println(<-number)
}
func channelData(number chan int) {
	number <- 100
}
func goRouteSample4() {
	message := make(chan string)
	go func() {
		message <- "Hello Duniya"
	}()

	fmt.Printf(<-message)
}
