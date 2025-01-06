package DSA

import (
	"math"
	"sort"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
28. Find the Index of the First Occurrence in a String

Dạng bài: String,
Phương pháp tiếp cận:
1. Sử dụng hàm strings.Index để tìm chỉ số của phần tử target.
2. Trả về chỉ số đó.
Level: Easy
Time Complexity: O(n)
Space Complexity: O(1)
*/
func StrStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

/*
135. Candy

Dạng bài: Array, Greedy
Phương pháp tiếp cận:
1. Sử dụng hai mảng left và right để lưu trữ số lượng kẹo tối thiểu cần cho mỗi học sinh.
2. Duyệt qua mảng ratings và cập nhật mảng left và right.
3. Tính tổng số kẹo tối thiểu cần cho mỗi học sinh.
4. Trả về tổng số kẹo tối thiểu.
Level: Hard
Time Complexity: O(n)
Space Complexity: O(n)
*/
func Candy(ratings []int) int {
	n := len(ratings)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = 1
		right[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	var total int
	for i := 0; i < n; i++ {
		total += max(left[i], right[i])
	}
	return total
}

/*
2089. Find Target Indices After Sorting Array

Dạng bài: Array,
Phương pháp tiếp cận:
1. Sắp xếp mảng.
2. Duyệt qua mảng và tìm các chỉ số của phần tử target.
3. Trả về các chỉ số đó.
Level: Easy
Time Complexity: O(n log n)
Space Complexity: O(n)
*/
func TargetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	var result []int
	for i, num := range nums {
		if num == target {
			result = append(result, i)
		}
	}
	return result
}

// 1365. How Many Numbers Are Smaller Than the Current Number
func SmallerNumbersThanCurrent(nums []int) []int {
	count := make([]int, 101)
	for _, num := range nums {
		count[num]++
	}
	for i := 1; i <= 100; i++ {
		count[i] += count[i-1]
	}
	var result []int
	for _, num := range nums {
		if num == 0 {
			result = append(result, 0)
		} else {
			result = append(result, count[num-1])
		}
	}
	return result
}

/*
2545. Sort the Students by Their Kth Score

Dạng bài: Array,
Phương pháp tiếp cận:
1. Sắp xếp mảng.
2. Duyệt qua mảng và tìm các chỉ số của phần tử target.
3. Trả về các chỉ số đó.
Level: Easy
Time Complexity: O(n log n)
Space Complexity: O(n)
*/
func SortStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})
	return score
}

/*
1122. Relative Sort Array

Dạng bài: Array, Hash Table, Sorting
Phương pháp tiếp cận:
1. Sử dụng map để lưu trữ số lần xuất hiện của mỗi phần tử trong arr1.
2. Duyệt qua mảng arr2 và thêm các phần tử tương ứng vào kết quả.
3. Thêm các phần tử còn lại từ arr1 vào kết quả.
4. Trả về kết quả.
Level: Easy
Time Complexity: O(n)
Space Complexity: O(n)
*/
func RelativeSortArray(arr1 []int, arr2 []int) []int {
	// Create a map to store the frequency of each element in arr1
	freqMap := make(map[int]int)
	for _, num := range arr1 {
		freqMap[num]++
	}

	// Initialize the result slice
	var result []int

	// First, add elements from arr2 in the order they appear in arr2
	for _, num := range arr2 {
		for freqMap[num] > 0 {
			result = append(result, num)
			freqMap[num]--
		}
	}

	// Then, add the remaining elements from arr1 in ascending order
	for num := range freqMap {
		for freqMap[num] > 0 {
			result = append(result, num)
			freqMap[num]--
		}
	}

	return result
}

/*
45. Jump Game II
Problem: Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Your goal is to reach the last index in the minimum number of jumps.
You can assume that you can always reach the last index.
Return the minimum number of jumps to reach the last index.
Example:
Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2.
Jump 1 step from index 0 to 1, then 3 steps to the last index.

Phương pháp tiếp cận:
1. Sử dụng quy hoạch động để lưu trữ số lượng nhảy tối thiểu để đạt đến mỗi chỉ số.
2. Khởi tạo một mảng dp với giá trị có thể lớn nhất.
3. Đặt chỉ số đầu tiên là 0 vì chỉ cần 0 nhảy để đạt đến chỉ số đầu tiên.
4. Duyệt qua mảng và đối với mỗi chỉ số, duyệt qua tất cả các chỉ số trước đó.
5. Nếu chỉ số hiện tại có thể đạt được từ chỉ số trước đó, cập nhật mảng dp.
6. Giá trị tại chỉ số cuối cùng sẽ là số lượng nhảy tối thiểu để đạt đến chỉ số cuối cùng.

	Time Complexity: O(n^2)
	Space Complexity: O(n)
*/
func Jump1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] >= i-j {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

/*
Phương pháp tiếp cận:
1. Sử dụng thuật toán tham lam để tìm số lượng nhảy tối thiểu để đạt đến điểm cuối cùng của mảng.
2. Duyệt qua mảng và cập nhật khoảng cách tối đa mà bạn có thể đạt được từ mỗi vị trí.
3. Khi bạn đạt đến điểm cuối cùng của mảng, trả về số lượng nhảy.

Time Complexity: O(n)
Space Complexity: O(1)
*/
func Jump2(nums []int) int {
	jumps := 0
	currentEnd := 0
	farthest := 0
	// Duyệt qua mảng và cập nhật khoảng cách tối đa mà bạn có thể đạt được từ mỗi vị trí.
	// chú ý: i < len(nums)-1 vì chỉ cần đến điểm cuối cùng của mảng.
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}
	return jumps
}
