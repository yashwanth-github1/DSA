1️⃣ Majority Element (Appears > n/2 times)

Problem: Find element that appears more than n/2 times in the array.

Input: arr = [3, 3, 4, 2, 3, 3, 3]

Output: 3

Constraints / Requirements:

Array can be unsorted

Exactly one element appears > n/2 times

Empty array → no majority → return None

Works for integers

Time & Space Complexity Reasoning:

HashMap Approach:

Time Complexity: O(n) → iterate array + count frequency

Space Complexity: O(k) → store distinct elements

Boyer–Moore Voting Algorithm:

Time Complexity: O(n) → single pass

Space Complexity: O(1) → only candidate & count stored

Algorithm:

HashMap Method:

Initialize empty map freq

Loop through array → freq[arr[i]]++

If freq[arr[i]] > n/2 → return element

Boyer–Moore Voting Algorithm:

Initialize candidate = None, count = 0

Loop through array:

If count == 0 → candidate = arr[i]

If arr[i] == candidate → count++

Else → count--

Return candidate

LeetCode / Reference: 169. Majority Element

2️⃣ Majority Element (> n/3 times)

Problem: Find all elements that appear more than n/3 times in the array.

Input: arr = [3,2,3]

Output: [3]

Constraints / Requirements:

Array can be unsorted

Maximum 2 elements can appear > n/3 times

Return list of elements

Empty array → return empty list

Time & Space Complexity Reasoning:

HashMap Approach:

Time Complexity: O(n) → iterate + count

Space Complexity: O(k) → map of distinct elements

Boyer–Moore Voting Algorithm (Generalized):

Time Complexity: O(n) → two passes: candidate selection & verification

Space Complexity: O(1) → at most 2 candidates

Algorithm:

HashMap Method:

Count frequencies in map

Loop map → if freq[element] > n/3 → add to result

Boyer–Moore Voting (Generalized for n/3):

Initialize candidate1 = None, candidate2 = None, count1 = 0, count2 = 0

First pass: select candidates

If arr[i] == candidate1 → count1++

Else if arr[i] == candidate2 → count2++

Else if count1 == 0 → candidate1 = arr[i]; count1 = 1

Else if count2 == 0 → candidate2 = arr[i]; count2 = 1

Else → count1--, count2--

Second pass: verify counts → add candidates with > n/3 occurrences

Example:
arr = [1,1,1,3,3,2,2,2] → n=8 → n/3=2.66 → elements with count>2 → [1,2]

LeetCode / Reference: 229. Majority Element II