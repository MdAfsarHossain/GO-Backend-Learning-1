package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go Backend API")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	user := User{
		ID:    "1",
		Name:  "Afsar",
		Email: "afsar@example.com",
	}

	json.NewEncoder(w).Encode(user)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var input CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(
			w,
			"Invalid JSON",
			http.StatusBadRequest,
		)
		return
	}

	fmt.Println("Name:", input.Name)
	fmt.Println("Email:", input.Email)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{
		"message": "User created successfully",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/user", getUserHandler)
	http.HandleFunc("/users", createUserHandler)

	fmt.Println("Server running on :5000")

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
