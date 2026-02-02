package main

import "math"

func findMaxMin(nums []int) (int, int) {
	maxVal, minVal := nums[0], nums[0]
	for _, val := range nums[1:] {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
	}
	return maxVal, minVal
}

//Find Index of Maximum Element (Variant of Max)

func maxIndex(nums []int) int {
	maxVal := nums[0]
	maxIdx := 0
	for i, val := range nums[1:] {
		if val > maxVal {
			maxVal = val
			maxIdx = i + 1 // adjust index for slice offset
		}
	}
	return maxIdx
}

//Find Second Maximum Element

func secondMax(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	maxVal := math.MinInt64
	secondMax := math.MinInt64

	for _, val := range nums {
		if val > maxVal {
			secondMax = maxVal
			maxVal = val
		} else if val > secondMax && val != maxVal {
			secondMax = val
		}
	}

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

	for _, val := range nums {
		if val < minVal {
			secondMin = minVal
			minVal = val
		} else if val < secondMin && val != minVal {
			secondMin = val
		}
	}

	if secondMin == math.MaxInt64 {
		return -1
	}
	return secondMin
}
