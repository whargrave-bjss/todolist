package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
	
)


const filename = "tasks.json"

func LoadTasks() ([]Task, error) {

	var tasks []Task

	
	rows, err := DB.Query("SELECT ID, UserID, Item, Done FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("error querying tasks: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.UserId, &task.Item, &task.Done)
		if err != nil {
			return nil, fmt.Errorf("error scanning task: %v", err)
		}
		tasks = append(tasks, task)
	}
	
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error unmarshalling data: %v", err)
	}
	return tasks, nil
} 

func SaveTasks(tasks []Task) error {
	tasks = ResetIDs(tasks)
	taskList := TaskList{Tasks: tasks}
	data, err := json.Marshal(taskList)
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil { 
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}

func ResetIDs(tasks []Task) []Task {
	for i := range tasks {
		tasks[i].ID = i + 1
	}
	return tasks
}
func AddTask(item string) []Task {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
	}
	newTask := Task{Item: item, ID: len(tasks) + 1, Done: false}
	tasks = append(tasks, newTask)
	err = SaveTasks(tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
	}
	return ResetIDs(tasks)
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

func CompleteTask(itemToComplete int)  []Task {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
	}
	for i, task := range tasks {
		if task.ID == itemToComplete { 
			tasks[i].Done = true
		} else {
			fmt.Println("Task not found")
		}
		err = SaveTasks(tasks)
		if err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
		}		
	}
	return ResetIDs(tasks)}


func DeleteTask(itemToDelete int) []Task {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
	}
	for i, task := range tasks {
		if task.ID == itemToDelete {
			if !task.Done {
				fmt.Printf((" %s is not completed. Are you sure you want to delete it? it\n"), task.Item)
				var confirm string
				fmt.Scanln(&confirm)
				if confirm != "yes" { 
					fmt.Println("Task not deleted")
					return tasks
				} 
			} else {
				fmt.Printf("%s has been deleted\n", task.Item)
				tasks = append(tasks[:i], tasks[i+1:]...)
			}
		} 
	}
	return ResetIDs(tasks)
}

func DeleteAllCompleteTasks(tasks []Task) []Task {
	for i, task := range tasks {
		if task.Done {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func CompletedCount(tasks []Task) int{
	count := 0
	for _, task := range tasks {
		if task.Done {
			count++
		}
	}
	return count
}

var startTune time.Time

func init() {
	startTune = time.Now()
}


func GetServerStatus() string {
	uptime := time.Since(startTune)
	return fmt.Sprintf("Server uptime: %v", uptime)
}

func GetAllTasks() string{
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
	}
	var result strings.Builder
	for _, task := range tasks { 
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Fprintf(&result, "%d %s - %s\n", task.ID, task.Item, status)
	}
	return result.String()
}

