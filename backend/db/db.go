package db

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"todolist/commands"
	
)

var db *sql.DB

func InitDB() () {
	var err error
	db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`
	if _, err := db.Exec(createUsersTable); err != nil {
        log.Fatal(err)
    }
	
	createTasksTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		UserID INTEGER NOT NULL,
		Item TEXT NOT NULL,
		DONE BOOLEAN NOT NULL DEFAULT FALSE,
		FOREIGN KEY (UserID) REFERENCES Users(ID) ON DELETE CASCADE
	);`

	if _, err := db.Exec(createTasksTable); err != nil {
		log.Fatal(err)
	}

	log.Println("Database initialized and tables created")

}

func SeedDB() {
	username := "BillyBob"
	password := "password"

	result, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		log.Fatal(err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	tasks, err := commands.LoadTasks()
	if err != nil {
		log.Fatal(err)
	}

	for _, task := range tasks {
		_, err = db.Exec("INSERT INTO tasks (UserID, Item, Done) VALUES (?, ?, ?)", userID, task.Item, task.Done)
		if err != nil {
			log.Fatal(err)
	}
}
}


func Close() {
    if err := db.Close(); err != nil {
        log.Fatal(err)
    }
}