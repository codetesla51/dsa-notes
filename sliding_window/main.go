package main

import "fmt"

// ============================================================================
// SLIDING WINDOW - O(n) Subarray/Substring Problems
// ============================================================================
//
// WHAT IS SLIDING WINDOW?
// A technique to process subarrays/substrings efficiently by:
// 1. Maintaining a "window" of elements
// 2. Sliding the window by adding new element and removing old element
// 3. Avoiding recalculation from scratch each time
//
// VISUAL (Fixed Window):
// Array: [1, 2, 3, 4, 5], k=3 (window size)
//
// Window 1: [1, 2, 3] sum=6
//            -------
// Window 2:    [2, 3, 4] sum=9 (added 4, removed 1)
//               -------
// Window 3:       [3, 4, 5] sum=12 (added 5, removed 2)
//                  -------
//
// TWO TYPES:
//
// 1. FIXED WINDOW:
//    - Window size is constant
//    - Slide by: add arr[right], remove arr[right-k]
//    - Use for: max sum of k elements, average of k elements
//
// 2. VARIABLE WINDOW:
//    - Window size changes
//    - Expand: right++ (add element)
//    - Shrink: left++ (remove element)
//    - Use for: longest substring with condition, minimum window
//
// WHY SLIDING WINDOW?
//
// Brute force (O(n*k) or O(n²)):
//   for i := 0; i <= n-k; i++ {
//       sum := 0
//       for j := i; j < i+k; j++ {  // Recalculate sum each time!
//           sum += arr[j]
//       }
//   }
//
// Sliding window (O(n)):
//   sum := sum of first k elements
//   for i := k; i < n; i++ {
//       sum += arr[i]      // Add new element
//       sum -= arr[i-k]    // Remove old element
//   }
//
// For n=1000, k=100: 100,000 operations → 1000 operations (100x faster!)
//
// ============================================================================

// slidingWindow finds maximum sum of k consecutive elements
//
// COMPLEXITY:
// Time:  O(n) - build first window O(k) + slide through array O(n-k) = O(n)
// Space: O(1) - only use sum and max variables
//
// PATTERN: Fixed window size
//
// KEY INSIGHTS:
// - arr[i] is the NEW element entering the window
// - arr[i-k] is the OLD element leaving the window
// - Window size is always k
//
// WHY arr[i-k] IS THE LEAVING ELEMENT:
// When we're at index i with window size k:
//
//	Window contains: arr[i-k+1], arr[i-k+2], ..., arr[i-1], arr[i]
//
//	Example: i=5, k=3
//	  Window = [arr[3], arr[4], arr[5]]
//	  When we move to i=6:
//	    Add arr[6]
//	    Remove arr[6-3] = arr[3] ← the leftmost element
//
// EXAMPLE TRACE:
// Input: arr = [1, 2, 3, 4, 5], k = 3
//
// Step 1: Build first window
//
//	Window: [1, 2, 3]
//	sum = 1 + 2 + 3 = 6
//	max = 6
//
// Step 2: Slide window (i=3 to i=4)
//
// i=3:
//
//	Before: [1, 2, 3] sum=6
//	        -------
//	Add arr[3]=4: sum = 6 + 4 = 10
//	Remove arr[0]=1: sum = 10 - 1 = 9
//	After: [2, 3, 4] sum=9
//	       -------
//	max = max(6, 9) = 9
//
// i=4:
//
//	Before: [2, 3, 4] sum=9
//	        -------
//	Add arr[4]=5: sum = 9 + 5 = 14
//	Remove arr[1]=2: sum = 14 - 2 = 12
//	After: [3, 4, 5] sum=12
//	       -------
//	max = max(9, 12) = 12
//
// Result: 12
func slidingWindow(arr []int, k int) int {
	// STEP 1: Build the first window
	// Calculate sum of first k elements
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	// Initialize max with first window's sum
	max := sum

	// STEP 2: Slide the window
	// Start from index k (the first element after initial window)
	for i := k; i < len(arr); i++ {
		// ADD: New element entering the window on the right
		sum += arr[i]

		// REMOVE: Old element leaving the window on the left
		// arr[i-k] is the element that's now outside the window
		// Example: if i=5 and k=3, we remove arr[2] (the element 3 positions back)
		sum -= arr[i-k]

		// Update max if current window has larger sum
		if sum > max {
			max = sum
		}
	}

	return max
}

