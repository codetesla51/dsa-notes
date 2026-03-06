package main

import (
	"fmt"
)

// ============================================================================
// STRINGS - Immutability, Runes vs Bytes
// ============================================================================
//
// STRINGS ARE IMMUTABLE IN GO
// - You cannot modify a string directly
// - s[0] = 'x' → COMPILE ERROR
// - Must convert to []byte or []rune to modify
//
// BYTE vs RUNE:
// - byte: uint8 (8 bits) - represents ASCII characters (0-127)
// - rune: int32 (32 bits) - represents Unicode code points
//
// String "hello" in memory:
// ┌─────┬─────┬─────┬─────┬─────┐
// │ 'h' │ 'e' │ 'l' │ 'l' │ 'o' │  Each is 1 byte (ASCII)
// └─────┴─────┴─────┴─────┴─────┘
// len("hello") = 5 bytes
//
// String "a界" (1 ASCII + 1 Chinese):
// ┌─────┬─────┬─────┬─────┐
// │ 'a' │  界 (3 bytes) │  UTF-8 encoding
// └─────┴─────┴─────┴─────┘
// len("a界") = 4 bytes
// len([]rune("a界")) = 2 runes
//
// WHEN TO USE EACH:
// - range over string → gives runes (handles Unicode correctly)
// - []byte conversion → when you need to modify string
// - string indexing s[i] → gives bytes (can break Unicode chars)
//
// ============================================================================

// countVowel counts how many vowels in a string
//
// COMPLEXITY:
// Time:  O(n) - loop through string once
// Space: O(1) - map has constant 5 entries
//
// KEY PATTERN: map[T]bool for EXISTENCE checking
// - We don't need to count occurrences
// - Just need to know "is this a vowel?"
// - Boolean is perfect for yes/no questions
//
// EXAMPLE TRACE:
// Input: s = "hello"
//
// Initial:
// voMap = {'a':true, 'e':true, 'i':true, 'o':true, 'u':true}
// count = 0
//
// Iteration 1: ch = 'h'
//
//	Is 'h' in voMap? voMap['h'] returns false (zero value for bool)
//	count = 0
//
// Iteration 2: ch = 'e'
//
//	Is 'e' in voMap? voMap['e'] returns true
//	count++ → count = 1
//
// Iteration 3: ch = 'l'
//
//	Is 'l' in voMap? false
//	count = 1
//
// Iteration 4: ch = 'l'
//
//	Is 'l' in voMap? false
//	count = 1
//
// Iteration 5: ch = 'o'
//
//	Is 'o' in voMap? voMap['o'] returns true
//	count++ → count = 2
//
// Result: 2
//
// WHY USE MAP INSTEAD OF IF STATEMENTS?
// Without map:
//
//	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
//	    count++
//	}
//
// This is O(5) per character (5 comparisons in worst case)
//
// With map:
//
//	if voMap[ch] { count++ }
//
// This is O(1) per character (single hash lookup)
//
// For 5 vowels, difference is small. For larger sets, map is much better!
func countVowel(s string) int {
	// Create map for O(1) vowel lookup
	// Key: the vowel character (as rune)
	// Value: true (we only care about presence, not count)
	voMap := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	// Counter to accumulate total vowels found
	count := 0

	// range on string gives us RUNES, not bytes
	// This correctly handles Unicode characters
	// Example: "café" with é is handled properly
	for _, ch := range s {
		// Map lookup: O(1) average case
		// If key doesn't exist, Go returns zero value (false for bool)
		// If key exists with value true, condition passes
		if voMap[ch] {
			count++ // Found a vowel, increment counter
		}
	}

	return count
}

// ============================================================================

