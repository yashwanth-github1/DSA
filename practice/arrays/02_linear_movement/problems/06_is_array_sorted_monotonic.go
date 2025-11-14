package main

// Check if array is sorted in ascending order
func isSortedAscending(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ { // single pointer i
		if nums[i] > nums[i+1] { // compare consecutive elements
			return false
		}
	}
	return true
}

// Check if array is sorted in descending order
func isSortedDescending(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			return false
		}
	}
	return true
}

// Check if array is sorted (either ascending or descending)
func isSorted(nums []int) bool {
	return isSortedAscending(nums) || isSortedDescending(nums)
}

//Check if Array is Strictly Increasing
func isStrictlyIncreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

//Check if Array is Non-Increasing
func isNonIncreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			return false
		}
	}
	return true
}
