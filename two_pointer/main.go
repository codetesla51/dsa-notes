package main

import (
	"fmt"
	"regexp"
	"strings"
)

// ============================================================================
// TWO POINTERS - O(n) Pattern for Arrays/Strings
// ============================================================================
//
// WHAT IS TWO POINTER TECHNIQUE?
// Using two indices that traverse data in tandem
// Eliminates need for nested loops: O(n²) → O(n)
//
// THREE MAIN PATTERNS:
//
// 1. OPPOSITE ENDS (convergence):
//    - Start: left=0, right=n-1
//    - Move toward each other
//    - Use for: palindromes, two sum (sorted), reversal
//
//    ┌───┬───┬───┬───┬───┐
//    │ 1 │ 2 │ 3 │ 4 │ 5 │
//    └───┴───┴───┴───┴───┘
//      ↑               ↑
//     left           right
//
// 2. FAST/SLOW (different speeds):
//    - Start: both at 0
//    - Slow moves 1 step, fast moves 2 steps
//    - Use for: cycle detection, finding middle
//
//    ┌───┬───┬───┬───┬───┐
//    │ 1 │ 2 │ 3 │ 4 │ 5 │
//    └───┴───┴───┴───┴───┘
//      ↑   ↑
//    slow fast
//
// 3. SAME DIRECTION (sliding/expanding):
//    - Start: both at 0 (or left at 0, right at some point)
//    - Move in same direction at different rates
//    - Use for: remove duplicates, partitioning
//
//    ┌───┬───┬───┬───┬───┐
//    │ 1 │ 1 │ 2 │ 2 │ 3 │
//    └───┴───┴───┴───┴───┘
//      ↑   ↑
//      i   j
//
// WHY TWO POINTERS BEATS NESTED LOOPS:
//
// Brute force (O(n²)):
//   for i := 0; i < n; i++ {
//       for j := i+1; j < n; j++ {
//           // Check pairs: n*(n-1)/2 comparisons
//       }
//   }
//
// Two pointers (O(n)):
//   left, right := 0, n-1
//   for left < right {
//       // Each element visited once: n comparisons
//   }
//
// For n=1000: 500,000 operations → 1000 operations (500x faster!)
//
// ============================================================================

// twoSum finds if any two numbers sum to target (sorted array)
//
// COMPLEXITY:
// Time:  O(n) - single pass with two pointers
// Space: O(1) - only two pointer variables
//
// PATTERN: Opposite ends (convergence)
//
// WHY THIS WORKS (SORTED ARRAY):
// - If sum too small → left++ (increase sum)
// - If sum too large → right-- (decrease sum)
// - We systematically explore all possibilities in one pass
//
// CONTRAST WITH UNSORTED:
// - Unsorted needs O(n) map or O(n²) brute force
// - Sorting first: O(n log n) + O(n) = still O(n log n)
// - But if already sorted → two pointers is O(n)!
//
// EXAMPLE TRACE:
// Input: arr = [1, 2, 3, 4, 5, 6], tar = 9
//
// Initial:
//
//	left = 0, right = 5
//	arr = [1, 2, 3, 4, 5, 6]
//	       ↑              ↑
//	     left           right
//
// Iteration 1:
//
//	sum = arr[0] + arr[5] = 1 + 6 = 7
//	Is 7 == 9? No
//	Is 7 > 9? No
//	7 < 9 → sum too small → left++
//	left = 1
//
// Iteration 2:
//
//	sum = arr[1] + arr[5] = 2 + 6 = 8
//	Is 8 == 9? No
//	Is 8 > 9? No
//	8 < 9 → sum too small → left++
//	left = 2
//
// Iteration 3:
//
//	sum = arr[2] + arr[5] = 3 + 6 = 9
//	Is 9 == 9? YES!
//	Return true
//
// Result: true (3 + 6 = 9)
//
// WHEN TO MOVE LEFT vs RIGHT:
// - sum < target: need larger sum → move left right (increase smaller number)
// - sum > target: need smaller sum → move right left (decrease larger number)
// - This greedy approach works because array is SORTED
func twoSum(arr []int, tar int) bool {
	// Initialize pointers at opposite ends
	left := 0             // Start at smallest value
	right := len(arr) - 1 // Start at largest value

	// Continue until pointers meet
	// When left == right, we can't make a pair
	for left <= right {
		// Calculate sum of current pair
		sum := arr[left] + arr[right]

		if sum == tar {
			// Found a pair that sums to target!
			return true

		} else if sum > tar {
			// Sum is too large
			// Decrease it by moving right pointer left
			// This uses a smaller number on the right side
			right--

		} else {
			// sum < tar
			// Sum is too small
			// Increase it by moving left pointer right
			// This uses a larger number on the left side
			left++
		}
	}

	// Exhausted all pairs, no match found
	return false
}

// ============================================================================

