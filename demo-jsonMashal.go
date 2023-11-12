package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Nattty tyty", "656565465", "hfewkhi@gmail.com"})
	fmt.Println(string(data))
}
