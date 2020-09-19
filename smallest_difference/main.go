package main

import (
	"fmt"
	"math"
	"sort"
)

// First implementation - straignt way
func SmallestDifference1(array1, array2 []int) []int {
	result := []int{array1[0], array2[0]}
	minDiff := int(math.Abs(float64(array1[0] - array2[0])))
	for _, n1 := range array1 {
		for _, n2 := range array2 {
			diff := int(math.Abs(float64(n2 - n1)))
			if minDiff > diff {
				minDiff = diff
				result = []int{n1, n2}
			}
		}
	}
	return result
}

// thought a bit about the complexity
func SmallestDifference2(array1, array2 []int) []int {
	sort.Ints(array1) // O(n log n)
	sort.Ints(array2) // O(m log m)

	i1, i2 := 0, 0
	result := []int{}
	minDiff, diff := math.MaxInt32, math.MaxInt32

	// O(n+m)
	for i1 < len(array1) && i2 < len(array2) {
		num1, num2 := array1[i1], array2[i2]

		if num1 < num2 {
			diff = num2 - num1
			i1++
		} else if num2 < num1 {
			diff = num1 - num2
			i2++
		} else {
			return []int{num1, num2}
		}
		if diff < minDiff {
			minDiff = diff
			result = []int{num1, num2}
		}
	}

	return result
}

func main() {
	array1 := []int{-1, 5, 0, 20, 28, 326}
	array2 := []int{26, 134, 322, 325, 326}

	fmt.Println(SmallestDifference1(array1, array2))
	fmt.Println(SmallestDifference2(array1, array2))
}
