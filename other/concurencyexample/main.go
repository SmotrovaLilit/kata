package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func unpredictableFunc() int {
	n := rand.Intn(1000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n

}

func predictableFunc(ctx context.Context) int {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	res := make(chan int)

	go func() {
		defer close(res)
		tmp := unpredictableFunc()
		select {
		case <-ctx.Done():
			return
		case res <- tmp:
		}
	}()

	select {
	case <-ctx.Done():
		return 100
	case r := <-res:
		return r
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	n := predictableFunc(ctx)
	fmt.Println(n)
}
