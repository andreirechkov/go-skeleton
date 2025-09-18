package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/andreirechkov/go-skeleton/internal/shared/db"
	"github.com/andreirechkov/go-skeleton/internal/modules/organisations/api"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, fallback to system env")
	}

	// init db
	conn, err := db.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	} else {
		log.Println("Successfully connected to Postgres")
	}

	// init http
	mux := http.NewServeMux()
	api.RegisterOrganisationRoutes(mux, conn)

	log.Println("API running on :8080")
	http.ListenAndServe(":8080", mux)
}