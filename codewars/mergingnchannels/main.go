package mergingnchannels

import (
	"fmt"
	"sync"
)

func Merge(c ...chan string) <-chan string {
	result := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(len(c))
	for number, ch := range c {
		go func(ch chan string, number int) {
			defer wg.Done()
			for v := range ch {
				result <- v
			}
			fmt.Printf("channel %d closed\n", number)
		}(ch, number)
	}
	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}
