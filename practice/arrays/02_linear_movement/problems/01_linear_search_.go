package main

func linearSearch(nums []int, target int) int {
	for i, val := range nums {
		if val == target {
			return i
		}
	}
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
	for i, val := range nums {
		if val == target {
			indices = append(indices, i)
		}
	}
	return indices
}

// First and Last Occurrence

func searchRange(nums []int, target int) []int {
	first, last := -1, -1
	for i, val := range nums {
		if val == target {
			if first == -1 {
				first = i
			}
			last = i
		}
	}
	return []int{first, last}
}
