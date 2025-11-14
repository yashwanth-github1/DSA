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

// Return All Indices of Target
func allIndices(nums []int, target int) []int {
	var indices []int
	for i, val := range nums { // start movement
		if val == target { // condition
			indices = append(indices, i) // operation
		}
	} // end movement
	return indices
}

// First and Last Occurrence

func searchRange(nums []int, target int) []int {
	first, last := -1, -1
	for i, val := range nums { // start movement
		if val == target { // condition
			if first == -1 { // first occurrence
				first = i // operation 1
			}
			last = i // operation 2
		}
	} // end movement
	return []int{first, last}
}
