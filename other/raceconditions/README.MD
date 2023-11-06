# Race conditions

Write a program that creates goroutines which modify a shared counter. 
Try running the program without synchronization and observe a data race. 
Then, use mutexes to prevent the data race.