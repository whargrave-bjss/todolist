package main

import (
	"fmt"
)

var mockTasks = []Task{
	{	
		ID: 1,
		Item: "Buy groceries",
		Done: false,
	},
	{
		ID: 2,
		Item: "Finish Go project",
		Done: false,
	},
	{
		ID: 3,
		Item: "Call mom",
		Done: true,
	},
	{
		ID: 4,
		Item: "Go for a run",
		Done: false,
	},
	{
		ID: 5,
		Item: "Read a book",
		Done: true,
	},
	{
		ID: 6,
		Item: "Plan vacation",
		Done: false,
	},
	{
		ID: 7,
		Item: "Water plants",
		Done: true,
	},
	{
		ID: 8,
		Item: "Write blog post",
		Done: false,
	},
	{
		ID: 9,
		Item: "Learn new recipe",
		Done: false,
	},
	{
		ID: 10,
		Item: "Meditate",
		Done: true,
	},
}
func main() {
    tasks := mockTasks // Start with your mock tasks

    for {
        fmt.Println("\nWelcome to Task Manager! Please choose what you want to do:")
        fmt.Println("1. List tasks")
        fmt.Println("2. Add task")
        fmt.Println("3. Complete task")
        fmt.Println("4. Delete task")
        fmt.Println("5. Quit")
        fmt.Print("Please enter the corresponding number: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            ListTasks(tasks)
        case 2:
            var newTask string
            fmt.Print("Enter the task you want to add: ")
            fmt.Scanln(&newTask)
            tasks = AddTask(tasks, newTask)
            ListTasks(tasks)
        case 3:
            var taskToComplete int
            fmt.Print("Enter the number of the task you want to mark as completed: ")
            fmt.Scanln(&taskToComplete)
            tasks = CompleteTask(tasks, taskToComplete)
            ListTasks(tasks)
        case 4:
            var taskToDelete int
            fmt.Println("Enter the number of the task you want to delete:")
            ListTasks(tasks)
            fmt.Scanln(&taskToDelete)
            tasks = DeleteTask(tasks, taskToDelete)
            ListTasks(tasks)
        case 5:
            fmt.Println("Thank you for using Task Manager. Goodbye!")
            return // This will exit the program
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}