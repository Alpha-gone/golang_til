package main

import (
	"embed"
	"net/http"
)

// go:embed static/*
var files embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(files)))
	http.ListenAndServe(":3000", nil)
}
