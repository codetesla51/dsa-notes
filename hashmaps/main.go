package main

import "fmt"

// ============================================================================
// HASH MAPS (Maps in Go) - O(1) Lookup Power
// ============================================================================
//
// WHAT IS A HASH MAP?
// A data structure that maps keys to values using a hash function
// Think: magical librarian who instantly knows which shelf has your book
//
// HOW IT WORKS:
// 1. Hash function converts key → number (array index)
// 2. Store value at that index in underlying array
// 3. Lookup: hash(key) → index → value in O(1) time!
//
// VISUAL:
// ┌─────────────────────────────────────┐
// │  Key    │  Hash Function │  Index  │
// ├─────────────────────────────────────┤
// │ "cat"   │  hash("cat")   │   2     │
// │ "dog"   │  hash("dog")   │   5     │
// │ "bird"  │  hash("bird")  │   1     │
// └─────────────────────────────────────┘
//
// Underlying array:
// ┌───┬────────┬───┬────┬───┬────────┬───┐
// │ 0 │   1    │ 2 │ 3  │ 4 │   5    │ 6 │
// ├───┼────────┼───┼────┼───┼────────┼───┤
// │   │ bird:3 │cat:1│   │   │ dog:2  │   │
// └───┴────────┴───┴────┴───┴────────┴───┘
//
// WHY MAPS ARE POWERFUL:
// - Lookup: O(1) average case vs O(n) for arrays
// - Eliminates nested loops: O(n²) → O(n)
// - Perfect for counting, duplicate detection, caching
//
// ============================================================================
// MAP OPERATIONS:
// ============================================================================
// Lookup:  map[key]           → O(1) average
// Insert:  map[key] = value   → O(1) average
// Delete:  delete(map, key)   → O(1) average
// Check:   val, ok := map[key] → O(1) average
//
// GO MAP SPECIFICS:
// - Zero value for missing key (0 for int, false for bool, "" for string)
// - Use "comma ok" idiom to distinguish missing vs zero: val, ok := map[key]
// - Maps are UNORDERED (iteration order is random)
// - Maps are reference types (passed by reference)
//
// ============================================================================

// isSeenMoreThanOnce checks if any element appears more than once
// Also known as: "Contains Duplicate"
//
// COMPLEXITY:
// Time:  O(n) - single pass through array with O(1) map operations
// Space: O(n) - worst case, all elements unique → store all in map
//
// KEY PATTERN: map[T]bool for EXISTENCE checking
// - We only need to know "have I seen this before?"
// - Don't need to count HOW MANY times
// - Boolean is perfect for yes/no questions
//
// WHY THIS BEATS NESTED LOOPS:
// Naive approach (O(n²)):
//
//	for i := 0; i < len(arr); i++ {
//	    for j := i+1; j < len(arr); j++ {
//	        if arr[i] == arr[j] {
//	            return true
//	        }
//	    }
//	}
//
// For 1000 elements: ~500,000 comparisons
//
// Map approach (O(n)):
//
//	for _, n := range arr {
//	    if seen[n] { return true }
//	    seen[n] = true
//	}
//
// For 1000 elements: ~1000 operations
// 500x FASTER!
//
// EXAMPLE TRACE:
// Input: arr = [1, 2, 3, 2, 4]
//
// Initial state:
// isSeenMap = {}  (empty map)
//
// Iteration 1: n = 1
//
//	Check: Is 1 in isSeenMap?
//	_, ok := isSeenMap[1] → ok = false (not found)
//	Add to map: isSeenMap = {1: true}
//
// Iteration 2: n = 2
//
//	Check: Is 2 in isSeenMap?
//	_, ok := isSeenMap[2] → ok = false
//	Add to map: isSeenMap = {1: true, 2: true}
//
// Iteration 3: n = 3
//
//	Check: Is 3 in isSeenMap?
//	_, ok := isSeenMap[3] → ok = false
//	Add to map: isSeenMap = {1: true, 2: true, 3: true}
//
// Iteration 4: n = 2
//
//	Check: Is 2 in isSeenMap?
//	_, ok := isSeenMap[2] → ok = TRUE (found it!)
//	Return true immediately (duplicate detected!)
//
// # Never reaches iteration 5 - early exit optimization
//
// Result: true
func isSeenMoreThanOnce(arr []int) bool {
	// Create map to track which numbers we've encountered
	// Key: the number we saw
	// Value: true (just a marker, actual value doesn't matter)
	// We could use map[int]struct{} to save memory, but bool is clearer
	isSeenMap := make(map[int]bool)

	// Loop through each number in array
	for _, n := range arr {
		// The "comma ok" idiom - Go's way to check if key exists
		// First value (_) would be the value (we don't need it)
		// Second value (ok) is boolean: true if key exists, false otherwise
		if _, ok := isSeenMap[n]; ok {
			// We've seen this number before!
			// This is a duplicate, return immediately
			return true
		}

		// Mark this number as "seen"
		// The value (true) is just a placeholder
		// What matters is the KEY being present in the map
		isSeenMap[n] = true
	}

	// Checked all numbers, found no duplicates
	return false
}

