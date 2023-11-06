package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go generate(ch1, 1)
	go generate(ch2, 2)
	go generate(ch3, 3)

	for {
		select {
		case num := <-ch1:
			fmt.Printf("Received from Generator 1: %d\n", num)
		case num := <-ch2:
			fmt.Printf("Received from Generator 2: %d\n", num)
		case num := <-ch3:
			fmt.Printf("Received from Generator 3: %d\n", num)
		}
	}
}

func generate(ch chan int, id int) {
	for {
		num := rand.Intn(100)

		ch <- num

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		fmt.Printf("Generator %d sent: %d\n", id, num)
	}
}
