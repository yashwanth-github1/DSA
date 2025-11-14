1️⃣ Remove Duplicates from Unsorted Array

Problem: Remove all duplicates from an unsorted array and return array of unique elements.

Input: arr = [4, 2, 7, 2, 1, 2, 3, 4]

Output: [4, 2, 7, 1, 3] (order of first occurrence preserved)

Constraints / Requirements:

Array can be unsorted

Handle duplicates

Maintain relative order of first occurrences

Empty array → return empty array

Time Complexity: O(n)

Reasoning: Single pass to check duplicates using a hash set/dictionary.

Space Complexity: O(k)

Reasoning: Need extra space to track k distinct elements.

Algorithm (HashSet / Single Pointer for order preservation):

Initialize empty set seen and pos = 0

Loop i = 0 to n-1:

If arr[i] not in seen:

arr[pos] = arr[i]; pos++

Add arr[i] to seen

Return first pos elements of array as unique array

Example: [4,2,7,2,1,2,3,4] → [4,2,7,1,3]

LeetCode / Reference: Remove Duplicates from Sorted Array II / General
 (adapted for unsorted)

2️⃣ Contains Duplicate

Problem: Check if array contains any duplicates.

Input: arr = [4, 2, 7, 2, 1, 2, 3, 4]

Output: true

Constraints / Requirements:

Array can be unsorted

Return boolean → true if any element repeats, false otherwise

Empty array → return false

Time Complexity: O(n)

Reasoning: Single pass to check and store seen elements in a set.

Space Complexity: O(k)

Reasoning: Extra space for set of k distinct elements.

Algorithm (HashSet):

Initialize empty set seen

Loop through array:

If arr[i] in seen → return true

Else → add arr[i] to seen

Return false

Example:

Input: [4,2,7,2,1,2,3,4] → true

Input: [1,2,3,4] → false

LeetCode / Reference: 217. Contains Duplicate