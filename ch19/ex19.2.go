package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./ch19/ex19_2_test.txt")
	if err != nil {
		fmt.Println("Faild to create a file")
	}

	defer fmt.Println("반드시 호출 됩니다.")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")

	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(f, "Hello World")
}
