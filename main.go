package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load .env file")
	}

	PORT := os.Getenv("PORT")
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/health", handlerReadiness)

	router.Mount("/v1", v1Router)

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
