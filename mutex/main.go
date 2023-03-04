package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup
func updateMessage(newMsg string) {
	defer wg.Done()
	msg = newMsg
}

func main(){
	msg = "Hello, World!"
 	wg.Add(2)
	go updateMessage("Hello, Go!")
	go updateMessage("Hello, Gopher!")
	wg.Wait()
	fmt.Println(msg)
}