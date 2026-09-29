package main

import (
	"fmt"
	"net/http"
)


func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go!")
}

func main() {
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":5000", nil)
}