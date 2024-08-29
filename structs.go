package main

type Task struct {
	ID int `json:"ID"`
	Item string `json:"Item"`
	Done bool `json:"Done"`
}

type TaskList struct {
	Tasks []Task `json:"Tasks"`
}

type Command struct {
	Type string
	Args string
	ResponseChan chan string
}