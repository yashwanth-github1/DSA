package main

//move zeros to end
func moveZeroes(nums []int) {
	n := len(nums)
	lastNonZero := 0

	// Move non-zero elements forward
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[lastNonZero] = nums[i]
			lastNonZero++ // move pointer
		}
	}

	// Fill remaining positions with zeros
	for i := lastNonZero; i < n; i++ {
		nums[i] = 0
	}
}

//move zeros to beginning

func moveZeroesToFront(nums []int) {
	n := len(nums)
	insertPos := n - 1 // position to insert non-zero from end

	// Move non-zero elements backward
	for i := n - 1; i >= 0; i-- {
		if nums[i] != 0 {
			nums[insertPos] = nums[i]
			insertPos--
		}
	}

	// Fill remaining positions with zeros
	for i := 0; i <= insertPos; i++ {
		nums[i] = 0
	}
}

// Move All x to the End

func moveXToEnd(nums []int, x int) {
	n := len(nums)
	last := 0 // pointer for next non-x element

	for i := 0; i < n; i++ {
		if nums[i] != x {
			nums[last] = nums[i]
			last++
		}
	}

	for i := last; i < n; i++ {
		nums[i] = x
	}
}
