package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		// send a message to the channel
		s := <-ping
		pong <- fmt.Sprintf("Pong: %s", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press enter (Q to quit) :")
	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		// wait for a response
		response := <-pong
		fmt.Println("Response:", response)
	}
	fmt.Println("all done! closing channels...")

	// close the channels
	close(ping)
	close(pong)

}
