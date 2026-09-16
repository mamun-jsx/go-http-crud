# Go HTTP CRUD API

A simple RESTful API built with Go's standard library (`net/http`). This project demonstrates basic CRUD (Create, Read, Update, Delete) operations using an in-memory global slice as a mock database.

## Local Setup

1.  **Ensure you have Go installed:** Check with `go version`.
2.  **Navigate to the project directory:** Open your terminal where the `main.go` file is located.
3.  **Run the Server:**
    ```bash
    go run main.go
    ```
    The server will output `server is running on 4000` and start listening on `http://localhost:4000`.

## In-Memory Database Workflow (Global Variable)

This application uses a global slice of `User` structs to simulate a database:

```go
var users = []User{
	{Id: 1, Name: "Kamal", Age: 11, Email: "kk@gmail.com"},
	{Id: 2, Name: "Jamal", Age: 32, Email: "jamal@gmail.com"},
    // ...
}
```

**How it works:**
- **State:** The `users` variable lives in the application's RAM (memory). As long as the server is running, the state is maintained.
- **Mutations:** When you interact with endpoints that modify data (like `POST /create-user`, `PUT /user/{id}`, or `DELETE /user/{id}`), the handler functions directly modify this global `users` slice (e.g., appending a new user, updating a struct at a specific index, or slicing out a deleted user).
- **Volatility:** Because this is an in-memory data structure without a physical database connection (like PostgreSQL or MySQL), any changes made via the API will be **lost** when the server is stopped. Every time you run `go run main.go`, the application resets to the initial four hardcoded users.

## API Endpoints Reference

Here are the available endpoints you can interact with.

### 1. Health Check
- **Endpoint:** `GET /health`
- **Description:** Checks if the server is up and running.

### 2. Get All Users
- **Endpoint:** `GET /users`
- **Description:** Returns a JSON array of all users in the system.
- **Command:** 
  ```bash
  curl http://localhost:4000/users
  ```

### 3. Get Single User
- **Endpoint:** `GET /users/{id}`
- **Description:** Returns the user corresponding to the provided ID.
- **Command:** 
  ```bash
  curl http://localhost:4000/users/1
  ```

### 4. Create User
- **Endpoint:** `POST /create-user`
- **Description:** Creates a new user. The ID is automatically generated based on the length of the array (`len(users) + 1`).
- **Command:**
  ```bash
  curl -X POST http://localhost:4000/create-user \
    -H "Content-Type: application/json" \
    -d '{"name": "New User", "age": 25, "email": "new@example.com"}'
  ```

### 5. Update User
- **Endpoint:** `PUT /user/{id}`
- **Description:** Updates an existing user by ID. Iterates over the global slice, finds the user, and replaces the struct with the decoded JSON request body.
- **Command:**
  ```bash
  curl -X PUT http://localhost:4000/user/1 \
    -H "Content-Type: application/json" \
    -d '{"id": 1, "name": "Kamal Updated", "age": 12, "email": "kamal_updated@gmail.com"}'
  ```

### 6. Delete User
- **Endpoint:** `DELETE /user/{id}`
- **Description:** Deletes a user by their ID. Removes the user from the global slice.
- **Command:** 
  ```bash
  curl -X DELETE http://localhost:4000/user/1 -i
  ```
