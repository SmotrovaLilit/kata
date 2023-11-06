package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic is recovered:", r)
		}
	}()

	panic("test panic recovery")
}
