1️⃣ Find Maximum and Minimum Together

Input: arr = [4, 2, 7, 1, 9, 3]

Output: max = 9, min = 1

Constraints / Requirements:

Array can be unsorted

Handle empty array → return (None, None)

Time Complexity: O(n)

Must scan all elements to determine max and min

Space Complexity: O(1)

Only two variables used

Algorithm:

Initialize max = arr[0], min = arr[0]

Loop through array → update max/min

Return (max, min)

Example: [4,2,7,1,9,3] → max=9, min=1

LeetCode / Reference: Maximum Product Subarray
 (uses max/min logic internally)

2️⃣ Find Index of Maximum Element (Variant of Max)

Input: arr = [4, 2, 7, 1, 9, 3]

Output: 4 (index of 9)

Constraints / Requirements:

Array can be unsorted

If multiple maximums → return first occurrence

Empty array → return -1

Time Complexity: O(n)

Space Complexity: O(1)

Algorithm:

Initialize max_val = arr[0], max_index = 0

Loop → update max_val and max_index

Return max_index

Example: [4,2,7,1,9,3] → index 4

LeetCode / Reference: LeetCode 152. Maximum Product Subarray
 (max logic inside)

3️⃣ Find Second Maximum Element

Input: arr = [4, 2, 7, 1, 9, 3]

Output: 7

Constraints / Requirements:

Array can be unsorted

At least 2 distinct elements

Handle duplicates → second largest distinct

Time Complexity: O(n) — single pass to find max1 and max2

Space Complexity: O(1)

Algorithm:

Initialize max1=-∞, max2=-∞

Loop → update max1/max2

Return max2

Example: [4,2,7,1,9,3] → 7

LeetCode / Reference: LeetCode 414. Third Maximum Number
 (similar logic for second/third max)

4️⃣ Find Second Minimum Element (Variant)

Input: arr = [4, 2, 7, 1, 9, 3]

Output: 2

Constraints / Requirements:

Array can be unsorted

At least 2 distinct elements

Handle duplicates

Time Complexity: O(n)

Space Complexity: O(1)

Algorithm:

Initialize min1=+∞, min2=+∞

Loop → update min1/min2

Return min2

Example: [4,2,7,1,9,3] → 2
LeetCode / Reference: LeetCode 414. Third Maximum Number
 (analogous for min)