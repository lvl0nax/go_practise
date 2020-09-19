package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	Array1 []int
	Array2 []int
	Result []int
}

func testCases() []TestCase {
	tc1 := TestCase{
		[]int{-1, 5, 10, 20, 28, 3},
		[]int{26, 134, 135, 15, 17},
		[]int{28, 26},
	}
	tc2 := TestCase{
		[]int{-1, 5, 0, 20, 28, 3},
		[]int{26, 134, 0, 15, 17},
		[]int{0, 0},
	}
	tc3 := TestCase{
		[]int{-1},
		[]int{134},
		[]int{-1, 134},
	}
	tc4 := TestCase{
		[]int{-1, 5, 0, 20, 28, 326},
		[]int{26, 134, 322, 325, 326},
		[]int{326, 326},
	}

	return []TestCase{tc1, tc2, tc3, tc4}
}

// go test -v
func TestSmallestDifference1(t *testing.T) {
	for _, testCase := range testCases() {
		returned := SmallestDifference1(testCase.Array1, testCase.Array2)

		if !reflect.DeepEqual(testCase.Result, returned) {
			t.Errorf(
				"\n Expected result: %v \n   Actual result: %v",
				testCase.Result,
				returned,
			)
		}
	}
}

func TestSmallestDifference2(t *testing.T) {
	for _, testCase := range testCases() {
		returned := SmallestDifference2(testCase.Array1, testCase.Array2)

		if !reflect.DeepEqual(testCase.Result, returned) {
			t.Errorf(
				"\n Expected result: %v \n   Actual result: %v",
				testCase.Result,
				returned,
			)
		}
	}
}

func ExampleSmallestDifference1() {
	for _, testCase := range testCases() {
		fmt.Println(SmallestDifference1(testCase.Array1, testCase.Array2))
	}
	// Output:
	// [28 26]
	// [0 0]
	// [-1 134]
	// [326 326]
}

func ExampleSmallestDifference2() {
	for _, testCase := range testCases() {
		fmt.Println(SmallestDifference2(testCase.Array1, testCase.Array2))
	}
	// Output:
	// [28 26]
	// [0 0]
	// [-1 134]
	// [326 326]
}
