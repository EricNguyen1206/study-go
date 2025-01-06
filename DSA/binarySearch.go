package DSA

// 367. Valid Perfect Square
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
