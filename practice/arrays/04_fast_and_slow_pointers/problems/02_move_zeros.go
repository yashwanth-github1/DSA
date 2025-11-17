package main

//Move Zeroes to end

func moveZerosToEnd(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			if slow != fast {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			slow++
		}
	}
}

//  Move Zeros to Start
func moveZerosToStart(nums []int) {
	slow := len(nums) - 1
	for fast := len(nums) - 1; fast >= 0; fast-- {
		if nums[fast] != 0 {
			if slow != fast {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			slow--
		}
	}
}