// ============================================================================

// longestSubString finds longest substring without repeating characters
//
// COMPLEXITY:
// Time:  O(n) - each character visited at most twice (once by right, once by left)
// Space: O(min(n, m)) - map stores unique characters, m = character set size
//
// PATTERN: Variable window with hashmap
//
// KEY INSIGHTS:
// - seen map stores LAST INDEX of each character (not just existence!)
// - idx >= left checks if character is in CURRENT window
// - When duplicate found: left = idx + 1 (skip past previous occurrence)
// - Window size = right - left + 1
//
// WHY STORE LAST INDEX NOT JUST EXISTENCE?
// We need to know WHERE we saw the character to update left pointer correctly
//
// WHY idx >= left CHECK?
// Character might be in map but OUTSIDE current window!
// Example: s = "abba"
//
//	When we see second 'a':
//	  seen = {'a':0, 'b':1, 'b':2}
//	  left = 2 (moved past first 'b')
//	  seen['a'] = 0, but 0 < 2 (outside current window!)
//	  We shouldn't move left backwards
//
// WINDOW SIZE FORMULA:
// right - left + 1
// Example: left=2, right=5 → indices 2,3,4,5 → 4 elements
//
// EXAMPLE TRACE:
// Input: s = "abcabcbb"
//
// Initial:
//
//	seen = {}
//	left = 0
//	max = 0
//
// right=0, ch='a':
//
//	Is 'a' in seen? No
//	seen = {'a': 0}
//	Window: "a" (length 1)
//	max = 1
//
// right=1, ch='b':
//
//	Is 'b' in seen? No
//	seen = {'a':0, 'b':1}
//	Window: "ab" (length 2)
//	max = 2
//
// right=2, ch='c':
//
//	Is 'c' in seen? No
//	seen = {'a':0, 'b':1, 'c':2}
//	Window: "abc" (length 3)
//	max = 3
//
// right=3, ch='a':
//
//	Is 'a' in seen AND idx >= left? Yes (idx=0 >= 0)
//	DUPLICATE FOUND!
//	left = idx + 1 = 0 + 1 = 1
//	seen['a'] = 3 (update last index)
//	Window: "bca" (length 3)
//	max = 3
//
// right=4, ch='b':
//
//	Is 'b' in seen AND idx >= left? Yes (idx=1 >= 1)
//	DUPLICATE FOUND!
//	left = idx + 1 = 1 + 1 = 2
//	seen['b'] = 4
//	Window: "cab" (length 3)
//	max = 3
//
// right=5, ch='c':
//
//	Is 'c' in seen AND idx >= left? Yes (idx=2 >= 2)
//	DUPLICATE FOUND!
//	left = idx + 1 = 2 + 1 = 3
//	seen['c'] = 5
//	Window: "abc" (length 3)
//	max = 3
//
// right=6, ch='b':
//
//	Is 'b' in seen AND idx >= left? Yes (idx=4 >= 3)
//	DUPLICATE FOUND!
//	left = idx + 1 = 4 + 1 = 5
//	seen['b'] = 6
//	Window: "cb" (length 2)
//	max = 3
//
// right=7, ch='b':
//
//	Is 'b' in seen AND idx >= left? Yes (idx=6 >= 5)
//	DUPLICATE FOUND!
//	left = idx + 1 = 6 + 1 = 7
//	seen['b'] = 7
//	Window: "b" (length 1)
//	max = 3
//
// Result: 3 (longest substring is "abc" or "bca" or "cab")
func longestSubString(s string) int {
	// Map to store last seen index of each character
	// Key: character, Value: last index where we saw it
	seen := make(map[byte]int)

	// Left boundary of current window
	left := 0

	// Maximum length found so far
	max := 0

	// Right pointer expands the window
	for right := 0; right < len(s); right++ {
		// Current character
		ch := s[right]

		// Check if character is in map AND inside current window
		if idx, ok := seen[ch]; ok && idx >= left {
			// DUPLICATE FOUND in current window!

			// Move left pointer past the previous occurrence
			// This shrinks the window to exclude the duplicate
			// idx is where we last saw this character
			// idx + 1 is the position right after it
			left = idx + 1
		}

		// Update last seen index of current character
		// Do this even if we found a duplicate (update to new position)
		seen[ch] = right

		// Calculate current window size
		// right - left + 1 gives number of characters in window
		// Example: left=2, right=5 → 5-2+1 = 4 characters
		windowSize := right - left + 1

		// Update max if current window is larger
		if windowSize > max {
			max = windowSize
		}
	}

	return max
}

