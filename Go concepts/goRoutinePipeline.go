// Pipeline1 demonstrates an infinite 3-stage pipeline.
package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	delay := 100 * time.Millisecond

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(delay)
		}
	}()

	// Squarer
	// go func() {
	// 	for {
	// 		x := <-naturals
	// 		squares <- x * x
	// 	}
	// }()
	go func() {
		for x := range naturals {
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
