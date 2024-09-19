package utils

import "time"
type Task struct {
	ID int  `json:"ID"`
	UserID int `json:"UserID"`
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
	ID int `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	CreatedAt time.Time `json:"CreatedAt"`
}

