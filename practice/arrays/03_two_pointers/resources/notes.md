Trapping rain water 

Problem Statement

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much rain water it can trap after raining.

🧠 Explanation (intuition)

Each bar represents a wall height.
Water can be trapped between taller bars, depending on the minimum of the tallest wall on the left and right.

📥 Input

An integer array height of length n

height[i] represents the height of the bar at index i

📤 Output

Return a single integer: total units of water trapped

🧪 Example 1
Input:  height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6


Explanation:

        █
    █   █ █
 █  █ █ █ █ █
----------------
Water trapped = 6 units

🧪 Example 2
Input:  height = [4,2,0,3,2,5]
Output: 9

🔒 Constraints

n == height.length

1 <= n <= 2 * 10⁴

0 <= height[i] <= 10⁵

🧠 Key Insight (important)

At any index i:

water[i] = min(maxLeft[i], maxRight[i]) - height[i]


If this value is negative → treat it as 0.


Trapping Rain Water — Explanation
The core idea

Water can only be trapped between walls.
At any position, the water level depends on the shorter wall on its left and right.

Why?
Because water will spill over the shorter side.

🧠 Key rule (this is everything)

At index i:

water at i = min(max height to the left,
                 max height to the right)
             − height[i]


If this value is negative → 0 water

🔍 Why this works

Imagine standing on a bar at index i.

Look left → what’s the tallest wall you see?

Look right → what’s the tallest wall you see?

Water can only rise as high as the shorter of those two walls

Anything above the current bar’s height is water.

🧪 Example walkthrough
Input
height = [0,1,0,2,1,0,1,3,2,1,2,1]

dry run for Example 1
We’ll use the intuition formula:

water[i] = min(maxLeft[i], maxRight[i]) - height[i]

Step 1: Draw the heights
Index:  0 1 2 3 4 5 6 7 8 9 10 11
Height: 0 1 0 2 1 0 1 3 2 1  2  1

Step 2: Compute maxLeft for each index

maxLeft[i] = tallest wall to the left of i (including i)

maxLeft: 0 1 1 2 2 2 2 3 3 3 3 3


Explanation:

i=0 → 0

i=1 → max(0,1)=1

i=2 → max(1,0)=1

i=3 → max(1,2)=2

i=4 → max(2,1)=2

i=5 → max(2,0)=2

i=6 → max(2,1)=2

i=7 → max(2,3)=3

i=8 → max(3,2)=3

i=9 → max(3,1)=3

i=10 → max(3,2)=3

i=11 → max(3,1)=3

Step 3: Compute maxRight for each index

maxRight[i] = tallest wall to the right of i (including i)

maxRight: 3 3 3 3 3 3 3 3 2 2 2 1


Explanation:

i=11 → 1

i=10 → max(1,2)=2

i=9 → max(2,1)=2

i=8 → max(2,2)=2

i=7 → max(2,3)=3

i=6 → max(3,1)=3

i=5 → max(3,0)=3

i=4 → max(3,1)=3

i=3 → max(3,2)=3

i=2 → max(3,0)=3

i=1 → max(3,1)=3

i=0 → max(3,0)=3

Step 4: Compute trapped water at each index
water[i] = min(maxLeft[i], maxRight[i]) - height[i]

i	height[i]	maxLeft[i]	maxRight[i]	water[i]
0	0	0	3	0
1	1	1	3	0
2	0	1	3	1
3	2	2	3	0
4	1	2	3	1
5	0	2	3	2
6	1	2	3	1
7	3	3	3	0
8	2	3	2	0
9	1	3	2	1
10	2	3	2	0
11	1	3	1	0
Step 5: Sum all trapped water
water = 1 + 1 + 2 + 1 + 1 = 6


✅ That’s why the total trapped water = 6 units

🚀 Optimized thinking (Two Pointers)

Instead of recomputing max left and right for every index, we:

Start with two pointers: left and right

Track:

leftMax

rightMax

