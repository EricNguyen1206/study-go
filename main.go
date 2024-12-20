package main

import (
	"fmt"
	"study-go/hello/Day1"
)

func main() {
	fmt.Println(Day1.TriangleArea(10, 20))
	fmt.Println(Day1.EvenString("hello"))
	fmt.Println(Day1.CalculateArray([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println(Day1.TwoSum([]int{2, 7, 11, 15}, 9))
}
