package main

import "fmt"

func main() {
	numbers := []int{5, 5, -4, 8, 11, 1, -2, 5}
	targetNum := 10

	fmt.Println(TwoNumberSum(numbers, targetNum))
}
