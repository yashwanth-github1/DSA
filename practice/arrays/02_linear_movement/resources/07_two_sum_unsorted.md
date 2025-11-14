1️⃣ Two Sum – Unsorted Array
Problem:

Find two indices whose values add up to the target.

Input:

arr = [4, 2, 7, 1, 9], target = 10

Output:

[0, 2]
(Because 4 + 7 = 11—example adjust—so correct example:)
Correct example:
arr = [4, 2, 6, 1, 9], target = 10 → Output: [0, 2]

Constraints / Requirements:

Array can be unsorted

Need indices, not values

Exactly one valid pair is expected (LeetCode version)

No need to handle duplicates separately

If no pair exists → return empty/[-1, -1] based on your code style

Time Complexity: O(n)

Reasoning:

We check each element once and use a hash map for O(1) lookups.

Space Complexity: O(n)

Reasoning:

Hash map stores at most all elements.

Algorithm (HashMap – Single Pass):

Create empty map m (value → index)

Loop i from 0 to n−1:

Compute need = target - arr[i]

If need exists in map → return [map[need], i]

Else store arr[i] in map

If no pair found → return no answer

LeetCode Reference:

1. Two Sum

2️⃣ Count Pairs With Given Sum (Unsorted Array)

(Not the same as Two Sum: here we count number of valid pairs, not indices)

Problem:

Count the number of unique pairs (i < j) such that:
arr[i] + arr[j] = target

Input:

arr = [1, 5, 7, 1], target = 6

Output:

2
Pairs: (1,5) at indices (0,1) and (1,3)

Constraints / Requirements:

Array can be unsorted

Elements may be repeated

Count all valid pairs (i < j)

Order does not matter

Do NOT return indices

Must handle duplicates correctly

Time Complexity: O(n)

Reasoning:

One pass using hashing

Frequency map helps count complement occurrences quickly

Space Complexity: O(n)

Reasoning:

We store element frequencies in a map.

Algorithm (HashMap Frequency Method):

Create empty map freq

Set count = 0

Loop through array:

Compute need = target - arr[i]

If need already in freq:

Add freq[need] to count

Increase freq[arr[i]]

Return count

Example:

arr = [1, 5, 7, 1], target = 6

When first 1 → freq = {1:1}

When 5 → need=1 → count=1

When 7 → need=-1 → skip

When second 1 → need=5 → count=2

Total pairs = 2

LeetCode Reference:

Closest problem:
LeetCode 1679 – Max Number of K-Sum Pairs
(Counts pairs, but uses sorting + two pointers)