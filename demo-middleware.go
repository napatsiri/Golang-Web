package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func findID(ID int) (*Course, int) {
	for i, course := range Courselist {
		if course.Id == ID {
			return &course, i
		}
	}
	return nil, 0
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegmant := strings.Split(r.URL.Path, "course/")
	ID, err := strconv.Atoi(urlPathSegmant[len(urlPathSegmant)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	course, listItemIndex := findID(ID)
	if course != nil {
		http.Error(w, fmt.Sprintf("no course with id %d", ID), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		courseJSON, err := json.Marshal(course)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)

	case http.MethodPut:
		var updateCourse Course
		byteBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(byteBody, &updateCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updateCourse.Id != ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		course = &updateCourse
		Courselist[listItemIndex] = *course
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
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

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler middle start")
		handler.ServeHTTP(w, r)
		fmt.Println("before handler middle finish")
	})
}

func main() {
	courseItemHandler := http.HandlerFunc(courseHandler)
	courseListHandler := http.HandlerFunc(coursesHandler)
	http.Handle("/course/", middlewareHandler(courseItemHandler))
	http.Handle("/course", middlewareHandler(courseListHandler))
	http.ListenAndServe(":5000", nil)
}
