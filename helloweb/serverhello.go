package main

import (
	"fmt"
	"net/http"
	"github.com/thienquach/golang/helloweb/homepage"
	"github.com/thienquach/golang/helloweb/radio"
)

func main() {
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/", homepage.Load)
	http.HandleFunc("/radios", radio.LoadRadios)
	http.HandleFunc("/selected", radio.LoadAnswer)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "Hello world")
}
