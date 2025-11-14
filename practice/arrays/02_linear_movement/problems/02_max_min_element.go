package main

import "math"

func findMaxMin(nums []int) (int, int) {
	maxVal, minVal := nums[0], nums[0]
	for _, val := range nums[1:] { // start movement
		if val > maxVal { // condition 1
			maxVal = val // operation
		}
		if val < minVal { // condition 2
			minVal = val // operation
		}
	} // end movement
	return maxVal, minVal
}

//Find Index of Maximum Element (Variant of Max)

func maxIndex(nums []int) int {
	maxVal := nums[0]
	maxIdx := 0
	for i, val := range nums[1:] { // start movement
		if val > maxVal { // condition
			maxVal = val   // operation
			maxIdx = i + 1 // adjust index for slice offset
		}
	} // end movement
	return maxIdx
}

//Find Second Maximum Element

func secondMax(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	maxVal := math.MinInt64
	secondMax := math.MinInt64

	for _, val := range nums { // start movement
		if val > maxVal { // condition 1
			secondMax = maxVal // operation 1
			maxVal = val       // operation 2
		} else if val > secondMax && val != maxVal { // condition 2
			secondMax = val // operation
		}
	} // end movement

	if secondMax == math.MinInt64 {
		return -1
	}
	return secondMax
}

//Find Second Minimum Element (Variant)

func secondMin(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	minVal := math.MaxInt64
	secondMin := math.MaxInt64

	for _, val := range nums { // start movement
		if val < minVal { // condition 1
			secondMin = minVal // operation 1
			minVal = val       // operation 2
		} else if val < secondMin && val != minVal { // condition 2
			secondMin = val // operation
		}
	} // end movement

	if secondMin == math.MaxInt64 {
		return -1
	}
	return secondMin
}
