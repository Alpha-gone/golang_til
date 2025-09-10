package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var fibMap [65525]int

func Fib(n int) int {
	f := fibMap[n]
	if f > 0 {
		return f
	}

	if n == 0 {
		return 0
	} else if n == 1 {
		f = 1
	} else {
		f = Fib(n-1) + Fib(n-2)
	}

	fibMap[n] = f
	return f
}

func main() {
	f, err := os.Create("./ch27/profiling2/cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	fmt.Println(Fib(50))

	time.Sleep(10 * time.Second)
}
