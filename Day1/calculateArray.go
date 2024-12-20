package Day1

import (
	"fmt"
	"sort"
)

func getArray() []int {
	fmt.Println("Please enter the array:")
	var arr []int
	fmt.Scanln(&arr)
	return arr
}

func Sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func Min(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func CalculateArray(arr []int) (int, int, int, int, []int) {
	var myArray []int
	if len(arr) == 0 {
		myArray = getArray()
	} else {
		myArray = arr
	}
	sort.Ints(myArray)
	return Sum(myArray), Min(myArray), Max(myArray), Sum(myArray) / len(myArray), myArray
}
