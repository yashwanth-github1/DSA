package main

// two sum unsorted
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
