package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)





func main() {

	done := make(chan struct{})
	commandChan := make(chan Command)
	go commandHandler(commandChan, done)

	
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add-task", addTaskHandler)
	http.HandleFunc("/delete-task/", deleteTaskHandler)
	http.HandleFunc("/update-task/", updateTaskHandler)


	fs := customFileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	
	go func() {
		log.Println("Listening on :3000...")
		err := http.ListenAndServe(":3000", nil)
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

		cmd := Command{Type: parts[0], ResponseChan: make(chan string)}
		commandChan <- cmd
		response := <-cmd.ResponseChan
		fmt.Println(response)
	}

	<-done
	fmt.Println("Application shutting down...")
}