package main

import (
	"fmt"
)


func main() {
    tasks, err := loadTasks()
		if err != nil {
			fmt.Printf("Error loading tasks: %v\n", err)
			return
		} 

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
			if err := saveTasks(tasks); err != nil {
				fmt.Printf("Error saving tasks: %v\n", err)
			}
            ListTasks(tasks)
        case 3:
            var taskToComplete int
            fmt.Print("Enter the number of the task you want to mark as completed: ")
            fmt.Scanln(&taskToComplete)
            tasks = CompleteTask(tasks, taskToComplete)
			if err := saveTasks(tasks); err != nil { 
				fmt.Printf("Error saving tasks: %v\n", err)
			}
            ListTasks(tasks)
        case 4:
            var taskToDelete int
            fmt.Println("Enter the number of the task you want to delete:")
            ListTasks(tasks)
            fmt.Scanln(&taskToDelete)
            tasks = DeleteTask(tasks, taskToDelete)
			if err := saveTasks(tasks); err != nil {
				fmt.Printf("Error saving tasks: %v\n", err)
			}
            ListTasks(tasks)
        case 5:
            fmt.Println("Thank you for using Task Manager. Goodbye!")
            return 
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}