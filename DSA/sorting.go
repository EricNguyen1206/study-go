package DSA

import "sort"

// 912. Sort an Array
func SortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}

// Bubble Sort
func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// Selection Sort
func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
	}
	return nums
}

// Quick Sort
// Ý tưởng: Chọn một phần tử pivot, sau đó chia mảng thành 2 mảng con, một mảng con bao gồm các phần tử nhỏ hơn pivot và một mảng con bao gồm các phần tử lớn hơn pivot
// Cách thực hiện:
// 1. Chọn một phần tử pivot
// 2. Chia mảng thành 2 mảng con
// 3. Gọi đệ quy để sắp xếp mỗi mảng con
// 4. Trả về mảng kết quả
// Độ phức tạp: O(n log n) time, O(log n) space
func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := nums[len(nums)/2]
	left, right := []int{}, []int{}
	for _, num := range nums {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

// Merge Sort
// Ý tưởng: Chia đôi mảng thành 2 mảng con, sau đó sắp xếp mỗi mảng con rồi merge lại
// Cách thực hiện:
// 1. Chia đôi mảng thành 2 mảng con
// 2. Sắp xếp mỗi mảng con
// 3. Merge 2 mảng con bằng cách so sánh phần tử nhỏ hơn của 2 mảng con, sau đó thêm phần tử nhỏ hơn vào mảng kết quả
// 4. Trả về mảng kết quả
// Độ phức tạp: O(n log n) time, O(n) space
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	// Chia đôi mảng thành 2 mảng con
	mid := len(nums) / 2
	// Gọi đệ quy để sắp xếp mỗi mảng con
	left, right := MergeSort(nums[:mid]), MergeSort(nums[mid:])
	// Merge 2 mảng con
	return merge(left, right)
}

// Hàm merge 2 mảng con
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	// Merge 2 mảng con
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// Thêm phần tử còn lại của mảng left vào mảng kết quả
	return append(result, left[i:]...)
}

// PRACTICE

// 1365. How Many Numbers Are Smaller Than the Current Number => check top150interview.go
// func SmallerNumbersThanCurrent(nums []int) []int {}

// 1122. Relative Sort Array => check top150interview.go
// func RelativeSortArray(arr1 []int, arr2 []int) []int {}