// isAnan checks if two strings are anagrams
// Anagram: same characters with same frequencies
// Example: "listen" and "silent"
//
// COMPLEXITY:
// Time:  O(n + m) where n=len(a), m=len(b)
// Space: O(k + j) where k=unique chars in a, j=unique chars in b
//
// KEY PATTERN: map[T]int for COUNTING
// - Need to track frequency of each character
// - map[T]bool wouldn't work (need actual counts)
//
// EXAMPLE TRACE:
// Input: a = "listen", b = "silent"
//
// Step 1: Build frequency map for "listen"
// Iteration by iteration:
//
//	'l' → map1 = {'l':1}
//	'i' → map1 = {'l':1, 'i':1}
//	's' → map1 = {'l':1, 'i':1, 's':1}
//	't' → map1 = {'l':1, 'i':1, 's':1, 't':1}
//	'e' → map1 = {'l':1, 'i':1, 's':1, 't':1, 'e':1}
//	'n' → map1 = {'l':1, 'i':1, 's':1, 't':1, 'e':1, 'n':1}
//
// Step 2: Build frequency map for "silent"
//
//	's' → map2 = {'s':1}
//	'i' → map2 = {'s':1, 'i':1}
//	'l' → map2 = {'s':1, 'i':1, 'l':1}
//	'e' → map2 = {'s':1, 'i':1, 'l':1, 'e':1}
//	'n' → map2 = {'s':1, 'i':1, 'l':1, 'e':1, 'n':1}
//	't' → map2 = {'s':1, 'i':1, 'l':1, 'e':1, 'n':1, 't':1}
//
// Step 3: Compare map1 against map2
//
//	'l': map2['l'] = 1, matches map1['l'] = 1 ✓
//	'i': map2['i'] = 1, matches map1['i'] = 1 ✓
//	's': map2['s'] = 1, matches map1['s'] = 1 ✓
//	't': map2['t'] = 1, matches map1['t'] = 1 ✓
//	'e': map2['e'] = 1, matches map1['e'] = 1 ✓
//	'n': map2['n'] = 1, matches map1['n'] = 1 ✓
//
// Step 4: Compare map2 against map1
//
//	(All match, same as step 3)
//
// Result: true (is anagram)
//
// WHY CHECK BOTH DIRECTIONS?
// Example: a = "listen", b = "listenx"
// First loop: all chars in "listen" are in "listenx" ✓
// Second loop: 'x' in "listenx" is NOT in "listen" ✗
// Catches extra characters!
func isAnan(a, b string) bool {
	// Create frequency maps for both strings
	// map[rune]int because we need to COUNT occurrences
	// Using rune handles Unicode correctly
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	// Build frequency map for first string
	for _, ch := range a {
		map1[ch]++
		// How map1[ch]++ works:
		// 1. If ch not in map, default int value is 0
		// 2. 0++ becomes 1 (first occurrence)
		// 3. If ch already in map with count=2, then 2++ becomes 3
	}

	// Build frequency map for second string
	for _, ch := range b {
		map2[ch]++
	}

	// Check: Does every character in map1 appear with same frequency in map2?
	for key, val := range map1 {
		// val = how many times 'key' appears in string a
		// map2[key] = how many times 'key' appears in string b

		if map2[key] != val {
			// Frequency mismatch! Not an anagram
			// Note: if map2 doesn't have key, returns 0
			return false
		}
	}

	// Check: Does every character in map2 appear with same frequency in map1?
	// This catches cases where b has extra characters not in a
	for key, val := range map2 {
		if map1[key] != val {
			return false
		}
	}

	// All characters match with same frequencies
	return true
}

// ============================================================================
// KEY STRING CONCEPTS:
// ============================================================================
//
// 1. IMMUTABILITY:
//    s := "hello"
//    s[0] = 'H'  // ❌ COMPILE ERROR
//
//    Instead:
//    b := []byte(s)
//    b[0] = 'H'
//    s = string(b)  // ✓ Works
//
// 2. RUNE vs BYTE:
//    s := "Go言語"
//    len(s) → 8 bytes (Go=2, 言=3, 語=3)
//    len([]rune(s)) → 4 runes
//
//    for i := 0; i < len(s); i++ {
//        // s[i] gives bytes - can break Unicode!
//    }
//
//    for _, ch := range s {
//        // ch is rune - handles Unicode correctly!
//    }
//
// 3. ZERO VALUE:
//    var s string → s = "" (empty string, not nil)
//
// 4. CONCATENATION:
//    s1 + s2 creates new string (immutable)
//    Use strings.Builder for efficient concatenation in loops
//
// ============================================================================

func main() {
	// Test countVowel
	fmt.Println("=== Count Vowels ===")
	fmt.Println("hello:", countVowel("hello"))             // 2 (e, o)
	fmt.Println("AEIOU:", countVowel("AEIOU"))             // 0 (uppercase not in map)
	fmt.Println("programming:", countVowel("programming")) // 3 (o, a, i)

	// Test isAnan
	fmt.Println("\n=== Anagram Check ===")
	fmt.Println("listen, silent:", isAnan("listen", "silent")) // true
	fmt.Println("hello, olleh:", isAnan("hello", "olleh"))     // true
	fmt.Println("hello, world:", isAnan("hello", "world"))     // false
	fmt.Println("a, aa:", isAnan("a", "aa"))                   // false (different lengths)

	// Demonstrate byte vs rune
	fmt.Println("\n=== Byte vs Rune Demo ===")
	s := "Go言語"
	fmt.Println("String:", s)
	fmt.Println("Byte length:", len(s))         // 8 bytes
	fmt.Println("Rune length:", len([]rune(s))) // 4 runes

	fmt.Println("\nIterating as bytes (WRONG for Unicode):")
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %d\n", i, s[i])
	}

	fmt.Println("\nIterating as runes (CORRECT for Unicode):")
	for i, ch := range s {
		fmt.Printf("Index %d: %c (rune value: %d)\n", i, ch, ch)
	}
}
