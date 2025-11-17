package main

// two sum
func twoSumSorted(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	// Movement → Condition → Operation → State → Result
	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{left, right} // Result → Found pair
		} else if sum < target {
			left++ // Movement → increase left
		} else {
			right-- // Movement → decrease right
		}
	}
	return []int{} // Result → not found
}

//Return All Pairs with Given Sum
func allPairsWithSum(nums []int, target int) [][2]int {
	left, right := 0, len(nums)-1
	var result [][2]int

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			result = append(result, [2]int{nums[left], nums[right]})
			left++
			right--
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return result
}

//Count Pairs With Given Sum

func countPairs(nums []int, target int) int {
	left, right := 0, len(nums)-1
	count := 0

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			count++
			left++
			right--
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return count
}

//Closest Sum to Target
func closestPair(nums []int, target int) (int, int) {
	left, right := 0, len(nums)-1
	bestL, bestR := left, right
	minDiff := 1 << 31 // very large

	for left < right {
		sum := nums[left] + nums[right]
		diff := sum - target
		if diff < 0 {
			diff = -diff
		}

		if diff < minDiff {
			minDiff = diff
			bestL, bestR = left, right
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}
	return nums[bestL], nums[bestR]
}

//Two Sum in Sorted + Duplicates Allowed
//Return all unique pairs (no duplicates in result) even if array has duplicates.

// func uniquePairs(nums []int, target int) [][2]int {
// 	left, right := 0, len(nums)-1
// 	var res [][2]int

// 	for left < right {
// 		sum := nums[left] + nums[right]
// 		if sum == target {
// 			res = append(res, [2]int{nums[left], nums[right]})
// 			leftVal, rightVal := nums[left], nums[right]
// 			for left < right && nums[left] == leftVal {
// 				left++
// 			}
// 			for left < right && nums[right] == rightVal {
// 				right--
// 			}
// 		} else if sum < target {
// 			left++
// 		} else {
// 			right--
// 		}
// 	}
// 	return res
// }
