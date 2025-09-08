package main

import "fmt"

type Node[T any] struct {
	val  T
	next *Node[T]
}

func NewNode[T any](v T) *Node[T] {
	return &Node[T]{val: v}
}

func (n *Node[T]) push(v T) *Node[T] {
	node := NewNode(v)

	n.next = node
	return node
}

func main() {
	node1 := NewNode(1)
	node1.push(2).push(3).push(4)

	for node1 != nil {
		fmt.Print(node1.val, " - ")
		node1 = node1.next
	}
	fmt.Println()

	node2 := NewNode("Hi")
	node2.push("Hello").push("How are you")

	for node2 != nil {
		fmt.Print(node2.val, " - ")
		node2 = node2.next
	}
	fmt.Println()
}
