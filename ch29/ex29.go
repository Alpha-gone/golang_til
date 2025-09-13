package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

type Students []student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

var students map[int]student
var lastId int

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/students", getStudentListHandler).Methods("GET")
	mux.HandleFunc("/student/{id:[0-9]+}", getStudentHandler).Methods("GET")
	mux.HandleFunc("/students", postStudentHandler).Methods("POST")
	mux.HandleFunc("/students/{id:[0-9]+}", deleteStudentHandler).Methods("DELETE")

	students = make(map[int]student)
	students[1] = student{1, "aaa", 16, 87}
	students[2] = student{2, "bbb", 18, 98}
	lastId = 2

	return mux
}

func getStudentListHandler(writer http.ResponseWriter, requst *http.Request) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(list)
}

func getStudentHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)

		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(student)
}

func postStudentHandler(writer http.ResponseWriter, req *http.Request) {
	var student student
	err := json.NewDecoder(req.Body).Decode(&student)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	lastId++
	student.Id = lastId
	students[lastId] = student

	writer.WriteHeader(http.StatusCreated)
}

func deleteStudentHandler(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)

		return
	}

	delete(students, id)
	writer.WriteHeader(http.StatusOK)
}
