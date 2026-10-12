package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
)

//go:embed static/*
var static embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	staticFiles, err := fs.Sub(static, "static")
	if err != nil {
		fmt.Println("Unable to load app assets:", err)
		os.Exit(1)
	}
	http.Handle("/", http.FileServer(http.FS(staticFiles)))

	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, nil)
}
