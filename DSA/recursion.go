package DSA

// 509. Fibonacci Number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// 326. Power of Three
func IsPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

// 202. Happy Number
func IsHappy(n int) bool {
	visited := map[int]bool{}

	for {
		nextValue := 0
		for ; n > 0; n = n / 10 {
			nextValue += (n % 10) * (n % 10)
		}

		if _, ok := visited[nextValue]; ok {
			// only cycled around 1 means happy
			return nextValue == 1
		} else {
			visited[nextValue] = true
		}

		n = nextValue
	}
}

// 1545. Find Kth Bit in Nth Binary String
func FindKthBit(n int, k int) byte {
	// base case
	if n == 1 {
		return '0'
	}
	// find the length of current string
	length := 1<<n - 1
	// find the middle index of current string
	mid := length/2 + 1
	// if k is the middle index, return '1'
	if k == mid {
		return '1'
	}
	// if k is less than the middle index, find the kth bit in the previous string
	if k < mid {
		return FindKthBit(n-1, k)
	}
	// if k is greater than the middle index,
	// get the kth bit in the previous string and invert it
	if FindKthBit(n-1, length-k+1) == '0' {
		return '1'
	}
	return '0'
}
