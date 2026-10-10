package main

import (
	"log"
	"net/http"
	"os"

	"entropy-library/server"
)

func main() {
	address := os.Getenv("ADDR")
	if address == "" {
		address = ":8080"
	}

	log.Printf("entropy server listening on %s", address)
	if err := http.ListenAndServe(address, server.NewHandler()); err != nil {
		log.Fatal(err)
	}
}