// ============================================================================
// SLIDING WINDOW PATTERNS:
// ============================================================================
//
// FIXED WINDOW:
//   sum := calculate first window
//   for i := k; i < n; i++ {
//       sum += arr[i]    // Add new element
//       sum -= arr[i-k]  // Remove old element
//       // process sum
//   }
//
//   Problems:
//   ✓ Maximum/minimum sum of k elements
//   ✓ Average of k elements
//   ✓ First negative in every window of size k
//
// VARIABLE WINDOW (expand right, shrink left):
//   left := 0
//   for right := 0; right < n; right++ {
//       // Add arr[right] to window
//
//       while window_condition_violated {
//           // Remove arr[left] from window
//           left++
//       }
//
//       // Process current window
//   }
//
//   Problems:
//   ✓ Longest substring without repeating characters
//   ✓ Minimum window substring
//   ✓ Longest substring with at most k distinct characters
//   ✓ Max consecutive ones after flipping k zeros
//
// ============================================================================
// COMMON PATTERNS:
// ============================================================================
//
// 1. FINDING MAXIMUM:
//    - Track max as you slide
//    - Update when current window beats max
//
// 2. FINDING MINIMUM:
//    - Track min as you slide
//    - Update when current window beats min
//
// 3. COUNTING:
//    - Use map to count frequency in window
//    - Add when right expands
//    - Subtract when left shrinks
//
// 4. CHECKING CONDITION:
//    - Expand right until condition met
//    - Shrink left while condition holds
//    - Track best window size
//
// ============================================================================

func main() {
	// Test fixed window
	fmt.Println("=== Fixed Window: Max Sum ===")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Array:", arr)
	fmt.Println("Max sum of 3 consecutive:", slidingWindow(arr, 3)) // 24 (7+8+9)

	// Demonstrate window sliding
	fmt.Println("\n=== Window Sliding Visualization ===")
	demo := []int{1, 2, 3, 4, 5}
	k := 3
	fmt.Printf("Array: %v, k=%d\n", demo, k)

	// First window
	sum := 0
	for i := 0; i < k; i++ {
		sum += demo[i]
	}
	fmt.Printf("Window 1: %v, sum=%d\n", demo[0:k], sum)

	// Slide windows
	for i := k; i < len(demo); i++ {
		sum += demo[i]
		sum -= demo[i-k]
		fmt.Printf("Window %d: %v, sum=%d (added %d, removed %d)\n",
			i-k+2, demo[i-k+1:i+1], sum, demo[i], demo[i-k])
	}

	// Test variable window
	fmt.Println("\n=== Variable Window: Longest Substring ===")
	fmt.Println("'abcabcbb':", longestSubString("abcabcbb")) // 3 ("abc")
	fmt.Println("'bbbbb':", longestSubString("bbbbb"))       // 1 ("b")
	fmt.Println("'pwwkew':", longestSubString("pwwkew"))     // 3 ("wke")
	fmt.Println("'abcdef':", longestSubString("abcdef"))     // 6 (whole string)

	// Demonstrate why last index matters
	fmt.Println("\n=== Why Store Last Index? ===")
	s := "abba"
	fmt.Printf("String: %s\n", s)
	fmt.Println("Without idx >= left check, we'd incorrectly shrink window")
	fmt.Println("With idx >= left check:", longestSubString(s)) // 2 ("ab" or "ba")
}
