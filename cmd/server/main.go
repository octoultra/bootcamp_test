package main

import (
	"entrytest/internal/api"
	"entrytest/internal/store"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("frontend")))

	api.Init(store.New()).RegisterRoutes(mux)

	log.Printf("сервер слушает http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
