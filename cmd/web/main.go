package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.NewJSONHandler(os.Stdout, nil)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	slog.Info("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
