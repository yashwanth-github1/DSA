package main

// remove duplicates unsorted
func removeDuplicates(a []int) []int {
	seen := make(map[int]bool) // track seen elements
	result := []int{}

	for _, val := range a { // traverse array
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

// count duplicates
func countDuplicatesInArray(a []int) (int, map[int]int) {
	count := make(map[int]int)
	for _, v := range a {
		count[v]++
	}
	duplicatescount := 0
	duplicates := make(map[int]int)
	for num, d := range count {
		if d > 1 {
			duplicatescount++
			duplicates[num] = d
		}
	}
	return duplicatescount, duplicates
}
