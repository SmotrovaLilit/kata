# Task: Implement a Program Using Context

Create a program that utilizes the Go context package to cancel the execution of a task after a specified duration or upon user request.

Requirements:

The program should define a function that performs a time-consuming task (e.g., simulating file download or data processing). This function should accept a context.Context parameter.

The program should allow the user to initiate the task and set a timeout for its execution.

If the user doesn't cancel the task manually, it should automatically be canceled after the specified timeout.

Implement a signal (e.g., via keyboard input) that allows the user to cancel the task at any time.

Ensure that the program handles the cancellation gracefully and prints a message when the task is canceled, whether due to a timeout or user request.

You may use the context.WithTimeout function to create a context with a timeout.

Test your program to ensure that it behaves as expected when tasks are completed or canceled.

This task aims to demonstrate how to use the context package to manage long-running operations with the ability to cancel them based on user input or a predefined timeout.