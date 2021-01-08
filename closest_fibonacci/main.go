package main

import (
	"fmt"
	"math"
)

func ClosestFibonacci(num int) int {
	n1 := 0
	n2 := 1
	diff := math.Abs(float64(num - n2))

	for {
		n1, n2 = n2, n1+n2
		if diff >= math.Abs(float64(num-n2)) {
			diff = math.Abs(float64(num - n2))
		} else {
			return n1
		}
	}
}

func main() {
	num1 := ClosestFibonacci(42)
	fmt.Println(num1) // => 34
}
