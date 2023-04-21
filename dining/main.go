package main

import (
	"fmt"
	"sync"
	"time"
)

// The dining philosophers problem is well known in computer science circles.
// five philo. numbered from 0 through 4, live in a house where the
// table is laid for them; each phi. has their own place at the table.
// Their only difficulty is that the dish
// served is a very difficult kind of spaghetti that can only be eaten with
// two forks. There are five forks on the table, one between each pair of
// plates. Each phi. needs to use two forks to eat. After eating, they
// put down the forks and continue thinking. The problem is to design a
// discipline or behavior pattern for the phi. that solves the problem of
// sharing the forks.

// This is a simple impl. of Dijkstra's solution to the dining philosophers
// Philosophers" dilemma.

// Phi. is a struct which stores information about phi.
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// philosophers list of all phis.
var phiolosophers = []Philosopher{
	{"Plato", 0, 4},
	{"Socrates", 1, 0},
	{"Aristotle", 2, 1},
	{"Pascal", 3, 2},
	{"Locke", 4, 3},
}

// define some variables
var hunger = 3 // how many times does a person eat?
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	// print out a welcome message
	fmt.Println("Dining philosophers problem")
	fmt.Println("===========================")
	fmt.Println("The table is empty.")
	// start the meal
	dine()
	// print out finished message
	fmt.Println("The table is empty.")
}

func dine() {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second

	wg := &sync.WaitGroup{}
	wg.Add(len(phiolosophers))
	// loop through the phi. list

	seated := &sync.WaitGroup{}
	seated.Add(len(phiolosophers))

	// forks is a map of all 5 forks.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(phiolosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal.
	for i := 0; i < len(phiolosophers); i++ {
		// fire off a go routine for the current phi.
		go diningProblem(phiolosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philo at the table
	fmt.Printf(("%s is seated at the table.\n"), philosopher.name)
	seated.Done()
	seated.Wait()

	// eat three times
	for i := 0; i < hunger; i++ {
		// get a lock on both forks

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s has the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s has the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s has the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s has the right fork.\n", philosopher.name)
		}

		fmt.Printf("%s is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("%s put down the forks.\n", philosopher.name)

	}
	fmt.Println(philosopher.name, "is satisfied.")
	fmt.Println(philosopher.name, "is left the table.")

}
