# Task: Implement a Worker Pool

Create a worker pool to process tasks concurrently using goroutines. 
Develop a pool of goroutines capable of executing task functions from a queue.

Requirements:

Create a worker pool that can execute multiple tasks concurrently.
Implement a task queue where tasks are added for processing.
The worker pool should consist of a fixed number of goroutines.
Tasks can be any arbitrary function that takes no arguments and returns no values (i.e., func()) for simplicity.
Demonstrate the ability to add tasks to the pool, which are then processed by the available workers.
Ensure proper synchronization to handle the concurrent execution of tasks.
Implement a graceful shutdown mechanism to stop the worker pool when all tasks are processed.
You can use Go's goroutines and channels to implement the worker pool. The worker pool should efficiently utilize available CPU cores for concurrent task execution.

