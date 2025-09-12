package main

import "net/http"

func main() {
	http.Handle(
		"/",
		http.FileServer(http.Dir("./ch28/static")))

	http.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./ch28/static"))))

	http.ListenAndServe(":3000", nil)
}
