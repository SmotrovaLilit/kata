package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- rand.Intn(10) + 1
	}()
	go func(ch chan int) {
		v := <-ch
		ch2 <- v * 2
	}(ch1)
	go func(ch chan int) {
		fmt.Println(<-ch)
	}(ch2)

	time.Sleep(3 * time.Second)
}
