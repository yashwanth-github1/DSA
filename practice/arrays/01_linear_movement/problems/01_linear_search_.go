package main

func linearSearch(nums []int, target int) int {
	for i, val := range nums { // start movement i=0
		if val == target { // condition
			return i // operation
		}
	} // end movement i=len(nums)
	return -1
}

// count occurences of target
func countOccurrences(nums []int, target int) int {
	count := 0
	for _, val := range nums { // start movement
		if val == target { // condition
			count++ // operation
		}
	} // end movement
	return count
}

func allIndices(nums []int, target int) []int {
	var indices []int
	for i, val := range nums { // start movement
		if val == target { // condition
			indices = append(indices, i) // operation
		}
	} // end movement
	return indices
}
