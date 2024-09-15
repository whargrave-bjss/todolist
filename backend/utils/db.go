package utils

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() () {
	var err error
	DB, err = sql.Open("sqlite3", "./tasks.db")
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
	if _, err := DB.Exec(createUsersTable); err != nil {
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

	if _, err := DB.Exec(createTasksTable); err != nil {
		log.Fatal(err)
	}

	SetDB(DB)

	log.Println("Database initialized and tables created")

}

func SeedDB() {

	_, err := DB.Exec("DROP TABLE IF EXISTS tasks;")
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec("DROP TABLE IF EXISTS users;")
	if err != nil {
		log.Fatal(err)
	}

	
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	if _, err := DB.Exec(createUsersTable); err != nil {
		log.Fatal(err)
	}

	createTasksTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		UserID INTEGER NOT NULL,
		Item TEXT NOT NULL,
		DONE BOOLEAN NOT NULL DEFAULT FALSE,
		CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		 FOREIGN KEY (UserID) REFERENCES users(ID) ON DELETE CASCADE
	);`
	if _, err := DB.Exec(createTasksTable); err != nil {
		log.Fatal(err)
	}

	username := "BillyBob"
	password := "password" 
	result, err := DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		log.Fatal(err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}


	taskItem := "Water Plants"
	_, err = DB.Exec("INSERT INTO tasks (UserID, Item, Done) VALUES (?, ?, ?)", userID, taskItem, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database seeded with test data.")
}

func SetDB(database *sql.DB) {
	DB = database
}

func Close() {
    if err := DB.Close(); err != nil {
        log.Fatal(err)
    }
}