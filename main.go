package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load .env file")
	}

	PORT := os.Getenv("PORT")
	router := chi.NewRouter()

	fmt.Printf("Sever starting on port: %v", PORT)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
