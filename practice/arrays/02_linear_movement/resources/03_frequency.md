1️⃣ Count Frequency of All Elements

Input: arr = [4, 2, 7, 2, 1, 2, 3, 4]

Output: {4:2, 2:3, 7:1, 1:1, 3:1}

Constraints / Requirements:

Array can be unsorted

Elements can be repeated

Empty array → return empty map/dictionary

Time Complexity: O(n)

Must iterate through all n elements

Space Complexity: O(k)

Extra space for k distinct elements

Algorithm:

Initialize empty dictionary/map freq

Loop through array:

If element exists → increment count

Else → initialize count as 1

Return freq

LeetCode / Reference: 347. Top K Frequent Elements

2️⃣ Most Frequent Element

Input: arr = [4, 2, 7, 2, 1, 2, 3, 4]

Output: 2 (frequency = 3)

Constraints / Requirements:

Array can be unsorted

Elements can be repeated

Empty array → return None or sentinel

Time Complexity: O(n)

Count all frequencies → O(n), then find max → O(k) ≈ O(n)

Space Complexity: O(k)

Store frequency map for k distinct elements

Algorithm:

Count frequencies → freq map

Initialize max_freq = 0, most_frequent = None

Loop through freq map → update if frequency > max_freq

Return most_frequent

LeetCode / Reference: 347. Top K Frequent Elements

3️⃣ Least Frequent Element

Input: arr = [4, 2, 7, 2, 1, 2, 3, 4]

Output: 7 (frequency = 1)

Constraints / Requirements:

Array can be unsorted

Handle duplicates

Empty array → return None

Time Complexity: O(n)

Count frequencies → O(n), then find min → O(k) ≈ O(n)

Space Complexity: O(k)

Algorithm:

Count frequencies → freq map

Initialize min_freq = ∞, least_frequent = None

Loop through freq map → update if frequency < min_freq

Return least_frequent

LeetCode / Reference: Variant of 347. Top K Frequent Elements

4️⃣ Sort Elements by Frequency

Input: arr = [4, 2, 7, 2, 1, 2, 3, 4]

Output: [2, 4, 7, 1, 3] (sorted by descending frequency; tie can be by order of first appearance)

Constraints / Requirements:

Array can be unsorted

Handle duplicates

Empty array → return empty list

Time Complexity: O(n log k)

Count frequencies → O(n), sort k distinct elements → O(k log k) ≈ O(n log k)

Space Complexity: O(k)

Store frequency map and list of distinct elements

Algorithm:

Count frequencies → freq map

Create list of distinct elements

Sort list based on frequency descending

Return sorted list

LeetCode / Reference: 451. Sort Characters By Frequency

💡 Notes / Tips