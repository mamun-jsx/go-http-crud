package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
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
	mux.HandleFunc("GET /users/{id}", getSingleUserHandler)
	mux.HandleFunc("PUT /user/{id}", updateUserHandler)
	mux.HandleFunc("DELETE /user/{id}", deleteUserHandler)

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

// create a user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser) //decode the user from json to struct
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(w, "Invalid request body")
		return
	}
	fmt.Println("user created", newUser)
	newUser.Id = len(users) + 1 // create id
	users = append(users, newUser)
	//* now send to client
	w.Header().Set("Content-Type", "Application/json") // set content type otherwise it will send default text type content
	w.WriteHeader(http.StatusCreated)                  // send code to client side
	json.NewEncoder(w).Encode(users)                   // send json response to client
}

// ? Get all User
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	// users, _ := json.Marshal(users)
	// ! send data to client side
	// w.Write(users)
	// json encoding
	encoder := json.NewEncoder(w)
	encoder.Encode(users)
	// encoding is more memory efficient
}

// ? get user by id

func getSingleUserHandler(w http.ResponseWriter, r *http.Request) {

	idParam := r.PathValue("id") // r.PathValue will provide the id from URL
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Invalid user id")
		return
	}
	for _, user := range users {
		if user.Id == id {
			// if id is matched then send it to client side
			w.Header().Set("Content-Type", "Application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	// send response status if user is not found
	w.WriteHeader(http.StatusNotFound)
}

// update user by Id
// Command to test: curl -X PUT http://localhost:4000/user/1 -H "Content-Type: application/json" -d '{"id": 1, "name": "Updated Name", "age": 20, "email": "updated@example.com"}'
// Description:
// 1. Extracts the 'id' from the URL path.
// 2. Decodes the JSON request body into a User struct, ensuring no unknown fields exist.
// 3. Iterates through the in-memory 'users' slice to find a matching user ID.
// 4. If found, updates the user at that index and returns the updated user as JSON.
// 5. If not found, returns a 404 Not Found status.
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "invalid user id")
		return
	}

	var updateUser User
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&updateUser); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body: " + err.Error()})
		return
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "request body must contain exactly one JSON object"})
		return
	}

	for idx, user := range users {
		if user.Id == id {
			updateUser.Id = id
			users[idx] = updateUser
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateUser)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "user not found")
}

// ? delete user by id
// Command to test: curl -X DELETE http://localhost:4000/user/1 -i
// Description:
// 1. Extracts the 'id' from the URL path.
// 2. Iterates through the in-memory 'users' slice to find the user with the matching ID.
// 3. If found, removes the user from the slice using slice slicing (users[:idx] and users[idx+1:]).
// 4. Returns a 204 No Content status on success (no response body).
// 5. If not found, returns a 404 Not Found error.
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	for idx, user := range users {
		if user.Id == id {
			// ! menual feature
			// users = append(users[:idx], users[idx+1:]...)
			// w.WriteHeader(http.StatusNoContent)
			// return

			// ! built in feature
			users = slices.Delete(users, idx, idx+1)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}
