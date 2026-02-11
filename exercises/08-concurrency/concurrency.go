package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// TODO:
// Run 5 goroutines that print numbers 1–5
// Ensure main waits for them to finish

func count() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("counter is finished")

}

func main() {

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		count()
	}
	wg.Wait()
}
