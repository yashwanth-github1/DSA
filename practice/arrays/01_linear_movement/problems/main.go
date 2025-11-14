package main

import "fmt"

func main() {
	nums := []int{4, 2, 7, 1, 9, 3}
	target := 7
	fmt.Println("Index:", linearSearch(nums, target))

	nums1 := []int{2, 3, 2, 4, 2, 5}
	target1 := 2
	fmt.Println("Count:", countOccurrences(nums1, target1))

	nums2 := []int{2, 3, 2, 4, 2, 5}
	target2 := 2
	fmt.Println("All Indices:", allIndices(nums2, target2))

}
