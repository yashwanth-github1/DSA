package main

// container with most water

//You are given an array height of n non-negative integers, where each element represents a vertical line on the x-axis.
//•	Find two lines which, together with the x-axis, forms a container that holds the most water.
//•	Return the maximum area of water the container can store.

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	// Movement → Condition → Operation → State → Result
	// Movement: left++ or right-- based on comparison
	// Condition: left < right
	// Operation: compute area, update maxArea, move smaller height pointer
	// State: updated maxArea, pointers moved inward
	// Result: final maxArea

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
