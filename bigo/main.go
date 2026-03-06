package main

import "fmt"

// ============================================================================
// BIG O NOTATION - Understanding Time Complexity
// ============================================================================
//
// Big O describes how an algorithm's runtime grows as input size increases
//
// Common complexities:
// O(1)      - Constant:     Same time regardless of input size
// O(log n)  - Logarithmic:  Divides problem in half each time
// O(n)      - Linear:       Time grows proportionally with input
// O(n log n)- Linearithmic: Efficient sorting algorithms
// O(n²)     - Quadratic:    Nested loops over input
// O(2^n)    - Exponential:  Doubles with each additional element
//
// Focus on WORST CASE and ignore constants:
//   5n + 3 → O(n)
//   n² + n → O(n²)  (higher order dominates)
//
// ============================================================================

// findTar searches for a target value in an array
//
// COMPLEXITY ANALYSIS:
// Time:  O(n) - worst case, we check every element
// Space: O(1) - only use a few variables regardless of input size
//
// WHY O(n)?
// - The loop runs at most n times (where n = len(arr))
// - If array has 100 elements, max 100 iterations
// - If array has 1000 elements, max 1000 iterations
// - Time scales linearly with input size
//
// EXAMPLE TRACE:
// Input: arr = [1, 2, 3, 4], tar = 4
//
// Iteration 1:
//
//	n = 1
//	Is 1 == 4? No
//	Continue...
//
// Iteration 2:
//
//	n = 2
//	Is 2 == 4? No
//	Continue...
//
// Iteration 3:
//
//	n = 3
//	Is 3 == 4? No
//	Continue...
//
// Iteration 4:
//
//	n = 4
//	Is 4 == 4? YES!
//	Return true immediately
//
// Result: true (found at index 3)
func findTar(arr []int, tar int) bool {
	// Loop through each element in the array
	// This is a linear search - checking elements one by one
	for _, n := range arr {
		// Compare current element with target
		// This is an O(1) operation (single comparison)
		if n == tar {
			// Found it! Return immediately (best case)
			// Best case: O(1) if found at first position
			return true
		}
	}

	// Checked all elements, target not found
	// This line only executes in worst case (target not in array)
	// Worst case: O(n) - had to check all n elements
	return false
}

// ============================================================================
// HOW TO READ BIG O FROM CODE:
// ============================================================================
//
// 1. COUNT THE LOOPS:
//    - Single loop over n elements → O(n)
//    - Nested loops → O(n²)
//    - Three nested loops → O(n³)
//
// 2. MAP/HASH OPERATIONS:
//    - map[key] lookup → O(1) average case
//    - map[key] = value → O(1) average case
//
// 3. ARRAY ACCESS:
//    - arr[i] → O(1) (direct memory access)
//
// 4. EXAMPLES:
//
//    for i := 0; i < n; i++ {     ← O(n)
//        // constant work
//    }
//
//    for i := 0; i < n; i++ {     ← O(n²)
//        for j := 0; j < n; j++ { ← nested loop
//            // constant work
//        }
//    }
//
//    for i := 0; i < n; i++ {     ← O(n) not O(n²)!
//        // constant work
//    }
//    for j := 0; j < n; j++ {     ← separate loop
//        // constant work
//    }
//    // Two sequential O(n) loops = O(n + n) = O(2n) = O(n)
//
// ============================================================================

func main() {
	// Test case 1: Target exists
	result1 := findTar([]int{1, 2, 3, 4}, 4)
	fmt.Println(result1) // true

	// Test case 2: Target doesn't exist
	result2 := findTar([]int{1, 2, 3, 4}, 5)
	fmt.Println(result2) // false

	// Test case 3: Target at beginning (best case - O(1))
	result3 := findTar([]int{10, 20, 30}, 10)
	fmt.Println(result3) // true
}
