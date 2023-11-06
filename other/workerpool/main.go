package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func main() {
	tasks := make(chan func(), 10)
	workersCount := runtime.NumCPU() - 1
	wg := &sync.WaitGroup{}
	wg.Add(workersCount)
	for i := 0; i < workersCount; i++ {
		go func(number int, pool chan func()) {
			for {
				select {
				case f, ok := <-pool:
					if !ok {
						fmt.Printf("%d: Worker done\n", number)
						wg.Done()

						pool = nil
						return
					}
					f()
				}
			}
		}(i, tasks)
	}
	go func() {
		testFunc := func() {
			fmt.Println("test function")
			math.Sqrt(6000000000000)
		}
		for i := 0; i < 100; i++ {
			tasks <- testFunc
		}
		close(tasks)
	}()
	wg.Wait()
}
