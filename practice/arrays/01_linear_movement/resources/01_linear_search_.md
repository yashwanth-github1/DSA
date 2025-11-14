linear search 

1️⃣ Problem Statement

Find the index of first occurrence of target in an array nums. Return -1 if not found.

2️⃣ Input / Output
nums = [4, 2, 7, 1, 9, 3]
target = 7
Output: 2

3️⃣ LeetCode Reference

NA

4️⃣ Brute Force Pseudocode
function linearSearch(nums, target):
    for i = 0 to len(nums)-1:
        if nums[i] == target:
            return i
    return -1


Explanation:

Traverse each element linearly from start to end

Stop immediately on first match → return index

Time Complexity Reasoning:

Best case: target at index 0 → O(1)

Worst case: target at last index or not present → O(n)

Average case: target in middle → O(n)

Space Complexity Reasoning:

Only loop variable used → O(1)

6️⃣ Movement → Condition → Operation → State → Result
Step	Start	End	Condition	Operation	State	Result
Linear Search	i=0	i=n-1	nums[i] == target	return i	current index	index of first occurrence or -1
7️⃣ Dry Run
nums = [4,2,7,1,9,3], target=7
i=0 → 4!=7 → continue
i=1 → 2!=7 → continue
i=2 → 7==7 → return 2
Output: 2


// Problem: Linear Search (Search Element in Array)
// Concept: Arrays
// Pattern: Brute Force / Basic Searching
// Link: https://www.geeksforgeeks.org/linear-search/
// Approach: Iterate through array and compare each element with target
// Time Complexity: O(n)
// Space Complexity: O(1)


Linear Search – Algorithm

Start from the first element of the array.

Compare the current element with the target value.

If they are equal,
→ Return the index (target found).

If not,
→ Move to the next element.

Continue until:

You find the target, or

You reach the end of the array.

If the entire array is scanned and the target is not found,
→ Return -1 (target not present).


-----------------------------------------------------------------------------------

Count all occurrences of target in nums.

1️⃣ Problem Statement
Count all occurrences of target in nums.
2️⃣ Input / Output
nums = [2,3,2,4,2,5]
target = 2
Output: 3
3️⃣ LeetCode Reference
NA
________________________________________
4️⃣ Brute Force Pseudocode
count = 0
for i = 0 to len(nums)-1:
    if nums[i] == target:
        count++
return count
Time Complexity Reasoning:
•	Must check all elements to count occurrences → O(n)
•	Early exit not possible
Space Complexity Reasoning:
•	Only count variable used → O(1)
________________________________________
6️⃣ Movement → Condition → Operation → State → Result
Step	Start	End	Condition	Operation	State	Result
Count Occurrences	i=0	i=n-1	nums[i] == target	count++	running count	total occurrences
7️⃣ Dry Run
nums=[2,3,2,4,2,5], target=2
i=0 → 2==2 → count=1
i=1 → 3!=2 → count=1
i=2 → 2==2 → count=2
i=3 → 4!=2 → count=2
i=4 → 2==2 → count=3
i=5 → 5!=2 → count=3
Output: 3

short notes



// Problem: Count Occurrences of a Target in Array
// Concept: Arrays
// Pattern: Brute Force / Frequency Counting
// Link: https://www.geeksforgeeks.org/count-number-of-occurrences/
// Approach: Iterate and count matches
// Time: O(n)
// Space: O(1)

Algorithm: Count Occurrences

Initialize a variable count = 0.

Loop from index 0 to n-1 through the array.

For each element:

If arr[i] == target,
→ increment count by 1.

After completing the loop:

Return count.


----------------------------------------------------------------------------------

1️⃣ Problem Statement
Return a slice of all indices where target occurs.
2️⃣ Input / Output
nums = [2,3,2,4,2,5]
target = 2
Output: [0,2,4]
3️⃣ LeetCode Reference
NA
________________________________________
4️⃣ Brute Force Pseudocode
indices = []
for i=0 to len(nums)-1:
    if nums[i]==target:
        append i to indices
return indices
Time Complexity Reasoning:
•	Must scan entire array → O(n)
•	Appending index → O(1) amortized → total still O(n)
Space Complexity Reasoning:
•	Slice stores k occurrences → O(k)
•	Loop variable → negligible → O(1)


6️⃣ Movement → Condition → Operation → State → Result
Step	Start	End	Condition	Operation	State	Result
Return all indices	i=0	i=n-1	nums[i] == target	append i to slice	slice grows	slice of all indices
________________________________________
7️⃣ Dry Run
nums=[2,3,2,4,2,5], target=2
i=0 → 2==2 → indices=[0]
i=1 → 3!=2 → indices=[0]
i=2 → 2==2 → indices=[0,2]
i=3 → 4!=2 → indices=[0,2]
i=4 → 2==2 → indices=[0,2,4]
i=5 → 5!=2 → indices=[0,2,4]
Output: [0,2,4]


short notes 

// Problem: Find All Indices of a Target
// Concept: Arrays
// Pattern: Brute Force
// Link: https://www.geeksforgeeks.org/find-all-occurrences-of-an-element-in-an-array/
// Approach: Traverse array and store positions where arr[i] == target
// Time: O(n)
// Space: O(k)   // k = number of matches



✅ Algorithm: Find All Indices
Create an empty list/array result.

Loop through the array from i = 0 to n-1.

For each element:

If arr[i] == target,
→ append i to result.

After the loop ends, return result.

If no indices found, return an empty list.

------------------------------------------------------------------------------------------



