package main

type Integer interface {
	int8 | int16 | int32 | int64 | int
}

func add[T Integer](a, b T) T {
	return a + b
}

type MyInt int

func main() {
	add(1, 2)

	var a MyInt = 3
	var b MyInt = 5
	add(a, b)
}
