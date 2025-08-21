package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
}

func (receiver *Student) String() string {
	return "Student"
}

type Actor struct {
}

func (a *Actor) String() string {
	return "Actor"
}

func ConvertType(stringer Stringer) {
	student := stringer.(*Student)
	fmt.Println(student)
}

func main() {
	actor := &Actor{}
	ConvertType(actor)
}
