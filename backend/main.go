package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"todolist/db"
	"todolist/types"
)

func main() {
	db.InitDB()
	defer db.Close()
	done := make(chan struct{})
	commandChan := make(chan types.Command)
	go commandHandler(commandChan, done)

	// Set up API routes
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/add-task", addTaskHandler)
	http.HandleFunc("/api/delete-task/", deleteTaskHandler)
	http.HandleFunc("/api/update-task/", updateTaskHandler)


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

		cmd := types.Command{Type: parts[0], ResponseChan: make(chan string)}
		commandChan <- cmd
		response := <-cmd.ResponseChan
		fmt.Println(response)
	}

	<-done
	fmt.Println("Application shutting down...")
}