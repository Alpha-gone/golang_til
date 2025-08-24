package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer) {
	writer("Hello world")
}

func main() {
	f, err := os.Create("./ch19/ex19_7_text.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
