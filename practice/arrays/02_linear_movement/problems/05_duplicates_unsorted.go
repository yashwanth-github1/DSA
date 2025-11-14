package main

// remove duplicates unsorted
func removeDuplicates(nums []int) []int {
	seen := make(map[int]bool) // track seen elements
	result := []int{}

	for _, val := range nums { // traverse array
		if !seen[val] { // check condition
			result = append(result, val) // operation
			seen[val] = true             // mark as seen
		}
	}

	return result
}

//Contains Duplicate

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool) // track seen elements

	for _, val := range nums { // traverse array
		if seen[val] { // condition: already seen
			return true // duplicate found
		}
		seen[val] = true // mark as seen
	}

	return false // no duplicates
}
