package main

import (
	"log"
	"net/http"
	"github.com/Khamitamantaev/restapigolang/rest-go/store"
	//"os"
	"github.com/gorilla/handlers"
)

func main() {
	port := "8000"

	router := store.NewRouter() // create routes

	// These two lines are important if you're designing a front-end to utilise this API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}