package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // under the hood default mux will be used.

	mux.HandleFunc("/", rootHandler)
	fmt.Println("server is running on 4000")
	mux.HandleFunc("/health", healthHandler)
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