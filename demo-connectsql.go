package main

import (
	"database/sql"
	"fmt"
	"log"
)

func creatingTabledb(db *sql.DB) {
	query := `CREATA TABLE users (
		id INT AUTO_INCREMENT
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	);`

func Insert (db *sql.DB)  {
	var username string
	var password string
	fmt.Scan(&username)
	fmt.Scan(&password)
	createAt := time.Now()

	result, err := db.Exec(`INSERT INTO user (username, password, created_at) VALUES (?,?,?)` ,username, password, created_at)
	if err != nil{
		log.Fatal(err)
	}
	
	id, err := result.LastInsertId()
	fmt.Println(id)
}

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

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

func delete(db *sql.DB){
	var deleteid
	fmt.Scan(&deleteidid)
	_, var err := db.Exec(`DELETE FROM users WHERE id = ?`, deleteid)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("mySQL", "root:krklde09@tcp(127.0.1:3306)/coursedb")
	if err != nil {
		fmt.Println("failed to connect")
	} else {
		fmt.Println("connect succesfully")
	}
	creatingTable(db)
	Insert(db)
	delete((db)
}
