package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool sync.Pool

	// Function for creating new objects
	pool.New = func() interface{} {
		return "new object"
	}

	// Get an object from the pool (or create a new one if empty)
	obj1 := pool.Get().(string)
	fmt.Println("First object from pool:", obj1)

	// Return an object to the pool
	pool.Put("reused object")

	// Get an object again
	obj2 := pool.Get().(string)
	fmt.Println("Second object from pool:", obj2)

	// Get an object again (pool is empty, so a new one is created)
	obj3 := pool.Get().(string)
	fmt.Println("Third object from pool:", obj3)
}
