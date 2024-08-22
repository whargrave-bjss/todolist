package main

import "fmt"

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

	fmt.Println("Welcome to Task Manager! Please choose what you want to do:")
	fmt.Println("1. List tasks, 2. Add task, 3. Complete task, 4. Delete task please enter the corresponding number:")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1: {
		ListTasks(mockTasks)
	}
	case 2: {
		var newTask string
		fmt.Println("Enter the task you want to add:")
		fmt.Scanln(&newTask)
		mockTasks = AddTask(mockTasks, newTask)
	 }
	case 3: {
		var taskToComplete int
		fmt.Println("Enter the number of the task you want to mark as completed:")
		fmt.Scanln(&taskToComplete)
		mockTasks = CompleteTask(mockTasks, taskToComplete)
	}
	case 4: {
		var taskToDelete int
		fmt.Println("Enter the number of the task you want to delete:")
		fmt.Scanln(&taskToDelete)	
		mockTasks = DeleteTask(mockTasks, taskToDelete)
	}
}
}