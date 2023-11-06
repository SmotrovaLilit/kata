package main

import (
	"time"
)

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func main() {
	c := &Counter{}

	go func(c *Counter) {
		for {
			c.Increment()
		}
	}(c)
	go func(c *Counter) {
		for {
			c.Increment()
		}
	}(c)

	time.Sleep(5 * time.Second)
}
