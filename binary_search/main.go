package main

import "fmt"

// ============================================================================
// BINARY SEARCH - O(log n) Divide and Conquer
// ============================================================================
//
// WHAT IS BINARY SEARCH?
// Searching a sorted array by repeatedly dividing the search space in half.
// Think: Finding a word in a dictionary - you don't start at page 1!
//
// CRITICAL REQUIREMENT: Array MUST be sorted!
//
// HOW IT WORKS:
// 1. Look at middle element
// 2. If target equals middle → found it!
// 3. If target < middle → search left half
// 4. If target > middle → search right half
// 5. Repeat until found or search space is empty
//
// VISUAL:
// Array: [1, 3, 5, 7, 9, 11, 13, 15, 17], target = 7
//
// Step 1: left=0, right=8, mid=4
//   [1, 3, 5, 7, 9, 11, 13, 15, 17]
//                 ↑
//              arr[4]=9 > 7 → search left
//
// Step 2: left=0, right=3, mid=1
//   [1, 3, 5, 7] (right half eliminated)
//       ↑
//    arr[1]=3 < 7 → search right
//
// Step 3: left=2, right=3, mid=2
//   [5, 7] (left half eliminated)
//    ↑
//  arr[2]=5 < 7 → search right
//
// Step 4: left=3, right=3, mid=3
//   [7]
//    ↑
//  arr[3]=7 == 7 → FOUND!
//
// COMPLEXITY:
// Time:  O(log n) - halve search space each iteration
// Space: O(1) - only use a few variables
//
// WHY O(log n)?
// - Iteration 1: n elements
// - Iteration 2: n/2 elements
// - Iteration 3: n/4 elements
// - Iteration k: n/(2^k) elements
// - Stop when n/(2^k) = 1 → k = log₂(n)
//
// EXAMPLES:
// - 1000 elements: max ~10 iterations (2^10 = 1024)
// - 1,000,000 elements: max ~20 iterations (2^20 = 1,048,576)
//
// WHEN TO USE:
// ✓ Array is sorted
// ✓ Need faster than O(n) search
// ✓ Finding insertion position
// ✓ Finding first/last occurrence
// ✓ Searching in ranges (not just arrays)
//
// ============================================================================

// binarySearch searches for target in sorted array
// Returns index if found, otherwise returns where it should be inserted
//
// COMPLEXITY:
// Time:  O(log n) - divide search space in half each time
// Space: O(1) - only use left, right, mid variables
//
// EXAMPLE TRACE:
// Input: arr = [1, 2, 3, 4], tar = 10
//
// Initial:
//
//	left = 0, right = 3
//	arr = [1, 2, 3, 4]
//	       ↑        ↑
//	     left    right
//
// Iteration 1:
//
//	left=0, right=3
//	mid = 0 + (3-0)/2 = 0 + 1 = 1
//	arr[mid] = arr[1] = 2
//	Is 2 == 10? No
//	Is 2 < 10? Yes → search right half
//	left = mid + 1 = 2
//
// Iteration 2:
//
//	left=2, right=3
//	mid = 2 + (3-2)/2 = 2 + 0 = 2
//	arr[mid] = arr[2] = 3
//	Is 3 == 10? No
//	Is 3 < 10? Yes → search right half
//	left = mid + 1 = 3
//
// Iteration 3:
//
//	left=3, right=3
//	mid = 3 + (3-3)/2 = 3 + 0 = 3
//	arr[mid] = arr[3] = 4
//	Is 4 == 10? No
//	Is 4 < 10? Yes → search right half
//	left = mid + 1 = 4
//
// Iteration 4:
//
//	left=4, right=3
//	left <= right? NO (4 <= 3 is false)
//	Exit loop
//
// Return left = 4 (insertion position)
//
// This means: 10 should be inserted at index 4 (after all existing elements)
// Result: 4
func binarySearch(arr []int, tar int) int {
	// Initialize search boundaries
	left := 0             // Start of search space
	right := len(arr) - 1 // End of search space

	// Continue while search space is valid
	// When left > right, search space is empty (not found)
	for left <= right {
		// Calculate middle index
		// WHY mid = left + (right-left)/2 NOT (left+right)/2?
		//
		// REASON 1: Prevent integer overflow
		//   If left and right are large (near max int), left+right could overflow
		//   Example: left=2^30, right=2^30 → left+right = 2^31 (overflow!)
		//   But left + (right-left)/2 never overflows
		//
		// REASON 2: Mathematical equivalence
		//   left + (right-left)/2
		//   = left + (right/2 - left/2)
		//   = left/2 + right/2
		//   = (left+right)/2  (without overflow risk)
		mid := left + (right-left)/2

		// Check if we found the target
		if arr[mid] == tar {
			// Found it! Return the index
			return mid

		} else if arr[mid] < tar {
			// Middle value is too small
			// Target must be in right half
			// Eliminate left half by moving left boundary
			left = mid + 1

		} else {
			// arr[mid] > tar
			// Middle value is too large
			// Target must be in left half
			// Eliminate right half by moving right boundary
			right = mid - 1
		}
	}

	// Target not found
	// left is now the insertion position
	//
	// SEARCH INSERT POSITION PATTERN:
	// When binary search fails, left points to where element should be inserted
	// This maintains sorted order
	//
	// Example: arr=[1,3,5,7], tar=4
	//   Final state: left=2 (between 3 and 5)
	//   Insert at index 2: [1,3,4,5,7] ✓
	return left
}

