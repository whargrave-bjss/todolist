package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
    "log"
)


func commandHandler(commandChan chan Command, done chan struct{}) {
    for {
        select {
        case <-done:
            return
        case cmd := <-commandChan:
            var response string
            switch cmd.Type {
            case "1":
                response = getServerStatus()
            case "2":
                response = getAllTasks()
            case "3":
                var newTask string
                fmt.Print("Enter the task you want to add: ")
                fmt.Scanln(&newTask)
                AddTask(newTask)
                response = fmt.Sprintf("%s has been added to the list of tasks", newTask)
            case "4":
                var taskToDelete int
                fmt.Println("Enter the number of the task you want to delete:")
                fmt.Scanln(&taskToDelete)
                DeleteTask(taskToDelete)
                response = "Task deleted"
            case "5":
                var taskToComplete int
                fmt.Print("Enter the number of the task you want to mark as completed: ")
                fmt.Scanln(&taskToComplete)
                CompleteTask(taskToComplete)
                response = "Task marked as completed"
            default:
                response = "Invalid command"
            }
            cmd.ResponseChan <- response
        }
    }
}

//handler functions
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
    enableCORS(&w, r)
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var task Task
	
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	} 
	tasks = append(tasks, task)

	err = saveTasks(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
    enableCORS(&w, r)
    if r.Method == "OPTIONS" {
        return
    }

    if r.Method != "DELETE" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Log the full request URL
    log.Printf("Delete request received for path: %s", r.URL.Path)

    // Extract the ID from the URL
    parts := strings.Split(r.URL.Path, "/")
    log.Printf("URL parts: %v", parts)

    if len(parts) < 4 {
        http.Error(w, "Invalid request path", http.StatusBadRequest)
        return
    }

    taskId, err := strconv.Atoi(parts[len(parts)-1])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }


    tasks, err := loadTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var taskDeleted bool
    for i, task := range tasks {
        if task.ID == taskId {
            tasks = append(tasks[:i], tasks[i+1:]...)
            taskDeleted = true
            break
        }
    }

    if !taskDeleted {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    err = saveTasks(tasks)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}


func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
    enableCORS(&w, r)
    if r.Method == "OPTIONS" {
        return
    }

    if r.Method != "PATCH" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 3 {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }
    taskID, err := strconv.Atoi(parts[len(parts)-1])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    var update struct {
        Done bool `json:"Done"`
    }
    if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    tasks, err := loadTasks()
    if err != nil {
        http.Error(w, "Failed to load tasks", http.StatusInternalServerError)
        return
    }
    var updatedTask *Task
    for i := range tasks {
        if tasks[i].ID == taskID {
            tasks[i].Done = update.Done
            updatedTask = &tasks[i]
            break
        }
    }
    if updatedTask == nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }
    if err := saveTasks(tasks); err != nil {
        http.Error(w, "Failed to save tasks", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedTask)
}

func enableCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}
