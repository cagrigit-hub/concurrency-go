package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"bravo",
		"charlie",
		"delta",
		"echo",
		"foxtrot",
		"golf",
		"hotel",
		"india",
		"juliet",
		"kilo",
	}
	wg.Add(len(words))
	for i, x := range words {
		go printSomething(fmt.Sprintf("This is the %dth word: %s", i, x), &wg)
	}
	wg.Wait()
	wg.Add(1)
	printSomething("This is a second thing to be printed!", &wg)
}
