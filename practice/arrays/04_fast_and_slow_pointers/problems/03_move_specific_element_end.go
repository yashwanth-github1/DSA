package main

// 3️⃣ Move All Specific Elements to End
func moveSpecificToEnd(nums []int, val int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			if slow != fast {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			slow++
		}
	}
}
