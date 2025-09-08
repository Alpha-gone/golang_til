package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera"}

	fmt.Println("Equal:", slices.Compare(names, []string{"Alice", "Bob", "Vera"}))
	fmt.Println("V < X:", slices.Compare(names, []string{"Alice", "Bob", "Xera"}))
	fmt.Println("V > C:", slices.Compare(names, []string{"Alice", "Bob", "Cat"}))
	fmt.Println("3 > 2:", slices.Compare(names, []string{"Alice", "Bob"}))
}
