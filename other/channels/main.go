package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://leetcode.com/problems/happy-number/",
		"https://leetcode.com/problems/contains-duplicate-ii/",
		"https://leetcode.com/problems/happy-number/",
		"https://leetcode.com/problems/contains-duplicate-ii/",
	}
	getRequests(http.DefaultClient, urls)
}

type Result struct {
	Number int
	Status int
	URL    string
}

type syncResultSlice struct {
	data []Result
	mu   sync.Mutex
}

func (s *syncResultSlice) add(r Result) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, r)
}

func getRequests(client *http.Client, urls []string) []Result {
	wg := sync.WaitGroup{}
	wg.Add(len(urls))
	quit := make(chan bool)
	result := syncResultSlice{}
	for n, u := range urls {
		go func(number int, url string) {
			time.Sleep(time.Duration(number * int(time.Second)))
			select {
			case _, _ = <-quit:
				wg.Done()
				fmt.Println(number+1, "request is canceled")
				quit <- true
				return
			default:
				resp, err := client.Get(url)
				wg.Done()
				if err != nil {
					fmt.Println(number+1, "an error occur")
					quit <- true
					return
				}
				defer resp.Body.Close()
				res := Result{
					Number: number + 1,
					Status: resp.StatusCode,
					URL:    url,
				}
				result.add(res)

				fmt.Println(number+1, "request is done")
			}
		}(n, u)
	}
	wg.Wait()
	return result.data
}
