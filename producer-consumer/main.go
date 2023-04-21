package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func pizzeria(pizzaMaker *Producer) {

}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The pizzeria is open for business!")
	color.Cyan("==================================")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	// run the producer in the background
	go pizzeria(pizzaJob)
	// create and run consumer

	// print out the ending message
}
