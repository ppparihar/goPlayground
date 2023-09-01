package main

import (
	"fmt"
	"math/rand"
)

// fanIn merges two channels of integers into one
func fanIn(ch1, ch2 <-chan int, done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case x, ok := <-ch1:
				if !ok {
					ch1 = nil
				} else {
					out <- x
				}
			case y, ok := <-ch2:
				if !ok {
					ch2 = nil
				} else {
					out <- y
				}
			case <-done:
				return
			}
			if ch1 == nil && ch2 == nil {
				return
			}
		}
	}()
	return out
}

// sum receives integers from a channel and returns their sum
func sum(ch <-chan int) int {
	total := 0
	for x := range ch {
		total += x
	}
	return total
}

// main creates two channels that send random numbers and uses fanIn and sum to get the total
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- rand.Intn(100)
		}
		close(ch1)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- rand.Intn(100)
		}
		close(ch2)
	}()
	out := fanIn(ch1, ch2, done)
	fmt.Println(sum(out)) // prints the sum of 20 random numbers
	close(done)
}