// ============================================================================
// BINARY SEARCH VARIATIONS:
// ============================================================================
//
// 1. SEARCH IN RANGES (not just arrays):
//    Problem: Find square root of n
//    Range: 0 to n
//    Check: if mid*mid == n → found
//           if mid*mid < n → search right
//           if mid*mid > n → search left
//
//    func sqrt(n int) int {
//        left, right := 0, n
//        for left <= right {
//            mid := left + (right-left)/2
//            square := mid * mid
//            if square == n {
//                return mid
//            } else if square < n {
//                left = mid + 1
//            } else {
//                right = mid - 1
//            }
//        }
//        return right  // Floor of square root
//    }
//
// 2. FIRST OCCURRENCE:
//    When arr[mid] == tar, don't return immediately
//    Continue searching left: right = mid - 1
//    Track last found index
//
// 3. LAST OCCURRENCE:
//    When arr[mid] == tar, don't return immediately
//    Continue searching right: left = mid + 1
//    Track last found index
//
// 4. ROTATED ARRAY:
//    Array sorted then rotated: [4,5,6,7,0,1,2]
//    One half is always sorted
//    Check which half is sorted, then decide where to search
//
// ============================================================================
// KEY CONCEPTS:
// ============================================================================
//
// WHY ARRAY MUST BE SORTED:
// Binary search assumes: if arr[mid] < target, all elements left of mid are < target
// This only works if array is sorted!
//
// Unsorted example: [3, 1, 4, 2]  target = 1
//   mid = 1, arr[1] = 1  ← We'd find it by luck
//   But: [3, 2, 4, 1]  target = 1
//   mid = 1, arr[1] = 2 > 1 → search left
//   left half is [3] → doesn't contain 1!  ✗ FAILS
//
// WHEN LEFT > RIGHT:
// This means the search space is empty (exhausted all possibilities)
// left has "crossed over" right, indicating target not found
//
// INVARIANT:
// If target exists, it's always between left and right
// We maintain this by only eliminating halves that can't contain target
//
// ============================================================================

func main() {
	// Test binary search on sorted array
	fmt.Println("=== Binary Search ===")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Array:", arr)
	fmt.Println("\nSearching for values:")
	fmt.Println("  5 found at index:", binarySearch(arr, 5))   // 4
	fmt.Println("  1 found at index:", binarySearch(arr, 1))   // 0
	fmt.Println("  10 found at index:", binarySearch(arr, 10)) // 9

	fmt.Println("\nSearching for missing values (returns insertion position):")
	fmt.Println("  0 should insert at:", binarySearch(arr, 0))   // 0 (before 1)
	fmt.Println("  11 should insert at:", binarySearch(arr, 11)) // 10 (after 10)
	fmt.Println("  5.5 (as 5) insert at:", binarySearch(arr, 6)) // Actually finds 6 at index 5

	// Demonstrate on larger array (see logarithmic scaling)
	fmt.Println("\n=== Logarithmic Scaling ===")
	large := make([]int, 1000)
	for i := range large {
		large[i] = i * 2 // 0, 2, 4, 6, ..., 1998
	}
	fmt.Println("Array size: 1000")
	fmt.Println("Find 998 at index:", binarySearch(large, 998)) // ~10 iterations max!

	// Example: Search insert position
	fmt.Println("\n=== Search Insert Position Pattern ===")
	sorted := []int{1, 3, 5, 7}
	fmt.Println("Array:", sorted)
	fmt.Println("Insert position for 4:", binarySearch(sorted, 4)) // 2 (between 3 and 5)
	fmt.Println("Insert position for 0:", binarySearch(sorted, 0)) // 0 (at beginning)
	fmt.Println("Insert position for 8:", binarySearch(sorted, 8)) // 4 (at end)
}
