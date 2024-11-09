package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	finished := make(chan struct{})
	go longOperation(ctx, finished)
	select {
	case <-finished:
		fmt.Println("Finished successfully.")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-interrupt:
		fmt.Println("Received an interrupt signal. Canceling the task...")
		cancel()
	}
	signal.Stop(interrupt)
	close(interrupt)
	fmt.Println("Program terminated.")
}

func longOperation(ctx context.Context, finished chan struct{}) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	case <-time.After(2 * time.Second): // imitating job
		finished <- struct{}{}
	}
}
