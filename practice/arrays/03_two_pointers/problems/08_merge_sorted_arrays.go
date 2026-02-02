package main

//Given two sorted arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
//•	You may assume nums1 has enough space (size m + n) to hold additional elements from nums2.
//•	Modify nums1 in-place.

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1 // Movement: three pointers at array ends

	for p1 >= 0 && p2 >= 0 { // Condition: both pointers valid
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1] //place larger element at end
			p1--                 // State: move p1
		} else {
			nums1[p] = nums2[p2] // place larger element at end
			p2--                 // State: move p2
		}
		p-- // merged pointer moves backward
	}

	// Copy remaining nums2 if any
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}

func mergeSortedArrays(a []int, b []int) []int {
	i, j := 0, 0
	result := []int{}
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	// for i<len(a){
	// result = append(result,a[i])
	// }
	// for j<len(b){
	// result = append(result,b[j])
	// }
	//
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}
