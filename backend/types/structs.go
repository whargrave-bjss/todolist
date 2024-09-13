package types

import "time"
type Task struct {
	ID int
	UserId int
	Item string
	Done bool
	CreatedAt time.Time
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

