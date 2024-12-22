package DSA

// easy

// medium
// 238. Product of Array Except Self
func ProductOfArrayExceptSelf(nums []int) []int {
	lprd, rprd, n := 1, 1, len(nums)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			temp[i] = lprd
		} else {
			temp[i] = lprd * nums[i-1]
			lprd *= nums[i-1]
		}
	}
	for i := n - 1; i >= 0; i-- {
		temp[i] *= rprd
		rprd *= nums[i]
	}
	return temp
}

// hard
