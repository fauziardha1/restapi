package main

import (
	"database/sql"
	"fmt"
)

const (
	DB_USER     = "fauzi"
	DB_PASSWORD = "hello"
	DB_NAME     = "solid"
)

// db setup
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	return db
}
