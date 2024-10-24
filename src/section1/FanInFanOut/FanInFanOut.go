package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

// Fan-in and fan-out are two patterns that are used to distribute work amongst a number of goroutines.
// Fan-in is the process of combining multiple input channels into a single channel
// and the message from these channels are consumed by a single goroutine. Think funnel.

// Fan-out is the process of distributing work across multiple channels.
// This is the opposite of fan-in. Think splitting the work across multiple channels.
// looks like a tree branching out.

// A bidirectional channel can send and receive messages e.g. ch chan int
// A send-only channel can only send messages e.g. ch chan<- int
// A receive-only channel can only receive messages e.g. ch <-chan int
func producerFanIn(ch chan int, name string) {
	for {
		sleep()
		n := rand.Intn(100)
		// send the message
		fmt.Printf("Channel %s -> %d\n", name, n)
		ch <- n
	}
}

func producerFanOut(ch chan int) {
	for {
		sleep()
		n := rand.Intn(100)
		// send the message
		fmt.Println(" ->", n)
		ch <- n
	}
}

func consumerFanIn(ch <-chan int) {
	for n := range ch {
		fmt.Printf("Received %d\n", n)
	}
}

func consumerFanOut(ch <-chan int, name string) {
	for n := range ch {
		fmt.Printf("%s Received %d \n", name, n)
	}
}

func fanIn(chA, chB <-chan int, chC chan<- int) {
	for {
		select {
		case n := <-chA:
			chC <- n
		case n := <-chB:
			chC <- n
		}
	}
}

func fanOut(chA <- chan int, chB, chC chan <- int) {
	for n := range chA {
		if n < 50{
			chB <- n
		} else {
			chC <- n
		}
	}
}

func sleep() {
	n := rand.Intn(1000)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func main() {
	// chA := make(chan int)
	// chB := make(chan int)
	// chC := make(chan int)

	// go producerFanIn(chA, "A")
	// go producerFanIn(chB, "B")
	// go consumerFanIn(chC)
	// fanIn(chA, chB, chC)

	// fan-out
	chD := make(chan int)
	chE := make(chan int)
	chF := make(chan int)
	go producerFanOut(chD)
	go consumerFanOut(chE, "E")
	go consumerFanOut(chF, "F")
	fanOut(chD, chE, chF)
}