Always move the pointer with the smaller max

Why this works

The smaller side is the limiting factor for water.

If:

leftMax < rightMax


Then:

water at left depends only on leftMax

the right side is tall enough already

Same logic applies symmetrically.

🧠 Visual intuition (important)

Think of filling water from the sides inward.
The side with the lower wall decides the water level.

⏱️ Complexity

Time: O(n) (single pass)

Space: O(1) (no extra arrays)




Container With Most Water — Problem Statement
❓ Problem:

You are given an array of positive integers height where each element represents a vertical line on the x-axis at position i.

Find two lines, which together with the x-axis, form a container, such that the container holds the most water.

Return the maximum area of water the container can store.

📥 Input

Array of integers height

height[i] = height of the line at position i

📤 Output

Maximum area (integer)

🧪 Example 1
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49


Explanation:

Lines at indices 1 (height=8) and 8 (height=7)

Width = 8 - 1 = 7

Height = min(8,7) = 7

Area = 7 * 7 = 49

🧪 Example 2
Input: height = [1,1]
Output: 1

🔒 Constraints

2 <= height.length <= 10^5

0 <= height[i] <= 10^4

🧠 Key Insight (Two Pointers Pattern)

You want max area = width × min(height[left], height[right])

Two pointers approach is optimal:

Start left = 0, right = n-1

Compute area = (right-left) × min(height[left], height[right])

Move the pointer pointing to the smaller height

Repeat until left >= right

Why move the smaller one?

Moving the larger one cannot increase the area because the height is still limited by the smaller line.

⏱ Complexity

Time: O(n) (single pass)

Space: O(1) (no extra arrays)


dry run 

Input:
height = [1,8,6,2,5,4,8,3,7]


We want max area using two pointers.

Step 0: Initialize
left = 0
right = 8  // last index
maxArea = 0

Step 1: Compute initial area
height[left] = 1
height[right] = 7
width = right - left = 8
area = min(1,7) * 8 = 1 * 8 = 8
maxArea = max(0, 8) = 8


Move smaller height → left (1 < 7)

left = 1, right = 8

Step 2:
height[left] = 8
height[right] = 7
width = 8 - 1 = 7
area = min(8,7) * 7 = 7 * 7 = 49
maxArea = max(8, 49) = 49


Move smaller height → right (7 < 8)

left = 1, right = 7

Step 3:
height[left] = 8
height[right] = 3
width = 7 - 1 = 6
area = min(8,3) * 6 = 3 * 6 = 18
maxArea = max(49, 18) = 49


Move smaller height → right (3 < 8)

left = 1, right = 6

Step 4:
height[left] = 8
height[right] = 8
width = 6 - 1 = 5
area = min(8,8) * 5 = 8 * 5 = 40
maxArea = max(49, 40) = 49


Move either (let’s move right)

left = 1, right = 5

Step 5:
height[left] = 8
height[right] = 4
width = 5 - 1 = 4
area = min(8,4) * 4 = 4 * 4 = 16
maxArea = max(49, 16) = 49


Move smaller height → right (4 < 8)

left = 1, right = 4

Step 6:
height[left] = 8
height[right] = 5
width = 4 - 1 = 3
area = min(8,5) * 3 = 5 * 3 = 15
maxArea = max(49, 15) = 49


Move smaller height → right (5 < 8)

left = 1, right = 3

Step 7:
height[left] = 8
height[right] = 2
width = 3 - 1 = 2
area = min(8,2) * 2 = 2 * 2 = 4
maxArea = max(49, 4) = 49


Move smaller height → right (2 < 8)

left = 1, right = 2

Step 8:
height[left] = 8
height[right] = 6
width = 2 - 1 = 1
area = min(8,6) * 1 = 6 * 1 = 6
maxArea = max(49, 6) = 49


Move smaller height → right (6 < 8)

left = 1, right = 1

Step 9:
left >= right → stop

✅ Result
maxArea = 49

