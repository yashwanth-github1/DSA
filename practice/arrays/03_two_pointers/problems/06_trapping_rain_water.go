package main

//Trapping Rain Water (Two Pointers)

//Given n non-negative integers representing elevation map where the width of each bar is 1,
// compute how much water it can trap after raining.
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	// Movement → Condition → Operation → State → Result
	// Movement: left++ or right-- based on height comparison
	// Condition: left < right
	// Operation: calculate trapped water
	// State: update leftMax, rightMax
	// Result: accumulated water

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}
	return water
}
