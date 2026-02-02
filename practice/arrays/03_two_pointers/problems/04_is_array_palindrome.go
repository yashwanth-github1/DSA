package main

// Check if array is palindrome using two pointers
func isPalindrome(nums []int) bool {

	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] != nums[r] {
			return false
		}
		l++
		r--
	}
	return true
}
