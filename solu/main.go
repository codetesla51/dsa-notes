package main

import (
	"fmt"
	"regexp"
	"strings"
)

// ============================================================================
// SOLUTIONS - Common Algorithm Problems
// ============================================================================
// This file contains solutions to frequently asked coding interview questions
// Each function demonstrates key patterns and techniques
// ============================================================================

// nonRepeating finds the first character that appears only once
//
// COMPLEXITY:
// Time:  O(n) - two passes through string
// Space: O(k) - map stores k unique characters where k ≤ n
//
// KEY PATTERN: When order matters, loop original string not map
// - Maps are UNORDERED in Go
// - Must preserve sequence from original string
// - First pass: count frequencies
// - Second pass: find first with count=1
//
// EXAMPLE TRACE:
// Input: s = "leetcode"
//
// Pass 1: Build frequency map
//
//	'l' → countMap = {'l':1}
//	'e' → countMap = {'l':1, 'e':1}
//	'e' → countMap = {'l':1, 'e':2}
//	't' → countMap = {'l':1, 'e':2, 't':1}
//	'c' → countMap = {'l':1, 'e':2, 't':1, 'c':1}
//	'o' → countMap = {'l':1, 'e':2, 't':1, 'c':1, 'o':1}
//	'd' → countMap = {'l':1, 'e':2, 't':1, 'c':1, 'o':1, 'd':1}
//	'e' → countMap = {'l':1, 'e':3, 't':1, 'c':1, 'o':1, 'd':1}
//
// Pass 2: Find first non-repeating (loop ORIGINAL string)
//
//	Check 'l': countMap['l'] = 1 ✓
//	Return "l"
//
// WHY NOT LOOP THE MAP?
// If we did `for ch := range countMap`:
//   - Go maps are unordered
//   - Might return 't' or 'c' or 'o' (all have count=1)
//   - We need the FIRST non-repeating, not just ANY
func nonRepeating(s string) string {
	// Map to count frequency of each character
	countMap := make(map[rune]int)

	// First pass: count all characters
	for _, ch := range s {
		_, ok := countMap[ch]
		if ok {
			// Character already seen, increment count
			countMap[ch]++
		} else {
			// First occurrence of this character
			countMap[ch] = 1
		}
		// Could simplify to: countMap[ch]++
		// But explicit if/else shows the logic clearly
	}

	// CRITICAL: Second pass through ORIGINAL STRING
	// This preserves the order from the input
	// Maps in Go are unordered - iteration order is random
	for _, ch := range s {
		// O(1) lookup in map
		if countMap[ch] == 1 {
			// Found first character with count=1
			return string(ch)
		}
	}

	// All characters repeat (or empty string)
	return ""
}

// ============================================================================

// twoSum finds two numbers that add up to target
// Returns their indices
//
// COMPLEXITY:
// Time:  O(n) - single pass through array
// Space: O(n) - map stores up to n elements
//
// KEY PATTERN: Map eliminates nested loop
// - Brute force: O(n²) with nested loops
// - With map: O(n) single pass
// - Map stores "what we've seen" to avoid re-scanning
//
// EXAMPLE TRACE:
// Input: arr = [2, 7, 11, 15], tar = 9
//
// Initial:
// compMap = {}
//
// Iteration 1: i=0, n=2
//
//	comp = 9 - 2 = 7 (what number would sum with 2 to make 9?)
//	Is 7 in compMap? No
//	Add: compMap = {2: 0}
//
// Iteration 2: i=1, n=7
//
//	comp = 9 - 7 = 2 (what number would sum with 7 to make 9?)
//	Is 2 in compMap? Yes! At index 0
//	Return: [0, 1]
//
// Result: [0, 1] (because arr[0] + arr[1] = 2 + 7 = 9)
//
// WHY THIS WORKS:
// For each number n, we ask: "Have I already seen (target - n)?"
// If yes, we found our pair!
// If no, remember n for future iterations
//
// WHY map[int]int NOT map[int]bool?
// We need to return INDICES, not just detect existence
// Key = number value, Value = its index
func twoSum(arr []int, tar int) []int {
	// Map to store numbers we've seen
	// Key: the number value
	// Value: its index in the array
	compMap := make(map[int]int)

	// Single pass through array
	for i, n := range arr {
		// Calculate complement: what number + n = target?
		comp := tar - n
		// Example: if n=2 and tar=9, we need 7

		// Check if we've already seen the complement
		if j, ok := compMap[comp]; ok {
			// Found it! comp at index j, n at index i
			// comp + n = target
			return []int{j, i}
		}

		// Haven't found pair yet
		// Store current number for future iterations
		// "Hey future iterations, we saw n at index i"
		compMap[n] = i
	}

	// No pair found that sums to target
	return nil
}

