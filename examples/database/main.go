package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db := initDB()
	defer db.Close()
	repository := newUserRepository(db)
	err := repository.createUser("John Doe", 30)
	if err != nil {
		panic(err)
	}

	users, err := repository.getUsers()
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable(db)
	return db
}

func createTable(db *sql.DB) {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            age INTEGER
        );
    `

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

type user struct {
	name string
	age  int
}
