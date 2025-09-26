// Package app contains the application bootstrap logic (wiring DB, HTTP server, and graceful shutdown).
package app

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
)

// Run initializes and starts the application, and handles graceful shutdown.
func Run() error {
	// load env
	_ = godotenv.Load()

	// port from env (default 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// init db
	conn, err := db.NewPostgres()
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	// init http mux
	mux := http.NewServeMux()
	orghttp.RegisterOrganisationRoutes(mux, conn)

	// server
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// run
	go func() {
		log.Printf("API running on :%s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}
