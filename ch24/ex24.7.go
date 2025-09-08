package main

import "fmt"

type nodeType1 struct {
	val  interface{}
	next *nodeType1
}

type nodeType2[T any] struct {
	val  T
	next *nodeType2[T]
}

func main() {
	node1 := &nodeType1{val: 1}
	node2 := &nodeType2[int]{val: 2}

	var v1 int = node1.val
	fmt.Println(v1)

	var v2 int = node2.val
	fmt.Println(v2)
}