// ============================================================================

// isMoreOften finds the character that appears most frequently
//
// COMPLEXITY:
// Time:  O(n) - single pass to count + O(k) to find max (k = unique chars)
// Space: O(k) - map stores k unique characters where k ≤ n
//
// KEY PATTERN: map[T]int for COUNTING
// - Need to track frequency of each character
// - map[T]bool wouldn't work (need actual counts, not just presence)
//
// EXAMPLE TRACE:
// Input: s = "aaabbc"
//
// Step 1: Build frequency map
//
// Iteration 1: ch = 'a'
//
//	isSeenMap['a']++ → 0++ = 1
//	isSeenMap = {'a': 1}
//
// Iteration 2: ch = 'a'
//
//	isSeenMap['a']++ → 1++ = 2
//	isSeenMap = {'a': 2}
//
// Iteration 3: ch = 'a'
//
//	isSeenMap['a']++ → 2++ = 3
//	isSeenMap = {'a': 3}
//
// Iteration 4: ch = 'b'
//
//	isSeenMap['b']++ → 0++ = 1
//	isSeenMap = {'a': 3, 'b': 1}
//
// Iteration 5: ch = 'b'
//
//	isSeenMap['b']++ → 1++ = 2
//	isSeenMap = {'a': 3, 'b': 2}
//
// Iteration 6: ch = 'c'
//
//	isSeenMap['c']++ → 0++ = 1
//	isSeenMap = {'a': 3, 'b': 2, 'c': 1}
//
// Step 2: Find character with maximum frequency
//
// Initial: maxCount = 0, maxChar = ' '
//
// Check 'a': val = 3
//
//	Is 3 >= 0? Yes
//	maxCount = 3, maxChar = 'a'
//
// Check 'b': val = 2
//
//	Is 2 >= 3? No
//	(no change)
//
// Check 'c': val = 1
//
//	Is 1 >= 3? No
//	(no change)
//
// Result: "a" (appeared 3 times)
//
// HOW map[rune]++ WORKS:
// When you do isSeenMap[ch]++:
// 1. Go checks if ch is in map
// 2. If NOT in map, returns zero value for int (which is 0)
// 3. Then increments: 0++ becomes 1
// 4. If already in map, gets current value and increments
// This is why we don't need to check "if ch in map" first!
func isMoreOften(s string) string {
	// Map to count frequency of each character
	// Key: the character (as rune for Unicode support)
	// Value: how many times it appears
	isSeenMap := make(map[rune]int)

	// Track the maximum frequency we've seen
	maxCount := 0

	// Track which character has the maximum frequency
	maxChar := ' ' // Space is placeholder (any rune works)

	// First pass: count all characters
	for _, ch := range s {
		// Increment count for this character
		// If ch not in map, Go returns 0 (zero value for int)
		// Then 0++ becomes 1 (first occurrence)
		// If ch already in map with count=2, then 2++ becomes 3
		isSeenMap[ch]++
	}

	// Second pass: find the character with highest count
	// WARNING: Go maps are UNORDERED
	// Iteration order is random (and can change between runs)
	// If multiple chars have same max frequency, result is unpredictable
	for key, val := range isSeenMap {
		// key = character (rune)
		// val = frequency (int)

		if val >= maxCount {
			// Found a character with higher (or equal) frequency
			maxCount = val // Update maximum frequency seen
			maxChar = key  // Update which character has this frequency
		}
	}

	// Convert rune back to string for return
	return string(maxChar)
}

