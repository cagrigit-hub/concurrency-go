package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup


type Income struct {
	Source string
	Amount int
}


func main(){
	// variable for bank balance
	var balance int
	var bMutex sync.Mutex
	// print out starting val
	fmt.Printf("Starting balance: %d.00 ", balance)
	fmt.Println()
	// define weekly revenue
	incomes := []Income{
		{"Salary", 1000},
		{"Side Hustle", 500},
		{"Dividends", 100},
		{"Interest", 50},
	}

	wg.Add(len(incomes))
	// loop through 52 weeks printing out the balance: keep a running total
	for i, income := range incomes {
		go func(i int , income Income){
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				bMutex.Lock()
				temp := balance
				temp += income.Amount
				balance = temp
				bMutex.Unlock()
				fmt.Printf("Week %d: %s: %d.00 \n", week, income.Source, balance)
			}
				
		}(i, income)
		
	}
	wg.Wait()
	// print out final balance
	fmt.Printf("Final balance: %d.00 ", balance)
	fmt.Println()
}