package main

func TwoNumberSum(array []int, target int) []int {
	l := len(array)

	for i, num := range array {
		for j := i + 1; j < l; j++ {
			num2 := array[j]
			if num+num2 == target {
				return []int{num, num2}
			}
		}
	}

	return []int{}
}
