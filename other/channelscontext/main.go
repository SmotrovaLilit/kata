package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type Result struct {
	Number int
	Status int
	URL    string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("handler started")
		defer log.Print("handler ended")
		result := make(chan string)
		errors := make(chan error)
		ctx := r.Context()
		ctx, cancel := context.WithCancel(ctx)
		go func(ctx context.Context, cancel context.CancelFunc) {
			defer close(result)
			defer close(errors)
			wg := sync.WaitGroup{}
			for i := 0; i < 5; i++ {
				wg.Add(1)
				go func(ctx context.Context, cancel context.CancelFunc, i int) {
					defer wg.Done()
					res, err := executeRequest(ctx, fmt.Sprintf("%d", i))
					if err != nil {
						errors <- err
						cancel()
						return
					}
					result <- res
				}(ctx, cancel, i)
			}
			wg.Wait()
		}(ctx, cancel)
		for {
			select {
			case v, ok := <-result:
				if !ok {
					return
				}
				fmt.Fprintln(w, "Hello world! status code is ", v)
				log.Print("getting result: ", v)
			case err, ok := <-errors:
				if !ok {
					return
				}
				log.Print(err)
			case <-ctx.Done():
				log.Print(ctx.Err())
				return
			}
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func executeRequest(ctx context.Context, number string) (string, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://127.0.0.1:8081/?number="+number, nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
