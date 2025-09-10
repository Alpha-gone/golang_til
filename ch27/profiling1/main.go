package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func Fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}

func main() {
	f, err := os.Create("./ch27/profiling1/cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	fmt.Println(Fib(50))

	time.Sleep(10 * time.Second)
}
