package main

import (
	"fmt"
	"hash/fnv"
)

type ComparableHasher interface {
	comparable
	Hash() uint32
}

type Mystring string

func (s Mystring) Hash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))

	return h.Sum32()
}

func equal[T ComparableHasher](a, b T) bool {
	if a == b {
		return true
	}

	return a.Hash() == b.Hash()
}

func main() {
	var str1 Mystring = "Hello"
	var str2 Mystring = "World"

	fmt.Println(equal(str1, str2))
}
