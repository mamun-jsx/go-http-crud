package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	// struct tag
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var users = []User{
	{Id: 1, Name: "Kamal", Age: 11, Email: "kk@gmail.com"},
	{Id: 2, Name: "Jamal", Age: 32, Email: "jamal@gmail.com"},
	{Id: 3, Name: "Ruhul", Age: 23, Email: "huhul@gmail.com"},
	{Id: 4, Name: "Habib", Age: 41, Email: "habib@gmail.com"},
}

func main() {
	mux := http.NewServeMux() // under the hood default mux will be used.

	mux.HandleFunc("/", rootHandler)
	fmt.Println("server is running on 4000")
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("POST /create-user", createUserHandler)
	mux.HandleFunc("GET /users", getUserHandler)

	// port listen
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		fmt.Println("server error ", err)
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to go server!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server is up and healthy")

}

func createUserHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "user created")

}

// ? Get all User
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	users, _ := json.Marshal(users)
	// ! send data to client side
	w.Write(users)
}
