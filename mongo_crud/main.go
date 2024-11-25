package main

import (
	"WEB_SERVER/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("CRUD using MongoDB\n")

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateUser(w, r)
		case http.MethodGet:
			handler.GetAllUsers(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type to text/plain (default is text/html)
		w.Header().Set("Content-Type", "text/plain")

		// Write a response to the client
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetUserByID(w, r)
		case http.MethodPut:
			handler.UpdateUser(w, r)
		case http.MethodDelete:
			handler.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("Server running on port 7779")
	log.Fatal(http.ListenAndServe(":7779", nil))

}
