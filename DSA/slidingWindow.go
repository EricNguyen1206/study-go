package DSA

import "math"

// 3. Longest Substring Without Repeating Characters
func LengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[byte]int)
	windowStart := 0
	maxLength := 0

	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		rightChar := s[windowEnd]
		if _, exists := charIndexMap[rightChar]; exists {
			windowStart = int(math.Max(float64(windowStart), float64(charIndexMap[rightChar]+1)))
		}
		charIndexMap[rightChar] = windowEnd
		maxLength = int(math.Max(float64(maxLength), float64(windowEnd-windowStart+1)))
	}
	return maxLength
}

// 1493. Longest Subarray of 1's After Deleting One Element
func LongestSubarray(nums []int) int {
	maxLength := 0
	zeroIndex := -1
	count := 0

	for i, num := range nums {
		if num == 0 {
			count++
			if count > 1 {
				// Move the left pointer to the right of the last zero
				maxLength = int(math.Max(float64(maxLength), float64(i-zeroIndex-1)))
				zeroIndex = i - 1 // Update zeroIndex to the last zero's position
			}
		}
	}
	// Handle the case where there is at most one zero
	return int(math.Max(float64(maxLength), float64(len(nums)-zeroIndex-1)))
}
