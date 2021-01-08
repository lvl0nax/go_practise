package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	Input  int
	Result int
}

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711
func testCases() []TestCase {
	return []TestCase{
		{25, 21},
		{150, 144},
		{10000, 10946},
		{88, 89},
	}
}

// go test -v
func TestClosestFibonacci(t *testing.T) {
	for _, testCase := range testCases() {
		actualResult := ClosestFibonacci(testCase.Input)

		if !reflect.DeepEqual(testCase.Result, actualResult) {
			t.Errorf("\n Expected result: %v \n   Actual result: %v", testCase.Result, actualResult)
		}
	}
}
