package main

import (
	"database/sql"
	"fmt"
	"log"
)

func query(db *sql.DB) {
	var (
		id         int
		coursname  string
		price      float64
		instructor string
	)
	for {
		var inputID int
		fmt.Scan(&inputID)
		query := "SELECT id ,coursname, price, instructor FROME onlinecourse WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &coursname, &price, &instructor); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, coursname, price, instructor)
	}
}

func main() {
	db, err := sql.Open("mySQL", "root:krklde09@tcp(127.0.1:3306)/coursedb")
	if err != nil {
		fmt.Println("failed to connect")
	} else {
		fmt.Println("connect succesfully")
	}
	query(db)
}
