package main

import (
	"log"

	"github.com/andreirechkov/go-skeleton/internal/app"
	_ "github.com/lib/pq" // register postgres driver
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
