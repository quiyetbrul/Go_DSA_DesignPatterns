package main

import (
	"fmt"
	"time"
)

func echoWorker(in, out chan int) {
	for {
		n := <-in
		time.Sleep(time.Duration(n) * time.Millisecond)
		out <- n
	}
}

func produce(ch chan<- int) {
	i := 0
	for {
		fmt.Println("-> send job", i)
		ch <- i
		i++
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)

	for i := 0; i < 4; i++ {
		go echoWorker(in, out)
		go produce(in)
	}

	for n := range out {
		fmt.Println("<- received job", n)
	}
}

