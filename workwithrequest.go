package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Course struct {
	Id         int     `json: "id"`
	Name       string  `json: "name"`
	Price      float64 `json: "price"`
	Instructor string  `json: "instructor"`
}

var Courselist []Course

func init() {
	CourseJSON := `[
		{
		"id":1,
		"name":"Aim"
		"price":43,
		"instructor":"lad"
		},
		{
			"id":2,
			"name":"lim"
			"price":90,
			"instructor":"ladgra"
			},
			{
				"id":3,
				"name":"jkom"
				"price":12,
				"instructor":"ladgrabang"
				}
	]`
	err := json.Unmarshal([]byte(CourseJSON), &Courselist)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextId() int {
	highestID := -1
	for _, course := range Courselist {
		if highestID < course.Id {
			highestID = course.Id
		}
	}
	return highestID + 1
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(Courselist)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "appilcation/json")
		w.Write(courseJSON)
	case http.MethodPost:
		var newCourse Course
		Bodybyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(Bodybyte, newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newCourse.Id != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newCourse.Id = getNextId()
		Courselist = append(Courselist, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}
func main() {
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)
}
