package main

import "fmt"

// ============================================================================
// ARRAYS - Contiguous Memory & Index Math
// ============================================================================
//
// WHAT IS AN ARRAY?
// Arrays store elements in CONTIGUOUS (adjacent) memory locations.
// Think of numbered parking spots in a row.
//
// MEMORY LAYOUT:
// ┌───────────────────────────────────────────────┐
// │  Index:  │  0  │  1  │  2  │  3  │  4  │  5  │
// ├───────────────────────────────────────────────┤
// │  Value:  │  1  │  2  │  3  │  4  │  5  │  6  │
// ├───────────────────────────────────────────────┤
// │ Address: │1000 │1004 │1008 │1012 │1016 │1020 │
// └───────────────────────────────────────────────┘
//
// INDEX MATH:
// To access arr[3]:
// memory_address = base_address + (index * element_size)
// memory_address = 1000 + (3 * 4 bytes) = 1012
//
// This is why array access is O(1) - simple math!
//
// ============================================================================
// ARRAY OPERATIONS COMPLEXITY:
// ============================================================================
// Access by index:     O(1)  - direct memory calculation
// Search (unsorted):   O(n)  - must check each element
// Insert at end:       O(1)* - just append (*amortized)
// Insert at beginning: O(n)  - must shift ALL elements right
// Insert at middle:    O(n)  - must shift remaining elements
// Delete at end:       O(1)  - just reduce length
// Delete at beginning: O(n)  - must shift ALL elements left
// Delete at middle:    O(n)  - must shift remaining elements
//
// WHY SHIFTING IS EXPENSIVE:
// To insert 'X' at index 1 in [A, B, C, D]:
//   [A, B, C, D]
//   [A, _, B, C, D]  ← shift B, C, D right (3 operations)
//   [A, X, B, C, D]  ← insert X
// Total: O(n) operations
//
// ============================================================================

// revArr reverses an array in-place using two pointers
//
// COMPLEXITY:
// Time:  O(n) - visit each element once (n/2 swaps, still O(n))
// Space: O(1) - only use 3 variables regardless of array size
//
// TWO POINTER PATTERN:
// Start at both ends, swap, move inward until pointers meet
//
// EXAMPLE TRACE:
// Input: arr = [1, 2, 3, 4, 5]
//
// Initial state:
// arr = [1, 2, 3, 4, 5]
//
//	 ↑           ↑
//	left       right
//
// Iteration 1:
//
//	left=0, right=4
//	temp = arr[0] = 1
//	arr[0] = arr[4] = 5
//	arr[4] = temp = 1
//	arr = [5, 2, 3, 4, 1]
//	left++ → 1, right-- → 3
//
// Iteration 2:
//
//	left=1, right=3
//	temp = arr[1] = 2
//	arr[1] = arr[3] = 4
//	arr[3] = temp = 2
//	arr = [5, 4, 3, 2, 1]
//	left++ → 2, right-- → 2
//
// Iteration 3:
//
//	left=2, right=2
//	left < right? NO (2 < 2 is false)
//	Loop exits
//
// Result: [5, 4, 3, 2, 1]
//
// WHY IS THIS EFFICIENT?
// - Only one pass through array
// - No extra array needed
// - In-place modification
func revArr(arr []int) []int {
	// Initialize two pointers
	left := 0             // Start at beginning (index 0)
	right := len(arr) - 1 // Start at end (last valid index)

	// Continue until pointers meet or cross
	// When left >= right, we've processed all elements
	for left < right {
		// Swap elements using temporary variable
		// This is the classic swap pattern

		temp := arr[left]      // Store left value (would be lost otherwise)
		arr[left] = arr[right] // Overwrite left with right value
		arr[right] = temp      // Put original left value into right position

		// Move pointers toward center
		left++  // Move left pointer one step right
		right-- // Move right pointer one step left
	}

	// Return the modified array
	// Note: array was modified in-place, but we return it for convenience
	return arr
}

// ============================================================================

// mulArr multiplies each element by its index
//
// COMPLEXITY:
// Time:  O(n) - single loop through array
// Space: O(n) - create new array of same size
//
// EXAMPLE TRACE:
// Input: arr = [1, 2, 3, 4, 5]
//
// Initial:
// newArr = []
//
// Iteration 1: i=0, n=1
//
//	mul = 1 * 0 = 0
//	newArr = [0]
//
// Iteration 2: i=1, n=2
//
//	mul = 2 * 1 = 2
//	newArr = [0, 2]
//
// Iteration 3: i=2, n=3
//
//	mul = 3 * 2 = 6
//	newArr = [0, 2, 6]
//
// Iteration 4: i=3, n=4
//
//	mul = 4 * 3 = 12
//	newArr = [0, 2, 6, 12]
//
// Iteration 5: i=4, n=5
//
//	mul = 5 * 4 = 20
//	newArr = [0, 2, 6, 12, 20]
//
// Result: [0, 2, 6, 12, 20]
//
// WHY NEW ARRAY?
// We're building results, not modifying in place.
// Original array remains unchanged.
func mulArr(arr []int) []int {
	// Create empty slice to collect results
	// Go slices automatically grow as needed
	newArr := []int{}

	// range gives us both index (i) and value (n)
	// This is idiomatic Go - better than C-style for loop
	for i, n := range arr {
		// Multiply value by its position
		mul := n * i

		// append adds element to end of slice
		// Go handles memory management automatically
		// If capacity exceeded, Go allocates larger array and copies
		// This is "amortized O(1)" - occasional O(n) copy, but rare
		newArr = append(newArr, mul)
	}

	return newArr
}

// ============================================================================
// KEY ARRAY CONCEPTS:
// ============================================================================
//
// 1. ZERO-INDEXED:
//    First element is arr[0], not arr[1]
//    Last element is arr[len(arr)-1]
//
// 2. BOUNDS CHECKING:
//    Go automatically checks bounds
//    arr[10] on 5-element array → panic: index out of range
//
// 3. SLICE VS ARRAY:
//    Array: Fixed size, value type [5]int
//    Slice: Dynamic size, reference type []int
//    Slices are more common in Go
//
// 4. CONTIGUOUS MEMORY = FAST ACCESS:
//    CPU cache-friendly
//    Sequential access is very fast
//    Random access is also O(1)
//
// 5. COST OF RESIZING:
//    Inserting in middle requires shifting
//    [1,2,3] insert 99 at index 1 → [1,99,2,3]
//    Must shift 2 and 3 right: O(n) operation
//
// ============================================================================

func main() {
	// Test mulArr
	fmt.Println("=== Multiply Array ===")
	result := mulArr([]int{1, 2, 3, 4, 5})
	fmt.Println(result) // [0, 2, 6, 12, 20]

	// Test revArr
	fmt.Println("\n=== Reverse Array ===")
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Before:", arr)
	reversed := revArr(arr)
	fmt.Println("After: ", reversed)
	fmt.Println("Original was modified:", arr) // Same as reversed

	// Demonstrate two-pointer pattern
	fmt.Println("\n=== Two Pointer Pattern Demo ===")
	palindrome := []int{1, 2, 3, 2, 1}
	fmt.Println("Palindrome array:", palindrome)
	revArr(palindrome)
	fmt.Println("Reversed:       ", palindrome)
	// Notice: palindrome reversed equals itself!
}
