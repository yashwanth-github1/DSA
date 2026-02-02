package main

// two sum unsorted(return indices)
func twoSumUnsorted(nums []int, target int) []int { //[2]int
	seen := make(map[int]int) // map[num] = index

	for i, num := range nums { // traverse array once
		complement := target - num
		if idx, found := seen[complement]; found { // check if complement exists
			return []int{idx, i} // return indices of the two numbers    //return [2]int{idx, i}
		}
		seen[num] = i // store current number and its index
	}

	return []int{} // if no pair found  //[2]int{-1, -1}
}

// two sum unsorted(return values)

func twoSumUnsortedValues(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index (unchanged)

	for _, num := range nums {
		complement := target - num
		if _, found := seen[complement]; found {
			return []int{complement, num} // return VALUES instead of indices
		}
		seen[num] = 1 // value doesn't matter anymore
	}
	return []int{}
}

//Count Pairs With Given Sum which equals target
func countPairs(nums []int, target int) int {
	freq := make(map[int]int)
	count := 0

	for _, num := range nums {
		complement := target - num
		if c, exists := freq[complement]; exists {
			count += c
		}
		freq[num]++
	}

	return count
}

func twoSumUnsortedBool(a []int, target int) []int {
	freq := make(map[int]bool)

	for _, v := range a {
		complement := target - v
		if freq[complement] {
			return []int{complement, v}
		}
		freq[v] = true
	}

	return nil
}
