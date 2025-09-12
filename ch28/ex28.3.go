package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"/",
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "Hello World")
		})

	mux.HandleFunc(
		"/bar",
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "Hello Bar")
		})

	http.ListenAndServe(":3000", mux)
}
