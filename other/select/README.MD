# Task: Create Multiple Channels and Use select

Create multiple channels and utilize the select statement to choose a random channel for sending data. Apply this concept to build a simple random number generator using parallelism.

You should create, for example, three channels, and use select to randomly select one of them for sending data. 
Each channel should be associated with a goroutine that generates a random number and sends it to the selected channel. 
The main program should receive and print these random numbers.

Your solution should demonstrate how the select statement can be used to handle multiple channels concurrently and randomly choose between them for data transmission.