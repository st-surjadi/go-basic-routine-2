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

func main() {
	// variable for bank balance
	var bankBalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Initial account balance: $%d.00.\n\n", bankBalance)

	// define weekly revenue
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Part-time Job", Amount: 50},
		{Source: "Investments", Amount: 100},
		{Source: "Gifts", Amount: 10},
	}

	wg.Add(len(incomes))

	// look through 52 weeks, print out how much is made and keep running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				bankBalance += income.Amount
				balance.Unlock()

				fmt.Printf("On week %d, you earned $%d.00 from %s.\n\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	// print out final balance
	fmt.Printf("Final bank balance: $%d.00.\n", bankBalance)
}
