package main

import (
	"fmt"
)

// go test -v
func ExampleBuildTreeAndCalculate() {
	fmt.Println(BuildTreeAndCalculate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(BuildTreeAndCalculate(1, 2, 3))
	fmt.Println(BuildTreeAndCalculate(1))
	fmt.Println(BuildTreeAndCalculate(3, 7))
	fmt.Println(BuildTreeAndCalculate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1))
	// Output:
	// [15 16 18 10 11]
	// [3 4]
	// [1]
	// [10]
	// [15 16 18 9 11 11 11]
}
