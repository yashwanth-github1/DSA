package main

import "sort"

//Count Occurrences of All Elements (Frequency Map)
func frequencyMap(nums []int) map[int]int {
	freq := make(map[int]int)
	for _, val := range nums { // traverse array
		freq[val]++ // increment count in map
	}
	return freq
}

//Find Most Frequent Element
func mostFrequent(nums []int) int {
	freq := make(map[int]int)
	maxFreq := 0
	mostFreq := nums[0]

	for _, val := range nums {
		freq[val]++
		if freq[val] > maxFreq {
			maxFreq = freq[val]
			mostFreq = val
		}
	}
	return mostFreq
}

//Find Least Frequent Element

func leastFrequent(nums []int) int {
	freq := make(map[int]int) // Map to store counts

	// Count frequency of each element
	for _, val := range nums {
		freq[val]++
	}

	// Initialize with first element
	leastFreq := nums[0]
	minFreq := freq[leastFreq]

	// Find element with smallest frequency
	for key, count := range freq {
		if count < minFreq { // If this element occurs less times
			minFreq = count // Update min frequency
			leastFreq = key // Update least frequent element
		}
	}

	return leastFreq
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
