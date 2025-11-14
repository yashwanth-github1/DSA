Linear Search Notes
1️⃣ Linear Search (Basic Search)

Input: arr = [4, 2, 7, 1, 9, 3], target = 7

Output: 2 (index of 7)

Constraints / Requirements:

Array can be unsorted

Works for duplicates

Handle empty array → return -1

Time Complexity: O(n)

Reasoning: Must scan all n elements in worst case (target at end or not present).

Best case: O(1) if target is at index 0.

Space Complexity: O(1)

Reasoning: Only a few variables used, no extra memory required.

Algorithm:

Start from index 0

Compare each element with target

If match → return index

If end of array reached → return -1

LeetCode / Reference: Basic concept used in many problems; no direct LeetCode problem.

2️⃣ Count Occurrences of Target

Input: arr = [4, 2, 7, 2, 1, 2, 3], target = 2

Output: 3

Constraints / Requirements:

Array can be unsorted

Works for duplicates

Empty array → return 0

Time Complexity: O(n)

Reasoning: Must check each element to count occurrences.

Space Complexity: O(1)

Reasoning: Only a counter variable is used.

Algorithm:

Initialize count = 0

Loop through array

If element == target → count++

Return count

LeetCode / Reference: Check If a Number Occurs Twice
 (variant)

3️⃣ Find All Indices of Target

Input: arr = [4, 2, 7, 2, 1, 2, 3], target = 2

Output: [1, 3, 5]

Constraints / Requirements:

Array can be unsorted

Works for duplicates

Empty array → return empty list []

Time Complexity: O(n)

Reasoning: Must scan all elements to find all indices.

Space Complexity: O(k)

Reasoning: Need to store k indices, where k = number of matches.

Algorithm:

Initialize empty list result

Loop through array

If element == target → append index to result

Return result

LeetCode / Reference: Can be implemented as helper in Two Sum

4️⃣ First and Last Occurrence (Linear Search)

Input: arr = [4, 2, 7, 2, 1, 2, 3], target = 2

Output: (1, 5) → first at index 1, last at index 5

Constraints / Requirements:

Array can be unsorted (for linear search)

Works for duplicates

Empty array → return (-1, -1)

Time Complexity: O(n)

Reasoning: Must scan all elements to find first and last occurrence.

Space Complexity: O(1)

Reasoning: Only two variables (first and last) used.

Algorithm:

Initialize first = -1, last = -1

Loop through array

If element == target:

If first == -1, set first = index

Always set last = index

Return (first, last)

LeetCode / Reference: 34. Find First and Last Position of Element in Sorted Array
 (binary search version)
