package main

import "fmt"

//Given an array nums with n objects colored 0, 1, or 2, sort them in-place so that objects of the same color are adjacent, with the colors in the order 0, 1, 2.
//•	Do not use the library’s sort function.
//•	Solve it in one pass using constant extra space.

func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		default:
			// Optional: handle unexpected values
			fmt.Println("Skipping unexpected value:", nums[mid])
			mid++
		}
	}
}