// ============================================================================
// WHEN TO USE map[T]bool vs map[T]int:
// ============================================================================
//
// Use map[T]bool when:
// ✓ Checking membership/existence
// ✓ Detecting duplicates
// ✓ Set operations (union, intersection)
// ✓ Marking visited nodes in graph
//
// Example:
//   vowels := map[rune]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true}
//   if vowels[ch] { // O(1) check
//       // ch is a vowel
//   }
//
// Use map[T]int when:
// ✓ Counting frequencies
// ✓ Finding most/least common element
// ✓ Anagram detection
// ✓ Tracking occurrences
//
// Example:
//   freq := make(map[rune]int)
//   for _, ch := range s {
//       freq[ch]++  // Count each character
//   }
//
// ============================================================================
// CRITICAL PATTERN: Map Eliminates Nested Loops
// ============================================================================
//
// BEFORE (O(n²)):
//   for i := 0; i < len(arr); i++ {
//       for j := i+1; j < len(arr); j++ {
//           if arr[i] == arr[j] {
//               // found duplicate
//           }
//       }
//   }
//
// AFTER (O(n)):
//   seen := make(map[int]bool)
//   for _, n := range arr {
//       if seen[n] {
//           // found duplicate
//       }
//       seen[n] = true
//   }
//
// The map stores previous values, eliminating the need for inner loop!
//
// ============================================================================

func main() {
	// Test isSeenMoreThanOnce
	fmt.Println("=== Duplicate Detection ===")
	fmt.Println("[1,2,3,2,4]:", isSeenMoreThanOnce([]int{1, 2, 3, 2, 4})) // true
	fmt.Println("[1,2,3,4,5]:", isSeenMoreThanOnce([]int{1, 2, 3, 4, 5})) // false
	fmt.Println("[5,5]:", isSeenMoreThanOnce([]int{5, 5}))                // true
	fmt.Println("[]:", isSeenMoreThanOnce([]int{}))                       // false

	// Test isMoreOften
	fmt.Println("\n=== Most Frequent Character ===")
	fmt.Println("aaabbc:", isMoreOften("aaabbc")) // "a" (appears 3 times)
	fmt.Println("hello:", isMoreOften("hello"))   // "l" (appears 2 times)
	fmt.Println("abcdef:", isMoreOften("abcdef")) // random (all appear once)

	// Demonstrate map behavior
	fmt.Println("\n=== Map Behavior Demo ===")

	// Zero values
	m := make(map[string]int)
	fmt.Println("Missing key returns zero:", m["nonexistent"]) // 0

	// Comma ok idiom
	val, ok := m["nonexistent"]
	fmt.Println("Value:", val, "Exists:", ok) // 0 false

	m["test"] = 0
	val, ok = m["test"]
	fmt.Println("Value:", val, "Exists:", ok) // 0 true

	// Maps are unordered
	fmt.Println("\nMaps are UNORDERED - order can vary:")
	charMap := map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4}
	for k, v := range charMap {
		fmt.Printf("%c:%d ", k, v)
	}
	fmt.Println()
}
