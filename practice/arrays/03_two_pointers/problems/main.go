package main

import "fmt"

func main() {

	nums := []int{2, 3, 4, 7, 11, 15}
	target := 10

	fmt.Println("Array:", nums)
	fmt.Println("Target:", target)
	fmt.Println("Output:", twoSumSorted(nums, target))

	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target1 := 10
	fmt.Println(allPairsWithSum(nums1, target1)) // [[1 9] [2 8] [3 7] [4 6]]

	nums2 := []int{1, 3, 4, 7, 10}
	target2 := 15
	fmt.Println(closestPair(nums2, target2))

	nums13 := []int{1, 3, 4, 7, 10}
	target5 := 15
	fmt.Println(closestPair(nums13, target5))

	nums3 := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums3)) // Output: [[-1 -1 2] [-1 0 1]]

	nums4 := []int{1, 2, -2, 0, -1, 1}
	target4 := 2

	fmt.Println(threeSumTarget(nums4, target4))

	nums5 := []int{-1, 2, 1, -4}
	target5 := 1
	fmt.Println(threeSumClosest(nums5, target5)) // Output: 2

	nums6 := []int{1, 2, 3, 4, 5, 6, 7}
	ReversePart(nums6, 2, 5)
	fmt.Println(nums6) // Output: [1, 2, 6, 5, 4, 3, 7]

	nums7 := []int{1, 2, 3, 4, 5}
	ReverseArray(nums7)
	fmt.Println(nums7) // Output: [5, 4, 3, 2, 1]

	nums8 := []int{1, 2, 3, 2, 1}
	fmt.Println(isPalindrome(nums8)) // Output: true

	nums9 := []int{1, 2, 3, 4, 5}
	fmt.Println(isPalindrome(nums9)) // Output: false

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height)) // Output: 49

	height1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height1)) // Output: 6

	nums10 := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums10)
	fmt.Println(nums10) // Output: [0 0 1 1 2 2]

	nums11 := []int{1, 2, 3, 0, 0, 0}
	nums12 := []int{2, 5, 6}
	merge(nums11, 3, nums12, 3)
	fmt.Println(nums11) // Output: [1 2 2 3 5 6]

}