// palindromeTwoSum checks if string is a palindrome
// Ignores non-alphanumeric characters and case
//
// COMPLEXITY:
// Time:  O(n) - single pass with two pointers (after cleanup)
// Space: O(n) - cleaned string storage
//
// PATTERN: Opposite ends with character comparison
//
// EXAMPLE TRACE:
// Input: s = "A man, a plan, a canal: Panama"
//
// Step 1: Clean string
//
//	Remove non-alphanumeric: "AmanaplanacanalPanama"
//	Lowercase: "amanaplanacanalpanama"
//	s = "amanaplanacanalpanama" (length 21)
//
// Step 2: Two pointer palindrome check
//
//	left=0, right=20
//
// Iteration 1:
//
//	s[0]='a', s[20]='a'
//	Match? Yes → continue
//	left=1, right=19
//
// Iteration 2:
//
//	s[1]='m', s[19]='m'
//	Match? Yes → continue
//	left=2, right=18
//
// ... (all match)
//
// Iteration 11:
//
//	left=10, right=10 (same position)
//	left < right? No (10 < 10 is false)
//	Exit loop
//
// Result: true (is palindrome)
//
// WHY REGEX FOR CLEANING:
// regexp.MustCompile("[^a-zA-Z0-9]+") means:
//
//	[^...] = NOT these characters
//	a-zA-Z0-9 = letters and numbers
//	+ = one or more
//
// Result: Matches all non-alphanumeric characters for removal
func palindromeTwoSum(s string) bool {
	// STEP 1: Clean the string

	// Compile regex to match non-alphanumeric characters
	// [^a-zA-Z0-9] means "anything that's NOT a letter or number"
	re := regexp.MustCompile("[^a-zA-Z0-9]+")

	// Replace all non-alphanumeric with empty string (remove them)
	// "A man, a plan!" → "Amanaplana"
	clean := re.ReplaceAllString(s, "")

	// Convert to lowercase for case-insensitive comparison
	// Trim any remaining whitespace
	// "Amanaplana" → "amanaplana"
	s = strings.ToLower(strings.TrimSpace(clean))

	// STEP 2: Two pointer palindrome check

	// Initialize pointers at opposite ends
	left := 0
	right := len(s) - 1

	// Check characters from both ends moving inward
	for left < right {
		// Compare characters at left and right positions
		if s[left] != s[right] {
			// Characters don't match → not a palindrome
			return false
		}

		// Characters match, move both pointers inward
		left++
		right--
	}

	// All characters matched → is a palindrome
	return true
}

// ============================================================================
// TWO POINTER PATTERNS SUMMARY:
// ============================================================================
//
// WHEN TO USE EACH PATTERN:
//
// 1. OPPOSITE ENDS:
//    ✓ Two sum (sorted array)
//    ✓ Palindrome checking
//    ✓ Reverse array/string
//    ✓ Container with most water
//    ✓ Trapping rain water
//
//    Template:
//      left, right := 0, n-1
//      for left < right {
//          // process arr[left] and arr[right]
//          // move pointers based on condition
//      }
//
// 2. FAST/SLOW:
//    ✓ Find middle of linked list
//    ✓ Cycle detection
//    ✓ Find kth element from end
//
//    Template:
//      slow, fast := head, head
//      for fast != nil && fast.Next != nil {
//          slow = slow.Next
//          fast = fast.Next.Next
//      }
//
// 3. SAME DIRECTION:
//    ✓ Remove duplicates in-place
//    ✓ Move zeros to end
//    ✓ Partition array
//
//    Template:
//      i := 0
//      for j := 0; j < n; j++ {
//          if condition {
//              arr[i], arr[j] = arr[j], arr[i]
//              i++
//          }
//      }
//
// ============================================================================
// KEY INSIGHTS:
// ============================================================================
//
// WHY IT WORKS:
// Two pointers maintain an "invariant" - a condition that's always true
//
// Example (two sum sorted):
//   Invariant: If a solution exists, it's between arr[left] and arr[right]
//   - If sum too small, left++ maintains invariant (might find larger sum)
//   - If sum too large, right-- maintains invariant (might find smaller sum)
//   - We never skip a potential solution
//
// WHEN NOT TO USE:
// ✗ Array is unsorted (unless you sort first)
// ✗ Need to check all pairs without elimination logic
// ✗ Problem doesn't have monotonic property
//
// ============================================================================

func main() {
	// Test two sum on sorted array
	fmt.Println("=== Two Sum (Sorted Array) ===")
	sorted := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Array:", sorted)
	fmt.Println("Target 9:", twoSum(sorted, 9))   // true (3+6)
	fmt.Println("Target 11:", twoSum(sorted, 11)) // true (5+6)
	fmt.Println("Target 15:", twoSum(sorted, 15)) // false

	// Test palindrome
	fmt.Println("\n=== Palindrome Check ===")
	fmt.Println("'A man a plan a canal Panama':", palindromeTwoSum("A man a plan a canal Panama")) // true
	fmt.Println("'race a car':", palindromeTwoSum("race a car"))                                   // false
	fmt.Println("'Was it a car or a cat I saw':", palindromeTwoSum("Was it a car or a cat I saw")) // true
	fmt.Println("'hello':", palindromeTwoSum("hello"))                                             // false

	// Demonstrate why two pointers is faster
	fmt.Println("\n=== Performance Comparison ===")
	fmt.Println("Array size: 1000")
	fmt.Println("Nested loops: ~500,000 comparisons (O(n²))")
	fmt.Println("Two pointers: ~1000 comparisons (O(n))")
	fmt.Println("Speed-up: 500x faster!")
}
