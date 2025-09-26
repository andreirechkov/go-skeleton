// Package main is the entry point of the application.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	orghttp "github.com/andreirechkov/go-skeleton/internal/modules/organisations/interfaces/http"
	"github.com/andreirechkov/go-skeleton/internal/shared/db"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // register postgres driver
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
	defer func() { _ = conn.Close() }()

	if err := conn.Ping(); err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}

	// init http mux and routes
	mux := http.NewServeMux()
	orghttp.RegisterOrganisationRoutes(mux, conn)

	// http server with sane timeouts
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		log.Println("API running on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	// wait for interrupt/terminate
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down...")

	// graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
