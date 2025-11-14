package main

//Find Majority Element
//Given an array nums, find the element that appears more than n/2 times.

func majorityElementMap(nums []int) int {
	freq := make(map[int]int)
	n := len(nums)

	for _, val := range nums {
		// Movement → iterate over array
		freq[val]++
		// Operation → increment count
		if freq[val] > n/2 {
			// Condition → count exceeds n/2
			return val
			// Result → majority element found
		}
	}

	return -1 // theoretically won't reach here
}

//for n/3

func majorityElementMap_(nums []int) []int {
	freq := make(map[int]int)

	// Step 1: Count frequencies
	for _, val := range nums {
		freq[val]++
	}

	// Step 2: Collect elements with freq > n/3
	result := []int{}
	limit := len(nums) / 3

	for key, count := range freq {
		if count > limit {
			result = append(result, key)
		}
	}

	return result
}

//Optimized Approach 2: Boyer-Moore Voting Algorithm (O(1) space)  (n/2)
func majorityElementBoyerMoore(nums []int) int {
	count := 0
	var candidate int

	for _, val := range nums {
		// Movement → iterate array
		if count == 0 {
			candidate = val
			// Operation → pick new candidate
			count = 1
			// State → candidate count reset
		} else if candidate == val {
			count++
			// Condition → matches candidate → increment count
		} else {
			count--
			// Condition → different → decrement count
		}
		// State → candidate and count updated
	}

	// Result → candidate is majority element
	return candidate
}

//Optimized Approach (Boyer–Moore Voting Algorithm for n/3 case)
func majorityElementNBy3(nums []int) []int {
	// Step 1: Find two possible candidates using voting method
	var cand1, cand2, count1, count2 int

	for _, num := range nums {
		switch {
		case num == cand1:
			count1++
		case num == cand2:
			count2++
		case count1 == 0:
			cand1, count1 = num, 1
		case count2 == 0:
			cand2, count2 = num, 1
		default:
			count1--
			count2--
		}
	}

	// Step 2: Verify both candidates
	count1, count2 = 0, 0
	for _, num := range nums {
		if num == cand1 {
			count1++
		} else if num == cand2 {
			count2++
		}
	}

	// Step 3: Return result
	n := len(nums)
	result := []int{}
	if count1 > n/3 {
		result = append(result, cand1)
	}
	if cand2 != cand1 && count2 > n/3 {
		result = append(result, cand2)
	}

	return result
}
