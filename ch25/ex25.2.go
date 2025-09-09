package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	PrintFile("./ch25/files/hamlet.txt")
}

func PrintFile(fileNames string) {
	file, err := os.Open(fileNames)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", fileNames)

		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
