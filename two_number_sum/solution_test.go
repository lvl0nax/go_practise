package main

import (
	"fmt"
	"reflect"
	"testing"
)

const target = 10

func getResults() [][]int {
	return [][]int{
		[]int{9, 1},
		[]int{11, -1},
		[]int{},
		[]int{5, 5},
		[]int{},
	}
}

func getInputs() [][]int {
	return [][]int{
		[]int{3, 5, 9, 8, 11, 1, -2, 6},
		[]int{3, 5, -4, 8, 11, -1, -2, 6},
		[]int{3, 5, -4, 8, -11, -1},
		[]int{5, 6, -4, 8, 11, -5, -2, 5},
		[]int{10},
	}
}

// go test -v
func TestTwoNumberSum(t *testing.T) {
	inputs := getInputs()
	for i, result := range getResults() {
		returned := TwoNumberSum(inputs[i], target)

		if !reflect.DeepEqual(result, returned) {
			t.Errorf(
				"\n Expected result: %v \n   Actual result: %v",
				result,
				returned,
			)
		}
	}
}

func ExampleTwoNumberSum() {
	inputs := getInputs()
	for _, example := range inputs {
		fmt.Println(TwoNumberSum(example, target))
	}
	// Output:
	// [9 1]
	// [11 -1]
	// []
	// [5 5]
	// []
}
