package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "Hello world")
}
