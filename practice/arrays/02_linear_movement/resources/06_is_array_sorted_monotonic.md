Array Sorted / Monotonic Checks – Notes
1️⃣ Check if Array is Sorted in Descending Order

Problem: Determine if array is sorted in non-increasing (descending) order.

Input: arr = [9, 7, 5, 3, 1]

Output: true

Constraints / Requirements:

Array can be unsorted

Handle empty array → consider sorted → true

Single element → sorted → true

Time Complexity: O(n)

Reasoning: Need to check every adjacent pair.

Space Complexity: O(1)

Algorithm:

Loop i = 1 to n-1:

If arr[i] > arr[i-1] → return false

Return true

Example: [9,7,5,3,1] → true, [9,5,7] → false

LeetCode / Reference: 896. Monotonic Array

2️⃣ Check if Array is Sorted in Ascending Order

Problem: Determine if array is sorted in non-decreasing (ascending) order.

Input: arr = [1, 2, 3, 5, 7]

Output: true

Constraints / Requirements:

Array can be unsorted

Empty or single-element array → sorted → true

Time Complexity: O(n)

Space Complexity: O(1)

Algorithm:

Loop i = 1 to n-1:

If arr[i] < arr[i-1] → return false

Return true

Example: [1,2,3,5,7] → true, [1,3,2] → false

LeetCode / Reference: 896. Monotonic Array

3️⃣ Check if Array is Strictly Increasing

Problem: Determine if array is strictly increasing → each element greater than previous.

Input: arr = [1, 2, 3, 4]

Output: true

Constraints / Requirements:

Array can be unsorted

Must strictly increase → no equal adjacent elements

Empty or single-element → return true

Time Complexity: O(n)

Space Complexity: O(1)

Algorithm:

Loop i = 1 to n-1:

If arr[i] <= arr[i-1] → return false

Return true

Example: [1,2,3,4] → true, [1,2,2,3] → false

LeetCode / Reference: Can be considered variant of 896. Monotonic Array

4️⃣ Check if Array is Non-Increasing

Problem: Determine if array is non-increasing → each element less than or equal to previous.

Input: arr = [9, 7, 7, 5, 3]

Output: true

Constraints / Requirements:

Array can be unsorted

Duplicates allowed

Empty or single-element → return true

Time Complexity: O(n)

Space Complexity: O(1)

Algorithm:

Loop i = 1 to n-1:

If arr[i] > arr[i-1] → return false

Return true

Example: [9,7,7,5,3] → true, [9,7,8,5] → false

LeetCode / Reference: 896. Monotonic Array