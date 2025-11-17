package main

//Reverse Array

func ReverseArray(nums []int) {
	// Movement → Condition → Operation → State → Result
	// Movement: l → right, r → left
	// Condition: l < r
	// Operation: swap(nums[l], nums[r])
	// State: update l++, r--
	// Result: array reversed in place

	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l] // swap
		l++
		r--
	}
}

//Reverse Only Part of Array
//Given an array, reverse the elements from index start to end.
func ReversePart(nums []int, start, end int) {
	l, r := start, end
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
