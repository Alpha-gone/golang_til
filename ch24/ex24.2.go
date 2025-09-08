package main

import "fmt"

func print[T1 any, T2 any](a T1, b T2) {
	fmt.Println(a, b)
}

func print2(a, b interface{}) {
	fmt.Println(a, b)
}

func main() {
	print(1, 2)
	print(3.14, 1.43)
	print("Hello", "World")
	print(1, "Hello")

	print2(1, 2)
	print2(3.14, 1.43)
	print2("Hello", "World")
	print2(1, "Hello")
}
