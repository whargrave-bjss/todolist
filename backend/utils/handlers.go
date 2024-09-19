package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
    "time"

)


func CommandHandler(commandChan chan Command, done chan struct{}) {
    for {
        select {
        case <-done:
            return
        case cmd := <-commandChan:
            var response string
            switch cmd.Type {
            case "1":
                response = GetServerStatus()
            case "2":
                response = GetAllTasks()
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

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
    EnableCORS(&w, r)
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var task Task
	
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
    userID, ok := r.Context().Value("UserID").(int)
    if !ok {
        http.Error(w, "User not authenticated", http.StatusUnauthorized)
        return
    }

    task.UserID = userID



    result, err := DB.Exec("INSERT INTO tasks (UserID, Item, Done, CreatedAt) VALUES (?, ?, ?, ?)", task.UserID, task.Item, task.Done, time.Now())
	if err != nil {
		http.Error(w, "Error adding task", http.StatusInternalServerError)
		return
	}

	taskID, err := result.LastInsertId() 
	if err != nil {
		http.Error(w, "Error adding task", http.StatusInternalServerError)
		return
	}
	task.ID = int(taskID) 

	w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}


func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
    EnableCORS(&w, r)
    if r.Method == "OPTIONS" {
        return
    }

    if r.Method != "DELETE" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    taskIDStr := r.URL.Path[len("/api/delete-task/"):]
    taskId, err := strconv.Atoi(taskIDStr)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

   
    _, err = DB.Exec("DELETE FROM tasks WHERE ID = ?", taskId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	tasks, err := LoadTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}


func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
    EnableCORS(&w, r)
	if r.Method == "OPTIONS" {
		return
	}
	
	taskIDStr := r.URL.Path[len("/api/update-task/"):]
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var update struct {
		Done bool `json:"done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	
	_, err = DB.Exec("UPDATE tasks SET Done = ? WHERE ID = ?", update.Done, taskID)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    EnableCORS(&w, r)
    if r.Method == "OPTIONS" {
        return
    }
    
    var user User

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
    }

    hashedPassword, err := HashPassword(user.Password)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    _, err = DB.Exec("INSERT INTO users (Username, Password, CreatedAt) VALUES (?, ?, ?)", user.Username, hashedPassword, time.Now())
    if err != nil {
        http.Error(w, "Error registering user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    EnableCORS(&w, r)
    if r.Method == "OPTIONS" {
        return
    }

    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    var storedPassword string
    err := DB.QueryRow("SELECT Password, ID FROM users WHERE Username = ?", user.Username).Scan(&storedPassword, &user.ID)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    if !CheckPasswordHash(user.Password, storedPassword) {
        http.Error(w, "Invalid password", http.StatusUnauthorized)
        return
    }

    token, err := CreateToken(user.ID)
    if err != nil {
        http.Error(w, "Error creating token", http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "id":       user.ID,
        "Username": user.Username,
        "token": token,
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response) 
}


func EnableCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}
