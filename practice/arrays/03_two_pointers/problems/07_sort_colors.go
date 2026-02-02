package main

import "fmt"

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
