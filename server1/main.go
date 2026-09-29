package main

import (
	"fmt"
	"net/http"
)


func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Request
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)

	// Response
	fmt.Fprintln(w, "Hello, Go!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Route")
}

func main() {
    http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

    fmt.Println("Server running on :5000")

    err := http.ListenAndServe(":5000", nil)

    if err != nil {
        fmt.Println("Server error:", err)
    }
}

// go run server1/main.go