// ============================================================================

// freqOcuring finds the most frequently occurring character
//
// COMPLEXITY:
// Time:  O(n) - single pass to count + O(k) to find max
// Space: O(k) - map stores k unique characters
//
// Note: Almost identical to isMoreOften in hashmaps/main.go
// Shows the same pattern in different context
func freqOcuring(s string) string {
	max := 0
	maxChar := ' '

	// Build frequency map
	countMap := make(map[rune]int)
	for _, ch := range s {
		_, ok := countMap[ch]
		if ok {
			countMap[ch]++
		} else {
			countMap[ch] = 1
		}
	}

	// Find character with maximum frequency
	// WARNING: Map iteration is UNORDERED
	// If multiple chars have same max frequency, result is random
	for key, value := range countMap {
		if value > max {
			max = value
			maxChar = key
		}
	}

	return string(maxChar)
}

// ============================================================================

// lastWord returns the length of the last word in a string
//
// COMPLEXITY:
// Time:  O(n) - string operations (trim, split) are O(n)
// Space: O(n) - split creates array of words
//
// EXAMPLE TRACE:
// Input: s = "   Hello World   "
//
// Step 1: TrimSpace
//
//	s = "Hello World"
//
// Step 2: Split by space
//
//	words = ["Hello", "World"]
//
// Step 3: Get last word
//
//	lastPart = words[2-1] = words[1] = "World"
//
// Step 4: Get length
//
//	len("World") = 5
//
// Result: 5
func lastWord(s string) int {
	// Remove leading and trailing whitespace
	// "  hello world  " → "hello world"
	s = strings.TrimSpace(s)

	// Split string by spaces into array of words
	// "hello world" → ["hello", "world"]
	words := strings.Split(s, " ")

	// Get the last word
	// words[len(words)-1] is last element
	lastPart := words[len(words)-1]

	// Return its length
	return len(lastPart)
}

// ============================================================================

// repeatedSubString checks if string is made of repeated pattern
//
// COMPLEXITY:
// Time:  O(n²) - outer loop O(n), inner repeat is O(n) worst case
// Space: O(n) - strings.Repeat creates new string
//
// EXAMPLE TRACE:
// Input: s = "abab"
//
// Iteration 1: i=1
//
//	chunk = s[0:1] = "a"
//	len(s)/len(chunk) = 4/1 = 4
//	strings.Repeat("a", 4) = "aaaa"
//	"aaaa" == "abab"? No
//
// Iteration 2: i=2
//
//	chunk = s[0:2] = "ab"
//	len(s)/len(chunk) = 4/2 = 2
//	strings.Repeat("ab", 2) = "abab"
//	"abab" == "abab"? Yes!
//	Return: true
//
// WHY i <= len(s)/2?
// A pattern can't be longer than half the string
// Example: "abcabc" (length 6)
//   - "abc" (length 3) can repeat: "abc"+"abc" = valid
//   - "abca" (length 4) can't repeat evenly in length 6
func repeatedSubString(s string) bool {
	// Try each possible pattern length
	// Pattern length from 1 to len(s)/2
	for i := 1; i <= len(s)/2; i++ {
		// Extract potential pattern
		chunk := s[0:i] // First i characters

		// Check if repeating this chunk equals original string
		// len(s)/len(chunk) = how many times to repeat
		if strings.Repeat(chunk, len(s)/len(chunk)) == s {
			return true
		}
	}

	// No repeating pattern found
	return false
}

// ============================================================================

// validPalindrome checks if string is a palindrome
// Ignores non-alphanumeric characters and case
//
// COMPLEXITY:
// Time:  O(n) - regex scan + reverse + comparison all O(n)
// Space: O(n) - cleaned string + byte slice
//
// # TWO POINTER PATTERN applied to string validation
//
// EXAMPLE TRACE:
// Input: s = "A man, a plan, a canal: Panama"
//
// Step 1: Remove non-alphanumeric
//
//	re matches: ", ", ", ", ", ", ": "
//	clean = "AmanaplanacanalPanama"
//
// Step 2: Lowercase and trim
//
//	s = "amanaplanacanalpanama"
//
// Step 3: Convert to byte slice and reverse
//
//	Initial: b = ['a','m','a','n',...,'m','a']
//	left=0, right=20
//
//	Swap process (two-pointer):
//	b[0] ↔ b[20]: 'a' ↔ 'a'
//	b[1] ↔ b[19]: 'm' ↔ 'm'
//	... continues until left >= right
//
//	Result: b = ['a','m','a','n',...,'m','a']
//
// Step 4: Compare
//
//	string(b) = "amanaplanacanalpanama"
//	s = "amanaplanacanalpanama"
//	Equal? Yes!
//
// Result: true
func validPalindrome(s string) bool {
	// Compile regex to match non-alphanumeric characters
	// [^...] means "NOT these characters"
	// a-zA-Z0-9 means letters and numbers
	// + means one or more
	re := regexp.MustCompile("[^a-zA-Z0-9]+")

	// Replace all non-alphanumeric with empty string (remove them)
	// "A man, a plan!" → "Amanaplana"
	clean := re.ReplaceAllString(s, "")

	// Convert to lowercase for case-insensitive comparison
	// Trim any remaining whitespace
	s = strings.ToLower(strings.TrimSpace(clean))

	// Two-pointer setup for reversal
	left := 0
	right := len(s) - 1

	// Convert to byte slice for mutation (strings are immutable)
	b := []byte(s)

	// Reverse the string using two pointers
	for left < right {
		// Swap characters at both ends
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}

	// If reversed equals original, it's a palindrome
	if string(b) == s {
		return true
	}
	return false
}

