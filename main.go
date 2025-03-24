package main

import (
	"net/http"
	"log"
	"auth_service/handlers"
)

func main () {
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
