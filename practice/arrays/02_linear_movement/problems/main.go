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

	nums3 := []int{5, 2, 2, 4, 2, 7}
	target3 := 2
	fmt.Println("First and Last Occurrence:", searchRange(nums3, target3))

	nums4 := []int{4, 2, 7, 1, 9, 3}
	maxVal, minVal := findMaxMin(nums4)
	fmt.Println("Maximum:", maxVal, "Minimum:", minVal)

	nums5 := []int{4, 2, 7, 1, 9, 3}
	fmt.Println("Index of Maximum:", maxIndex(nums5))

	nums6 := []int{4, 2, 7, 1, 9, 3}
	fmt.Println("Second Maximum:", secondMax(nums6))

	nums7 := []int{4, 2, 7, 1, 9, 3}
	fmt.Println("Second Minimum:", secondMin(nums7))

	nums8 := []int{1, 7, 3, 7, 7, 2}
	fmt.Println("Frequency of elements:", frequencyMap(nums8))

	nums11 := []int{1, 7, 3, 7, 7, 2}
	element, freq := mostfreqElement(nums11)
	fmt.Println("Most Frequent Element:", element, "Frequency:", freq)

	nums9 := []int{1, 7, 3, 7, 7, 2}
	element1, freq1 := leastfreqElement(nums9)
	fmt.Println("Least Frequent Element:", element1, "Frequency:", freq1)

	nums10 := []int{1, 7, 3, 7, 7, 2}
	fmt.Println("Sorted by frequency:", sortByFrequency(nums10))

	nums12 := []int{0, 1, 0, 3, 12}
	moveZeroes(nums12)
	fmt.Println("Array after moving zeros:", nums12)

	a := []int{1, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5}
	duplicatescount, values := countDuplicatesInArray(a)
	fmt.Println("duplicates count", duplicatescount, "and its values", values)

	nums13 := []int{0, 1, 0, 3, 12}
	moveZeroesToFront(nums13)
	fmt.Println("Array after moving zeros to front:", nums13)

	nums14 := []int{2, 1, 2, 3, 2, 4}
	x := 2
	moveXToEnd(nums, x)
	fmt.Println("Array after moving x to end:", nums14)

	nums15 := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	fmt.Println("Array after removing duplicates:", removeDuplicates(nums15))

	nums16 := []int{1, 2, 3, 1}
	nums17 := []int{1, 2, 3, 4}

	fmt.Println("Contains duplicate nums1:", containsDuplicate(nums16)) // true
	fmt.Println("Contains duplicate nums2:", containsDuplicate(nums17)) // false

	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{5, 4, 4, 2, 1}
	arr3 := []int{1, 3, 2, 4}

	fmt.Println("arr1 Ascending?", isSortedAscending(arr1))   // true
	fmt.Println("arr2 Descending?", isSortedDescending(arr2)) // true
	fmt.Println("arr3 Sorted?", isSorted(arr3))               // false

	arr4 := []int{5, 4, 4, 2}
	fmt.Println("arr4 strictly increasing?", isStrictlyIncreasing(arr4))

	arr5 := []int{5, 4, 4, 2}
	fmt.Println("arr5 strictly increasing?", isNonIncreasing(arr5))

	nums18 := []int{3, 2, 4}
	target4 := 6
	fmt.Println("Indices:", twoSumUnsorted(nums18, target4))

	nums19 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target5 := 10
	fmt.Println("Count pair:", countPairs(nums19, target5)) //4

	nums20 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("Majority element:", majorityElementMap(nums20))

	nums21 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("Majority element (Boyer-Moore):", majorityElementBoyerMoore(nums21))

	nums22 := []int{1, 2, 3, 1, 2, 1, 1}
	fmt.Println("Majority > n/3 using map:", majorityElementMap_(nums22))

	nums23 := []int{3, 2, 3, 2, 2, 1, 1}
	fmt.Println("Elements appearing > n/3 times:", majorityElementNBy3(nums23))

	a1 := []int{1, 2, 3, 4, 5}
	a2 := []int{1, 2, 3, 4, 5}

	k := 2

	rotateRight(a1, k)
	fmt.Println("Right Rotate:", a1) // [4 5 1 2 3]

	rotateLeft(a2, k)
	fmt.Println("Left Rotate :", a2) // [3 4 5 1 2]

	a3 := []int{1, 2, 2, 4, 5, 2, 2}
	target7 := 1
	fmt.Println(firstAndLastOccurunceinArray(a3, target7))

}
