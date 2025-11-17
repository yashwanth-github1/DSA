package main

// Check if array is palindrome using two pointers
func isPalindrome(nums []int) bool {
	// Movement → Condition → Operation → State → Result
	// Movement: l → right, r → left
	// Condition: l < r
	// Operation: Compare nums[l] and nums[r]
	// State: Update l++, r--
	// Result: Return true if all pairs matched

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
