package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length := removeDuplicates(nums)
	fmt.Println("New length:", length)           // Output: 5
	fmt.Println("Modified nums:", nums[:length]) // Output: [0 1 2 3 4]

	nums1 := []int{0, 1, 0, 3, 12}
	moveZerosToEnd(nums1)
	fmt.Println("Move zeros to end:", nums1) // [1 3 12 0 0]

	nums2 := []int{1, 0, 2, 0, 3}
	moveZerosToStart(nums2)
	fmt.Println("Move zeros to start:", nums2) // [0 0 1 2 3]

	nums3 := []int{2, 1, 2, 3, 2, 4}
	moveSpecificToEnd(nums3, 2)
	fmt.Println("Move 2 to end:", nums3) // [1 3 4 2 2 2]

	nums4 := []int{3, 2, 2, 3, 4}
	length1 := removeElement(nums4, 3)
	fmt.Println("Remove 3 - new length:", length1) // 3
	fmt.Println("Modified nums:", nums4[:length1]) // [2 2 4]

}
