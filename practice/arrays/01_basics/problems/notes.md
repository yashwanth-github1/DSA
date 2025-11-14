Array Basic Operations

Example Array:

arr = [10, 20, 30, 40, 50]
length = 5

1. Access / Read

Operation: Access an element by its index.

Example: arr[2] → 30

Time Complexity: O(1) → Direct access via index
Space Complexity: O(1) → No extra space

Explanation: Arrays store elements in contiguous memory. Index gives direct address.

2. Update / Modify

Operation: Change value at a specific index.

Example: arr[1] = 25 → [10, 25, 30, 40, 50]

Time Complexity: O(1) → Direct index access
Space Complexity: O(1) → In-place modification

3. Search / Linear Search

Operation: Find the index of a target element.

Example: Find 30

Algorithm:

for i in 0 to n-1:
    if arr[i] == target:
        return i
return -1


Time Complexity:

Best: O(1) → target at first index

Worst: O(n) → target at last or not present

Average: O(n)

Space Complexity: O(1) → no extra memory

4. Insertion

Operation: Insert element at a given index.

Example: Insert 15 at index 1 → [10, 15, 25, 30, 40, 50]

Algorithm:

Shift all elements from the index to the right by 1

Insert element

Time Complexity:

Best: O(1) → inserting at end

Worst: O(n) → inserting at start requires shifting all elements

Space Complexity: O(1) → in-place if array has extra capacity
O(n) → if a new array is created to increase size

5. Deletion

Operation: Remove element at a given index.

Example: Delete element at index 2 → [10, 25, 40, 50]

Algorithm:

Shift all elements after the index to the left by 1

Time Complexity:

Best: O(1) → deleting last element

Worst: O(n) → deleting first element requires shifting all

Space Complexity: O(1) → in-place

6. Traversal

Operation: Visit all elements.

Example: Print all elements → 10 25 30 40 50

Time Complexity: O(n) → need to visit all elements
Space Complexity: O(1) → if just reading

7. Finding Maximum / Minimum

Operation: Traverse array to find max/min.

Algorithm:

max_val = arr[0]
for i in 1 to n-1:
    if arr[i] > max_val:
        max_val = arr[i]
return max_val


Time Complexity: O(n) → visit all elements
Space Complexity: O(1)

8. Sum / Average / Aggregate Operations

Operation: Compute sum or average of elements

Algorithm:

sum = 0
for i in 0 to n-1:
    sum += arr[i]
average = sum / n


Time Complexity: O(n)
Space Complexity: O(1)

9. Reverse Array

Operation: Reverse elements in-place

Algorithm (Two-pointer):

i = 0, j = n-1
while i < j:
    swap(arr[i], arr[j])
    i++, j--


Time Complexity: O(n) → swap n/2 elements
Space Complexity: O(1) → in-place

10. Copy / Clone Array

Operation: Make a copy of an array

Example: newArr = arr[:]

Time Complexity: O(n) → all elements copied
Space Complexity: O(n) → new array created

11. Rotate Array

Operation: Rotate left or right

Example: Left rotate [10, 20, 30, 40, 50] by 2 → [30, 40, 50, 10, 20]

Time Complexity: O(n) → need to move elements
Space Complexity:

O(n) → using extra array

O(1) → using reversal method

12. Count Frequency / Occurrences

Operation: Count how many times a value occurs

Algorithm:

count = 0
for i in 0 to n-1:
    if arr[i] == x:
        count++


Time Complexity: O(n)
Space Complexity: O(1) → simple count

✅ Summary Table of Array Operations

Operation	Time Complexity	Space Complexity	Notes
Access / Read	O(1)	O(1)	Direct index
Update / Modify	O(1)	O(1)	In-place
Linear Search	O(n)	O(1)	Target unknown
Insertion	O(n)	O(1)/O(n)	Shift elements
Deletion	O(n)	O(1)	Shift elements
Traversal	O(n)	O(1)	Visit all
Find Max / Min	O(n)	O(1)	Single pass
Sum / Average	O(n)	O(1)	Single pass
Reverse	O(n)	O(1)	Two-pointer
Copy / Clone	O(n)	O(n)	New array
Rotate	O(n)	O(n)/O(1)	Depends on method
Count Frequency	O(n)	O(1)	Simple count