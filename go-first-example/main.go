package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	words := []string{
		"alpha", "beta", "delta", "gamma", "pi", "zeta", "eta", "theta", "epsilon", "omega",
	}

	wg.Add(len(words))

	for i, word := range words {
		go printString(fmt.Sprintf("%d: %s", i, word), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printString("This is the second thing to be printed!", &wg)

}

func printString(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
