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




// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var msg string
// var wg sync.WaitGroup
// func updateMessage(newMsg string, mutex *sync.Mutex) {
// 	defer wg.Done()
// 	mutex.Lock()
// 	msg = newMsg
// 	mutex.Unlock()
// }

// func main(){
// 	msg = "Hello, World!"

// 	var mutex sync.Mutex
//  	wg.Add(2)
// 	go updateMessage("Hello, Go!", &mutex)
// 	go updateMessage("Hello, Gopher!", &mutex)
// 	wg.Wait()
// 	fmt.Println(msg)
// }