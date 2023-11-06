# Task: Creating Parallel Goroutines with Channels

Write a program in Golang that creates 3 goroutines to perform the following tasks in parallel:

Goroutine 1: Generates a random number between 1 and 10 and sends it to channel ch1.

Goroutine 2: Reads a number from ch1, multiplies it by 2, and sends the result to channel ch2.

Goroutine 3: Reads a number from ch2, prints it to the screen, and then terminates.

Use channels for data communication between goroutines. Ensure that the goroutines run concurrently and that the data is passed in the correct order.

This task will help you understand how to create goroutines and use channels for data exchange between them in parallel computations.