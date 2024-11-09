package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var ready = false                   // Resource state
	cond := sync.NewCond(&sync.Mutex{}) // Create a sync.Cond with a mutex

	var wg sync.WaitGroup

	// Function that waits for the resource to be ready
	waitForResource := func(id int) {
		defer wg.Done()
		cond.L.Lock() // Lock the mutex before checking the condition
		// We wrap the condition check with a for loop (for !ready { cond.Wait() }) to protect against spurious wakeups
		// (a rare but possible situation where a goroutine may wake up erroneously in cond.Wait()).
		// Without using a mutex and a loop, the goroutine could continue execution even if the condition has not been met.
		for !ready { // Wait while the resource is not ready
			fmt.Printf("Goroutine %d: Resource is not ready, waiting\n", id)
			cond.Wait() // Wait for the notification. The goroutine temporarily releases the mutex and blocks here
		}
		cond.L.Unlock() // Unlock the mutex after the condition is satisfied
		fmt.Printf("Goroutine %d: Resource is ready, starting work\n", id)
	}

	// Start several goroutines that are waiting for the resource to be ready
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go waitForResource(i)
	}

	// Simulate work to prepare the resource
	time.Sleep(2 * time.Second)
	cond.L.Lock()
	ready = true // Set the resource readiness flag
	cond.L.Unlock()
	cond.Broadcast() // Notify all goroutines that the resource is ready

	wg.Wait() // Wait for all goroutines to finish
}
