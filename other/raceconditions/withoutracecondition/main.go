package main

import (
	"sync"
	"time"
)

type Counter struct {
	count int
	mw    sync.Mutex
}

func (c *Counter) Increment() {
	c.mw.Lock()
	defer c.mw.Unlock()
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
