package main

//Remove Duplicates from Sorted Array (In-place)

// RemoveDuplicatesSorted removes duplicates in-place using two pointers
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0

	for fast := 1; fast < len(nums); fast++ { // movement
		if nums[fast] != nums[slow] { // condition
			slow++                  // operation
			nums[slow] = nums[fast] // place unique element
		}
	}

	return slow + 1 // length of unique portion
}
