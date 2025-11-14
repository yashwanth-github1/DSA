1️⃣ Move Zeros to End (Single Pointer)

Problem: Move all 0s to the end while keeping the relative order of non-zero elements.

Input: arr = [0, 1, 0, 3, 12]

Output: [1, 3, 12, 0, 0]

Constraints / Requirements:

Array can be unsorted

Maintain relative order of non-zero elements

In-place using single pointer

Empty array → return empty array

Time Complexity: O(n)

Must scan all n elements once.

Space Complexity: O(1)

Only single pointer variable used.

Algorithm (Single Pointer):

Initialize pos = 0 → next index to place non-zero

Loop i = 0 to n-1:

If arr[i] != 0:

arr[pos] = arr[i]

If i != pos → arr[i] = 0

pos++

Return arr

LeetCode Reference: 283. Move Zeroes

2️⃣ Move Zeros to Front (Single Pointer)

Problem: Move all 0s to the front while maintaining relative order of non-zero elements.

Input: arr = [0, 1, 0, 3, 12]

Output: [0, 0, 1, 3, 12]

Constraints / Requirements:

Array can be unsorted

Maintain relative order of non-zero elements

In-place using single pointer

Empty array → return empty array

Time Complexity: O(n)

Scan all elements once.

Space Complexity: O(1)

Algorithm (Single Pointer from End):

Initialize pos = n-1 → next index from end for non-zero

Loop i = n-1 to 0:

If arr[i] != 0:

arr[pos] = arr[i]

If i != pos → arr[i] = 0

pos--

Return arr

LeetCode Reference: Reverse logic of 283. Move Zeroes

3️⃣ Move All X to End (Single Pointer)

Problem: Move all occurrences of x to the end while preserving order of other elements.

Input: arr = [2, 1, 2, 3, 2, 4], x = 2

Output: [1, 3, 4, 2, 2, 2]

Constraints / Requirements:

Array can be unsorted

Maintain relative order of non-x elements

In-place using single pointer

Empty array → return empty array

Time Complexity: O(n)

Scan all elements once.

Space Complexity: O(1)

Algorithm (Single Pointer):

Initialize pos = 0 → next index to place non-x

Loop i = 0 to n-1:

If arr[i] != x:

arr[pos] = arr[i]

If i != pos → arr[i] = x

pos++

Return arr

LeetCode Reference: Generalization of 283. Move Zeroes