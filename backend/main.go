package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"todolist/utils"
)

func main() {
	utils.InitDB()
	defer utils.Close()
	done := make(chan struct{})
	commandChan := make(chan utils.Command)
	go utils.CommandHandler(commandChan, done)

	// Set up API routes
	http.HandleFunc("/api/tasks", utils.TasksHandler)
	http.HandleFunc("/api/add-task", utils.AuthMiddleware(utils.AddTaskHandler))
	http.HandleFunc("/api/delete-task/", utils.DeleteTaskHandler)
	http.HandleFunc("/api/update-task/", utils.UpdateTaskHandler)
	http.HandleFunc("/api/register", utils.RegisterHandler)
	http.HandleFunc("/api/login", utils.LoginHandler)


	go func() {
		log.Println("Starting server on :8080...")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
		close(done)
	}()

	
	for {
		fmt.Println("\nAvailable commands: 1: Server_Status, 2: TASKS 3: Add Task 4: Delete Task 5: Complete Task Q: Quit")
		fmt.Print("Enter command '1', '2', '3', '4', '5': ")
		var input string
		fmt.Scanln(&input)

		if input == "Q" {
			close(done)
			break
		}

		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		cmd := utils.Command{Type: parts[0], ResponseChan: make(chan string)}
		commandChan <- cmd
		response := <-cmd.ResponseChan
		fmt.Println(response)
	}

	<-done
	fmt.Println("Application shutting down...")
}