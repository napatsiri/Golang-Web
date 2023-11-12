package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var Db *sql.DB

const coursePath = "courses"
const BasePath = "/api"

type Course struct {
	CourseID  int     `json: "courseid"`
	Cousename string  `json: "cousename"`
	Price     float64 `json: "price"`
	ImageURL  string  `json: "imageURL"`
}

func SetupDB() {
	var err error
	Db, err = sql.Open("mysql", "root:mflv[DB2022@tcp(127.0.0.1:3306)/coursedb")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Db)
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}

func getCourseList() ([]Course, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := DbConn.QueryContext(ctx, `SELECT
	courseid,
	coursename,
	price,
	image_url
	FROM test`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	courses := makr([]Course, 0)
	for results.Next(){
		var course Course
		results.Scan(&course.CourseID,
		&course.Cousename,
		&course.Price,
		&course.ImageURL)

		courses = append(courses, course)
	}
	return courses, nil
}

func handleCourses(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		courseList, err := getCourseList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(courseList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func corsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "appilcation/json")
		w.Header().Set("Access-Control-Allow-Method", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Header", "Accept, Content-Type, Content-Len")
	})
}

func SetupRoutes(apiBasePath string) {
	coursesHandler := http.HandlerFunc(handleCourses)

	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, coursePath), cosMiddleware(coursesHandler))
}

func main() {
	SetupDB()
	SetupRoutes(BasePath)
	log.Fatel(http.ListenAndServe(":5000", nil))
}
