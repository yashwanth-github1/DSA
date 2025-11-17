package main

import "sort"

//three sum                //Three Sum = 0
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // skip duplicates
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// skip duplicates
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

// three sum with target
func threeSumTarget(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums)-2; i++ {

		// Skip duplicates for i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// Skip duplicates for left
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// Skip duplicates for right
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return res
}

//3Sum Closest

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2]

	// Movement → Condition → Operation → State → Result
	// Movement: i loop, left & right pointers
	// Condition: left < right
	// Operation: sum check, move pointers
	// State: updated closestSum
	// Result: final closest sum

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(closestSum-target) {
				closestSum = sum
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else { // exact match
				return sum
			}
		}
	}
	return closestSum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
