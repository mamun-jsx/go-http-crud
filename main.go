package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	fmt.Println("server is running on 4000")
	
	// port listen 
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("server error ", err)
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to go server!")
}
