package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got", i, "from channel")

		// sleep for a second
		time.Sleep(time.Second)
	}
}

func main() {
	// buffered channel
	ch := make(chan int, 10)
	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("Sending", i, "to channel")
		ch <- i
		fmt.Println("Sent", i, "to channel")
	}

	fmt.Println("Done sending to channel")
	close(ch)
}
