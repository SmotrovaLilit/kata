package main

import (
	"database/sql"
	"log"
)

type userRepository struct {
	db *sql.DB
}

func newUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) createUser(name string, age int) error {
	insertDataSQL := "INSERT INTO users (name, age) VALUES (?, ?)"

	_, err := r.db.Exec(insertDataSQL, name, age)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) getUsers() ([]user, error) {
	queryDataSQL := "SELECT name, age FROM users"

	rows, err := r.db.Query(queryDataSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	users := make([]user, 0)
	for rows.Next() {
		var name string
		var age int
		if err := rows.Scan(&name, &age); err != nil {
			return nil, err
		}
		users = append(users, user{name: name, age: age})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
