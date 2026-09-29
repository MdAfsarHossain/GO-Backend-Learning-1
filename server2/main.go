package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go Backend!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Production API-তে JSON response পাঠালে header সেট করা উচিত:

	// http://localhost:5000/user?name=Afsar&age=26
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Println(`Name: `, name)
	fmt.Println("Age: ", age)

	response := map[string]string{
		"name":  "Afsar",
		"email": "afsar@example.com",
	}

	json.NewEncoder(w).Encode(response)
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func userHandlerStruct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := User{
		ID:    "1",
		Name:  "Afsar",
		Email: "afsar@example.com",
	}

	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/user-struct", userHandlerStruct)

	fmt.Println("Welcome to Server 2")
	fmt.Println("Server running on PORT: 5000")

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}

}
