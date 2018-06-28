package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type PageVariables struct {
	Date string
	Time string
}

func main() {
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	HomePageVars := PageVariables{

		//Jan 2 15:04:05 2006 MST
		//1   2 3  4  6  6    7
		Date: now.Format("02 Jan 2006 MST"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("homepage.html")

	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
