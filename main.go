package main

import (
	"fmt"
	"log"
	"net/http"
	"study-go/hello/Day3"
)

func main() {
	usersLoaded, err := Day3.LoadUsers()
	if err != nil {
		fmt.Println("Error loading users:", err)
		return
	}

	server := &Day3.Server{
		Users: usersLoaded,
	}

	http.HandleFunc("/register", server.RegisterHandler)
	http.HandleFunc("/login", server.LoginHandler)

	fmt.Println("Server is running on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
