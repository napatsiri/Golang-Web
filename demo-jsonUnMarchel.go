package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID" :101, "EmployeeName" : "Sirisit", "Tel" : "9434902", "Email" : "jndkgnekw@gmaul.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e.ID)
}
