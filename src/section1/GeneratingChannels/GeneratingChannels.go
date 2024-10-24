package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Channels in Go are a powerful feature used for communication between goroutines.
They allow you to send and receive values between different goroutines,
enabling concurrent programming.
Channels can be thought of as pipes through which data can flow.
*/

func process2(ch chan int) {
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	ch <- n
}

func main() {
	ch := make(chan int)
	go process2(ch) // this process happens concurrently as the rest of the code
	fmt.Println("Waiting for the process to finish...")
	n := <-ch
	fmt.Printf("Process took %dms\n", n)
}
