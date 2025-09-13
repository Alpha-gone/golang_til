package main

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

var students map[int]Student
var lastId int

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func main() {
	r := gin.Default()
	setupHandlers(r)
	r.Run(":3000")
}

func setupHandlers(g *gin.Engine) {
	g.GET("/students", getStudentsHandler)
	g.GET("/student/:id", getStudentHandler)
	g.POST("/student", postStudentHandler)
	g.DELETE("/student/:id", deleteStudentHandler)

	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2
}

func getStudentsHandler(c *gin.Context) {
	list := make(Students, 0)

	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)
	c.JSON(http.StatusOK, list)
}

func getStudentHandler(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	student, ok := students[id]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)

		return
	}

	c.JSON(http.StatusOK, student)
}

func postStudentHandler(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	lastId++
	student.Id = lastId
	students[lastId] = student

	c.String(http.StatusOK, "Success to add id: %d", lastId)
}

func deleteStudentHandler(c *gin.Context) {
	idstr := c.Param("id")
	if idstr == "" {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	delete(students, id)
	c.String(http.StatusOK, "sucess to delete")
}
