package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true // send a value to the channel to indicate completion  (arrow points in direction data should flow)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true // send a value to the channel to indicate completion  (arrow points in direction data should flow)
	close(doneChan)  // close channel when you know you are done - if you close it before its done its not going to support any more data
	// above close only works if you know which channel takes the longest
}

func main() { // go says run them as go routines concurrently
	//	dones := make([]chan bool, 4) // will transmit data between go routines - communication device
	done := make(chan bool)

	// variable will say whether a go routine is done or not

	//	dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)

	//	dones[1] = make(chan bool)
	go greet("How are you?", done)

	//	dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)

	//	dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)

	for range done {
		// fmt.Println(doneChan)
		// <-done
	}

	/*
		<-done // wait here until we receive a value from the done channel
		<-done // wait here until we receive a value from the done channel
		<-done // wait here until we receive a value from the done channel
		<-done // wait here until we receive a value from the done channel
	*/
	// if you use same channel for multiple go routines you have wait for them all
	// can also work with a slice of channels if you want to track each one individually
}

// when you run go functions the println will not work because the main function will exit before the go routines complete
