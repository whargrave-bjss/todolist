package utils

import "time"
type Task struct {
	ID int  `json:"ID"`
	UserId int `json:"UserID"`
	Item string `json:"Item"`
	Done bool `json:"Done"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type TaskList struct {
	Tasks []Task `json:"Tasks"`
}

type Command struct {
	Type string
	Args string
	ResponseChan chan string
}

type User struct {
	ID int
	Username string
	Password string
	CreatedAt time.Time
}

