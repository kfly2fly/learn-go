package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	// close the channel, but doesn't work when you dont know which goroutine will be last
	close(doneChan)

}

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go greet("Nice to meet you again!", done)
	go slowGreet("How... are... you..?", done)
	go greet("I hope you like the course", done)

	for range done {
		// fmt.Println(doneChan)
	}
}
