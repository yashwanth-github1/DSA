package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 20, 30, 40, 50}
	fmt.Println("Original Array:", arr)

	// 1. Access / Read
	fmt.Println("Access index 2:", arr[2])

	// 2. Update / Modify
	arr[1] = 25
	fmt.Println("After Update index 1:", arr)

	// 4. Insertion at index 2, value 15
	arr = insertAt(arr, 2, 15)
	fmt.Println("After Insertion at index 2:", arr)

	// 5. Deletion at index 3
	arr = deleteAt(arr, 3)
	fmt.Println("After Deletion at index 3:", arr)

	// 6. Traversal
	fmt.Print("Traversal: ")
	traverse(arr)

	// 8. Sum / Average
	sum := findSum(arr)
	fmt.Println("Sum:", sum, "Average:", float64(sum)/float64(len(arr)))

	// 9. Reverse
	reverse(arr)
	fmt.Println("After Reverse:", arr)

	// 10. Copy / Clone
	arrCopy := copyArray(arr)
	fmt.Println("Copied Array:", arrCopy)

}

// Insert at index
func insertAt(arr []int, index int, val int) []int {
	arr = append(arr, 0) // increase slice size
	for i := len(arr) - 1; i > index; i-- {
		arr[i] = arr[i-1]
	}
	arr[index] = val
	return arr
}

// Delete at index
func deleteAt(arr []int, index int) []int {
	for i := index; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	return arr[:len(arr)-1]
}

// Traversal
func traverse(arr []int) {
	for _, val := range arr {
		fmt.Print(val, " ")
	}
}

// Sum
func findSum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

// Reverse
func reverse(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// Copy / Clone
func copyArray(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}
