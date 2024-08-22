package main

import "fmt"



func AddTask(tasks []Task, item string) []Task {
	newTask := Task{Item: item, ID: len(tasks) + 1, Done: false}
	return append(tasks, newTask)
}

func ListTasks(tasks []Task) {
	for _, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d %s - %s\n", task.ID, task.Item, status)
	}
}

func CompleteTask(tasks []Task, itemToComplete int)  []Task {
	for i, task := range tasks {
		if task.ID == itemToComplete { 
			tasks[i].Done = true
		}
	}
	fmt.Println("Task not found")
	return tasks
}

func DeleteTask(tasks []Task, itemToDelete int) []Task {
	for i, task := range tasks {
		if task.ID == itemToDelete { 
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	fmt.Println("Task not found")
	return tasks
}