// ============================================================================

// revVowel reverses only the vowels in a string
//
// COMPLEXITY:
// Time:  O(n) - single pass with two pointers
// Space: O(n) - byte slice copy of string
//
// # TWO POINTER PATTERN with conditional advancement
//
// EXAMPLE TRACE:
// Input: s = "hello"
//
// Initial:
// b = ['h', 'e', 'l', 'l', 'o']
//
//	 ↑                   ↑
//	left               right
//
// voMap = {'a','e','i','o','u': true}
//
// Iteration 1:
//
//	b[left]='h' not vowel → left++ to 1
//
// Iteration 2:
//
//	b[left]='e' is vowel ✓
//	b[right]='o' is vowel ✓
//	Swap: b = ['h', 'o', 'l', 'l', 'e']
//	left++ to 2, right-- to 3
//
// Iteration 3:
//
//	b[left]='l' not vowel → left++ to 3
//
// Iteration 4:
//
//	b[right]='l' not vowel → right-- to 2
//
// Iteration 5:
//
//	left=3, right=2
//	left < right? No (3 < 2 is false)
//	Exit loop
//
// Result: "holle"
//
// WHY SEPARATE IF STATEMENTS?
// We check left and right independently
// Both might need to advance before swapping
// Only swap when BOTH are vowels
func revVowel(s string) string {
	// Map for O(1) vowel lookup
	// Using byte (not rune) because we only care about ASCII vowels
	voMap := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	// Two-pointer setup
	left := 0
	right := len(s) - 1

	// CRITICAL: Convert string to byte slice
	// Strings are immutable in Go - can't modify directly
	// []byte is mutable - we can swap elements
	b := []byte(s)

	// Process until pointers meet
	for left < right {
		// If left is NOT a vowel, advance it
		if !voMap[b[left]] {
			left++
		}

		// If right is NOT a vowel, retreat it
		if !voMap[b[right]] {
			right--
		}

		// If BOTH are vowels, swap them
		if voMap[b[left]] && voMap[b[right]] {
			// Go's simultaneous assignment - swaps in one line
			b[left], b[right] = b[right], b[left]

			// Move both pointers past the swapped vowels
			left++
			right--
		}
	}

	// Convert byte slice back to string
	return string(b)
}

// ============================================================================
// KEY PATTERNS DEMONSTRATED:
// ============================================================================
//
// 1. TWO POINTER PATTERN (revVowel, validPalindrome):
//    - Start at both ends
//    - Move toward center
//    - Process/swap elements
//
// 2. MAP TO ELIMINATE NESTED LOOPS (twoSum):
//    - Store "what we've seen"
//    - O(n²) → O(n) optimization
//
// 3. FREQUENCY COUNTING (nonRepeating, freqOcuring):
//    - map[T]int for counts
//    - Two passes: count, then find
//
// 4. ORDER PRESERVATION (nonRepeating):
//    - Maps are unordered
//    - Loop original sequence when order matters
//
// 5. STRING IMMUTABILITY (revVowel, validPalindrome):
//    - Convert to []byte to modify
//    - Convert back to string
//
// ============================================================================

func containsDuplicate(nums []int) bool {
	occur := make(map[int]int)

	for _, n := range nums {
		occur[n]++
	}
	fmt.Println(occur)
	for _, val := range occur {
		if val >= 2 {
			return true
		}
	}
	return false
}
func squart(x int) int {
	left := 0
	right := x
	if x < 2 {
		return x
	}
	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid
		if square == x {
			return mid
		} else if square < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
func main() {

	fmt.Println(squart(8)) // "Aa" (case preserved, only ASCII)
}
