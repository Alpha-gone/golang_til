package main

import "fmt"

func main() {
	str1 := "Hello\t'World'"
	str2 := `Go is "awesome"!\nGo is simple and\t'powerful'`

	fmt.Println(str1)
	fmt.Println(str2)
}
