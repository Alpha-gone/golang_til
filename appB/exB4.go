package main

import "fmt"

func main() {
	var slice []int

	slice, allocCnt := append1000times(slice)
	fmt.Println("allocCnt:", allocCnt, "cap:", cap(slice))

	var slice2 = make([]int, 0, 1000)
	slice, allocCnt = append1000times(slice2)
	fmt.Println("allocCnt:", allocCnt, "cap:", cap(slice))
}

func append1000times(slice []int) ([]int, int) {
	var lastCap = cap(slice)
	var allocCnt int = 0

	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
		if lastCap != cap(slice) {
			allocCnt++
			lastCap = cap(slice)
		}
	}

	return slice, allocCnt
}
