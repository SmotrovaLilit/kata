package main

import (
	"fmt"
	"sync"
)

var once sync.Once

// The function to be executed only once
func initialize() {
	fmt.Println("Initialization happens only once")
}

func main() {
	var wg sync.WaitGroup

	// Launch multiple goroutines that will call the initialize function
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d is trying to call initialize\n", id)
			once.Do(initialize) // Ensure the function is called only once
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines have finished executing")
}
