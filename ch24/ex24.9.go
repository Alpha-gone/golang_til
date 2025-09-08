package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera"}
	n, found := slices.BinarySearch(names, "Vera")
	fmt.Println("Vara:", n, found)

	n, found = slices.BinarySearch(names, "Bill")
	fmt.Println("Bill:", n, found)
}
