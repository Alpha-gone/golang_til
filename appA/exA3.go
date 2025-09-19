package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("appA/output.txt")
	if err != nil {
		fmt.Errorf("Create: %v\n", err)

		return
	}
	defer f.Close()

	const name, age = "Kim", 7
	n, err := fmt.Fprint(f, name, " is ", age, " years old.\n")
	if err != nil {
		fmt.Errorf("Fprint: %v\n", err)
	}

	fmt.Print(n, " byte written\n")
}
