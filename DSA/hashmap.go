package DSA

import (
	"reflect"
)

// 1. Two sum
func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		if _, ok := hashMap[target-num]; ok {
			return []int{hashMap[target-num], i}
		}
		hashMap[num] = i
	}
	return nil
}

// 1748. Sum of Unique Elements
func SumOfUniqueElements(nums []int) int {
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}
	sum := 0
	for num, count := range hashMap {
		if count == 1 {
			sum += num
		}
	}
	return sum
}

// 383. Ransom Note
func CanConstruct(ransomNote string, magazine string) bool {
	hashMap := make(map[rune]int)
	for _, char := range magazine {
		hashMap[char]++
	}
	for _, char := range ransomNote {
		if hashMap[char] == 0 {
			return false
		}
		hashMap[char]--
	}
	return true
}

// Solution 2
// func canConstruct(ransomNote string, magazine string) bool {
// 	for _, v := range ransomNote {
// 		if strings.Count(ransomNote, string(v)) > strings.Count(magazine, string(v)) {
// 			return false
// 		}
// 	}
// 	return true
// }

// 454. 4Sum II
// time complexity: O(n^2)
// space complexity: O(n)
func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hashMap := make(map[int]int)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			hashMap[num1+num2]++
		}
	}

	count := 0
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			count += hashMap[-num3-num4]
		}
	}
	return count
}

// 1394. Find Lucky Integer in an Array
func FindLucky(arr []int) int {
	hashMap := make(map[int]int)
	for _, num := range arr {
		hashMap[num]++
	}
	for num, count := range hashMap {
		if num == count {
			return num
		}
	}
	return -1
}

// 1742. Maximum Number of Balls in a Box
func sumOfDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func CountBalls(lowLimit int, highLimit int) int {
	hashMap := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		hashMap[sumOfDigits(i)]++
	}
	max := 0
	for _, count := range hashMap {
		if count > max {
			max = count
		}
	}
	return max
}

// 771. Jewels and Stones
func NumJewelsInStones(jewels string, stones string) int {
	hashMap := make(map[rune]bool)
	for _, jewel := range jewels {
		hashMap[jewel] = true
	}
	count := 0
	for _, stone := range stones {
		if hashMap[stone] {
			count++
		}
	}
	return count
}

// 2405. Optimal Partition of String
func partitionString(s string) int {
	hashMap := make(map[rune]bool)
	count := 1
	for _, char := range s {
		if hashMap[char] {
			count++
			hashMap = make(map[rune]bool)
		}
		hashMap[char] = true
	}
	return count
}

// 20. Valid Parentheses
func isValid(s string) bool {
	hashMap := make(map[rune]rune)
	hashMap['('] = ')'
	hashMap['['] = ']'
	hashMap['{'] = '}'
	stack := []rune{}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if len(stack) > 0 && hashMap[stack[len(stack)-1]] == char {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

// 1657. Determine if Two Strings Are Close
func CloseStrings(word1 string, word2 string) bool {
	hashMap1 := make(map[rune]int)
	hashMap2 := make(map[rune]int)
	for _, char := range word1 {
		hashMap1[char]++
	}
	for _, char := range word2 {
		hashMap2[char]++
	}
	return reflect.DeepEqual(hashMap1, hashMap2)
}
