1️⃣ Linear Search (Basic Search)

Input: arr = [4, 2, 7, 1, 9, 3], target = 7

Output: 2 (index of 7)

Constraints / Requirements:

Array can be unsorted

Works for duplicates

Handle empty array → return -1

Time Complexity: O(n)

Space Complexity: O(1)

Algorithm:

Start from index 0

Compare each element with target

If match → return index

If end of array reached → return -1

LeetCode: Not a direct problem, but basic concept used in many problems (e.g., 704. Binary Search
 — but that’s for sorted arrays)

2️⃣ Count Occurrences of Target

Input: arr = [4, 2, 7, 2, 1, 2, 3], target = 2

Output: 3 (2 appears 3 times)

Constraints / Requirements:

Array can be unsorted

Works for duplicates

Empty array → return 0

Time Complexity: O(n)

Space Complexity: O(1)

Algorithm:

Initialize count = 0

Loop through array

If element == target → count++

Return count

LeetCode: Q1150. Check If a Number Occurs Twice
 (variant for counting)

3️⃣ Find All Indices of Target

Input: arr = [4, 2, 7, 2, 1, 2, 3], target = 2

Output: [1, 3, 5] (indices of 2)

Constraints / Requirements:

Array can be unsorted

Works for duplicates

Empty array → return empty list []

Time Complexity: O(n)

Space Complexity: O(k) — k = number of matches

Algorithm:

Initialize empty list result

Loop through array

If element == target → append index to result

Return result

LeetCode: Can be implemented as helper in 1. Two Sum
 to find all indices

4️⃣ First and Last Occurrence (Linear Search)

Input: arr = [4, 2, 7, 2, 1, 2, 3], target = 2

Output: (1, 5) → first at index 1, last at index 5

Constraints / Requirements:

Array can be unsorted for linear search

Works for duplicates

Empty array → return (-1, -1)

Time Complexity: O(n)

Space Complexity: O(1)

Algorithm:

Initialize first = -1, last = -1

Loop through array

If element == target:

If first == -1, set first = index

Always set last = index

Return (first, last)

LeetCode: 34. Find First and Last Position of Element in Sorted Array
 (for sorted array → binary search version)
