package main

import "sort"

//Count Occurrences of All Elements (Frequency Map)
func frequencyMap(a []int) map[int]int {
	freq := make(map[int]int)
	for _, val := range a { // traverse array
		freq[val]++ // increment count in map
	}
	return freq
}

//Find Most Frequent Element
func mostfreqElement(a []int) (int, int) {
	freq := make(map[int]int)
	for _, v := range a {
		freq[v]++
	}
	mostFreq := -1
	element := -1
	for num, val := range freq {
		if val > mostFreq {
			mostFreq = val
			element = num
		}
	}
	return element, mostFreq
}

//Find Least Frequent Element

func leastfreqElement(a []int) (int, int) {
	freq := make(map[int]int)
	for _, v := range a {
		freq[v]++
	}
	leastFreq := a[0] //leastFreq := math.MaxInt64
	element := a[0]
	for num, val := range freq {
		if val < leastFreq {
			leastFreq = val
			element = num
		}
	}
	return element, leastFreq
}

//Sort Elements by Frequency

func sortByFrequency(nums []int) []int {
	freq := make(map[int]int)

	// Count frequency
	for _, val := range nums {
		freq[val]++
	}

	// Get unique numbers
	keys := []int{}
	for key := range freq {
		keys = append(keys, key)
	}

	// Sort keys based on frequency (high to low)
	sort.Slice(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})

	// Build result array
	result := []int{}
	for _, key := range keys {
		for i := 0; i < freq[key]; i++ {
			result = append(result, key)
		}
	}

	return result
}
