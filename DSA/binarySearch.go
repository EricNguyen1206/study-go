package DSA

// PRACTICE

// 704. Binary Search
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 367. Valid Perfect Square
// Ý tưởng: Sử dụng binary search để tìm căn bậc hai của num
// Cách thực hiện:
// 1. Nếu num < 0, trả về false
// 2. Khởi tạo left = 0, right = num
// 3. Lặp cho đến khi left <= right
// 4. Tính mid = (left + right) / 2
// 5. Tính square = mid * mid
// 6. Nếu square == num, trả về true
// 7. Nếu square < num, tăng left lên 1
// 8. Nếu square > num, giảm right xuống 1
// 9. Trả về false
// Độ phức tạp: O(log n) time, O(1) space
func IsPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}
	left, right := 0, num
	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid
		if square == num {
			return true
		} else if square < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 441. Arranging Coins
// This function is used to arrange coins in a staircase.
// It uses binary search to find the maximum number of complete rows that can be formed.
// The total number of coins in a complete row is given by the formula (n * (n + 1)) / 2.
// If the total number of coins is equal to the given number, the function returns the current row.
// If the total number of coins is less than the given number, the function moves to the next row.
// If the total number of coins is greater than the given number, the function moves to the previous row.
func ArrangeCoins(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		total := (mid * (mid + 1)) / 2
		if total == n {
			return mid
		} else if total < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

// 852. Peak Index in a Mountain Array
// Ý tưởng: Sử dụng binary search để tìm đỉnh của mảng
// Cách thực hiện:
// 1. Khởi tạo left = 0, right = len(arr)-1
// 2. Lặp cho đến khi left < right
// 3. Tính mid = (left + right) / 2
// 4. Nếu arr[mid] > arr[mid+1], giảm right xuống 1
// 5. Nếu arr[mid] < arr[mid+1], tăng left lên 1
// 6. Trả về left
// Độ phức tạp: O(log n) time, O(1) space
func PeakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
