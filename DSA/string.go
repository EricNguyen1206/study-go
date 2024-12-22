package DSA

import "strings"

// 125. Valid Palindrome
func IsPalindrome(s string) bool {
	// remove non-alphanumeric characters
	s = strings.ReplaceAll(s, " ", "")
	// convert to lowercase
	s = strings.ToLower(s)
	// check if the string is a palindrome
	return s == reverse(s)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 242. Valid Anagram
func IsAnagram(s string, t string) bool {
	// check if the length of the strings are the same
	if len(s) != len(t) {
		return false
	}
	// count the frequency of each character in s
	count := make(map[rune]int)
	for _, char := range s {
		count[char]++
	}
	// check if the frequency of each character in t is the same as in s
	for _, char := range t {
		count[char]--
		if count[char] < 0 {
			return false
		}
	}
	return true
}

// 1528. Shuffle String
func RestoreString(s string, indices []int) string {
	// create a new string with the same length as s
	result := make([]rune, len(s))
	// iterate over the indices and place the characters in the correct position
	for i, index := range indices {
		result[index] = rune(s[i])
	}
	return string(result)
}
