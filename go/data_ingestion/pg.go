package main

import (
	"database/sql"
	"fmt"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "project01"
	password = "project01"
	dbname   = "project01"
)

func pg_read(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
