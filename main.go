package main

import (
	"auth-api/routes"
	"auth-api/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.InitializeMongoDB("mongodb://mongo:27017")
	r := mux.NewRouter()
	routes.RegisterAuthRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
