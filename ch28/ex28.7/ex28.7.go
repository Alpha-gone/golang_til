package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/",
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "Hello World")
		})
	runPath := "./ch28/ex28.7/"

	err := http.ListenAndServeTLS(":3000", runPath+"localhost.crt", runPath+"localhost.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
