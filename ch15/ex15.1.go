package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())

	n := rand.Intn(100)
	fmt.Println(n)
}
