package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	Name  string
	Age   int
	Score int
}

func main() {
	http.ListenAndServe(":3000", makeWebHandler())
}

func makeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", studentHandler)

	return mux
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	student := student{"aaa", 16, 87}

	data, _ := json.Marshal(student)

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))

}
