# Data Structures & Algorithms - Comprehensive Study Notes

## Table of Contents
1. [Big O Notation](#big-o-notation)
2. [Arrays](#arrays)
3. [Strings](#strings)
4. [Hash Tables](#hash-tables)
5. [Linked Lists](#linked-lists)
6. [Key Patterns Summary](#key-patterns-summary)

---

## Big O Notation

### Real World Analogy
Think of Big O as measuring how much longer a task takes as you scale it up:
- **O(1)** - Looking up a word in a dictionary if you know the exact page number
- **O(n)** - Reading every page in a book to find a specific word
- **O(n²)** - Comparing every person in a room with every other person (handshakes)

### Plain English Explanation
Big O notation describes how the runtime or space requirements of an algorithm grow as the input size increases. We focus on the **worst case** and ignore constants.

### How to Read Complexity from Code

#### Example 1: O(n) - Linear Time
**From bigo/main.go:5**
```go
func findTar(arr []int, tar int) bool {
    // Iterate through entire array once
    for _, n := range arr {
        // Single comparison - O(1) operation
        if n == tar {
            return true
        }
    }
    return false
}
```

**Line-by-line explanation:**
```go
func findTar(arr []int, tar int) bool {
    // Input: arr = [1, 2, 3, 4], tar = 4
    
    for _, n := range arr {
        // This loop runs n times where n = len(arr)
        // Iteration 1: n = 1, check if 1 == 4? No
        // Iteration 2: n = 2, check if 2 == 4? No
        // Iteration 3: n = 3, check if 3 == 4? No
        // Iteration 4: n = 4, check if 4 == 4? Yes!
        
        if n == tar {
            return true  // Found it! Return immediately
        }
    }
    return false  // Searched entire array, not found
}
```

**Step-by-step trace:**
```
Input: arr = [1, 2, 3, 4], tar = 4

Step 1: n = 1
        Compare: 1 == 4? false
        Continue...

Step 2: n = 2
        Compare: 2 == 4? false
        Continue...

Step 3: n = 3
        Compare: 3 == 4? false
        Continue...

Step 4: n = 4
        Compare: 4 == 4? true
        Return true
```

**Complexity: O(n)**
- We potentially check every element in the array
- If array doubles in size, worst case doubles the operations
- Best case: O(1) if found immediately
- Worst case: O(n) if not found or at end

---

## Arrays

### Real World Analogy
Arrays are like numbered parking spots in a parking garage:
- Each spot has a number (index)
- All spots are in a continuous row
- You can instantly drive to spot #47 without passing spots 1-46
- But if you want to insert a car between spot #3 and #4, you have to move every car from #4 onward down one spot

### Plain English Explanation
Arrays store elements in **contiguous memory locations**. Each element can be accessed in constant time using its index because the computer calculates the exact memory address using simple math: `base_address + (index * element_size)`.

### ASCII Diagram
```
Array in Memory:
┌───────────────────────────────────────────────┐
│  Index:  │  0  │  1  │  2  │  3  │  4  │  5  │
├───────────────────────────────────────────────┤
│  Value:  │  1  │  2  │  3  │  4  │  5  │  6  │
├───────────────────────────────────────────────┤
│ Address: │1000 │1004 │1008 │1012 │1016 │1020 │
└───────────────────────────────────────────────┘

To access arr[3]:
memory_address = base_address + (index * 4 bytes)
memory_address = 1000 + (3 * 4) = 1012
```

### Example 1: Array Multiplication
**From arrays/main.go:17**

```go
func mulArr(arr []int) []int {
    newArr := []int{}
    for i, n := range arr {
        mul := n * i
        newArr = append(newArr, mul)
    }
    return newArr
}
```

**Line-by-line explanation:**
```go
func mulArr(arr []int) []int {
    // Purpose: Multiply each element by its index
    // Input: [1, 2, 3, 4, 5]
    
    newArr := []int{}
    // Create empty slice to store results
    // Why new array? We're building results, not modifying in place
    
    for i, n := range arr {
        // i = index (0, 1, 2, 3, 4)
        // n = value at that index
        // range gives us both automatically
        
        mul := n * i
        // Multiply value by its position
        // 1*0=0, 2*1=2, 3*2=6, 4*3=12, 5*4=20
        
        newArr = append(newArr, mul)
        // append adds element to end of slice
        // Go automatically resizes if needed
    }
    return newArr  // [0, 2, 6, 12, 20]
}
```

**Step-by-step trace:**
```
Input: arr = [1, 2, 3, 4, 5]

Initial state:
arr =    [1, 2, 3, 4, 5]
newArr = []

Iteration 1: i=0, n=1
    mul = 1 * 0 = 0
    newArr = [0]

Iteration 2: i=1, n=2
    mul = 2 * 1 = 2
    newArr = [0, 2]

Iteration 3: i=2, n=3
    mul = 3 * 2 = 6
    newArr = [0, 2, 6]

Iteration 4: i=3, n=4
    mul = 4 * 3 = 12
    newArr = [0, 2, 6, 12]

Iteration 5: i=4, n=5
    mul = 5 * 4 = 20
    newArr = [0, 2, 6, 12, 20]

Return: [0, 2, 6, 12, 20]
```

**Complexity: O(n) time, O(n) space**
- Time: Loop through array once (n iterations)
- Space: Create new array of same size

### Example 2: Two Pointer Array Reversal (Commented Code)
**From arrays/main.go:5-16**

```go
func revArr(arr []int) []int {
    left := 0               // Start pointer at beginning
    right := len(arr) - 1   // Start pointer at end
    
    for left < right {      // While pointers haven't met
        // Swap elements at left and right
        temp := arr[left]       // Store left value temporarily
        arr[left] = arr[right]  // Overwrite left with right
        arr[right] = temp       // Put original left into right
        
        left++   // Move left pointer right
        right--  // Move right pointer left
    }
    return arr
}
```

**Step-by-step trace:**
```
Input: arr = [1, 2, 3, 4, 5]

Initial:
arr = [1, 2, 3, 4, 5]
       ↑           ↑
      left       right

Step 1: Swap arr[0] and arr[4]
    temp = 1
    arr[0] = 5
    arr[4] = 1
    arr = [5, 2, 3, 4, 1]
    left++ → 1, right-- → 3

Step 2: Swap arr[1] and arr[3]
    temp = 2
    arr[1] = 4
    arr[3] = 2
    arr = [5, 4, 3, 2, 1]
    left++ → 2, right-- → 2

Step 3: left == right (both at index 2)
    Loop exits (left < right is false)

Result: [5, 4, 3, 2, 1]
```

**Complexity: O(n) time, O(1) space**
- Time: Visit each element once (n/2 swaps, but O(n))
- Space: Only 3 variables (left, right, temp) regardless of array size

### Array Complexity Table
| Operation | Time Complexity | Reason |
|-----------|----------------|---------|
| Access by index | O(1) | Direct memory calculation |
| Search (unsorted) | O(n) | Must check each element |
| Insert at end | O(1) amortized | Append is constant* |
| Insert at beginning | O(n) | Must shift all elements right |
| Delete at end | O(1) | Just reduce length |
| Delete at beginning | O(n) | Must shift all elements left |
| Delete at middle | O(n) | Must shift remaining elements |

*Go slices auto-resize with amortized O(1) append

---

## Strings

### Real World Analogy
Strings are like messages carved in stone:
- You can read any character instantly
- You can make a copy with changes
- But you can't modify the original stone - it's immutable
- Each letter might be a single symbol (byte) or a complex hieroglyph (rune)

### Plain English Explanation
In Go, strings are **immutable** sequences of bytes. To modify a string, you must create a new one. Go distinguishes between:
- **byte**: 8-bit value (ASCII characters)
- **rune**: int32 value (Unicode code point, can represent any character including emoji)

### ASCII Diagram
```
String "hello" in memory:

As a string (immutable):
┌─────┬─────┬─────┬─────┬─────┐
│ 'h' │ 'e' │ 'l' │ 'l' │ 'o' │  ← Cannot modify
└─────┴─────┴─────┴─────┴─────┘

Converted to []byte (mutable):
┌─────┬─────┬─────┬─────┬─────┐
│ 104 │ 101 │ 108 │ 108 │ 111 │  ← Can modify
└─────┴─────┴─────┴─────┴─────┘

Why convert?
string → []byte → modify → string
         ↑                   ↑
    Can swap elements    Convert back
```

### Byte vs Rune
```
String: "a界" (1 ASCII + 1 Chinese character)

Using bytes (len = 4):
Index: 0    1    2    3
Byte:  97   231  149  140
       'a'  └─────┬─────┘
               '界' (3 bytes in UTF-8)

Using runes (len = 2):
Index: 0      1
Rune:  97     30028
       'a'    '界'
```

### Example 1: Count Vowels
**From strings/main.go:7**

```go
func countVowel(s string) int {
    voMap := map[rune]bool{
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
    }
    count := 0
    for _, ch := range s {
        if voMap[ch] {
            count++
        }
    }
    return count
}
```

**Line-by-line explanation:**
```go
func countVowel(s string) int {
    // Purpose: Count how many vowels in a string
    
    voMap := map[rune]bool{
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
    }
    // Why map[rune]bool? We only care if it EXISTS (is a vowel)
    // Don't need to count occurrences, just check membership
    // O(1) lookup vs O(5) if we used array and looped
    
    count := 0
    // Accumulator to track total vowels found
    
    for _, ch := range s {
        // range on string gives us runes, not bytes
        // Handles Unicode correctly (e.g., "café" with é)
        // _ ignores index because we don't need it
        
        if voMap[ch] {
            // Map lookup is O(1)
            // If key doesn't exist, returns false (zero value for bool)
            // If key exists with value true, condition passes
            
            count++  // Found a vowel, increment
        }
    }
    return count
}
```

**Step-by-step trace:**
```
Input: s = "hello"

Initial state:
voMap = {'a':true, 'e':true, 'i':true, 'o':true, 'u':true}
count = 0

Iteration 1: ch = 'h'
    Is 'h' in voMap? No (returns false)
    count = 0

Iteration 2: ch = 'e'
    Is 'e' in voMap? Yes (returns true)
    count = 1

Iteration 3: ch = 'l'
    Is 'l' in voMap? No
    count = 1

Iteration 4: ch = 'l'
    Is 'l' in voMap? No
    count = 1

Iteration 5: ch = 'o'
    Is 'o' in voMap? Yes
    count = 2

Return: 2
```

**Complexity: O(n) time, O(1) space**
- Time: Loop through string once (n characters)
- Space: Map has constant 5 entries regardless of input size

### Example 2: Anagram Check
**From strings/main.go:23**

```go
func isAnan(a, b string) bool {
    map1 := make(map[rune]int)
    map2 := make(map[rune]int)
    for _, ch := range a {
        map1[ch]++
    }
    for _, ch := range b {
        map2[ch]++
    }
    for key, val := range map1 {
        if map2[key] != val {
            return false
        }
    }
    for key, val := range map2 {
        if map1[key] != val {
            return false
        }
    }
    return true
}
```

**Line-by-line explanation:**
```go
func isAnan(a, b string) bool {
    // Purpose: Check if two strings are anagrams
    // Anagram: same characters with same frequencies
    // Example: "listen" and "silent"
    
    map1 := make(map[rune]int)
    map2 := make(map[rune]int)
    // Why map[rune]int? We need to COUNT occurrences
    // Can't use bool - need to know "appeared twice" vs "appeared once"
    
    for _, ch := range a {
        map1[ch]++
        // If ch not in map, default int value is 0
        // 0++ becomes 1 (first occurrence)
        // Subsequent hits increment: 1++, 2++, etc.
    }
    
    for _, ch := range b {
        map2[ch]++
        // Build frequency map for second string
    }
    
    for key, val := range map1 {
        // Check every character in first string
        if map2[key] != val {
            // Does second string have same count?
            // If map2 doesn't have key, returns 0
            // So if val=2 but map2[key]=0, not anagram
            return false
        }
    }
    
    for key, val := range map2 {
        // Must check both directions!
        // What if b="listenx" and a="listen"?
        // First loop passes (all of "listen" chars match)
        // But b has extra 'x', so this loop catches it
        if map1[key] != val {
            return false
        }
    }
    
    return true  // All characters match with same frequencies
}
```

**Step-by-step trace:**
```
Input: a = "listen", b = "silent"

Step 1: Build map1 from "listen"
    'l': 1
    'i': 1
    's': 1
    't': 1
    'e': 1
    'n': 1

Step 2: Build map2 from "silent"
    's': 1
    'i': 1
    'l': 1
    'e': 1
    'n': 1
    't': 1

Step 3: Check map1 against map2
    'l' in map2? Yes, count = 1 ✓
    'i' in map2? Yes, count = 1 ✓
    's' in map2? Yes, count = 1 ✓
    't' in map2? Yes, count = 1 ✓
    'e' in map2? Yes, count = 1 ✓
    'n' in map2? Yes, count = 1 ✓

Step 4: Check map2 against map1
    's' in map1? Yes, count = 1 ✓
    'i' in map1? Yes, count = 1 ✓
    'l' in map1? Yes, count = 1 ✓
    'e' in map1? Yes, count = 1 ✓
    'n' in map1? Yes, count = 1 ✓
    't' in map1? Yes, count = 1 ✓

Return: true (is anagram)
```

**Complexity: O(n + m) time, O(n + m) space**
- Time: Loop through string a once (n), string b once (m), then both maps
- Space: Two maps storing unique characters from both strings

### Example 3: Reverse Vowels
**From solu/main.go:93**

```go
func revVowel(s string) string {
    voMap := map[byte]bool{
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
    }
    left := 0
    right := len(s) - 1
    b := []byte(s)
    for left < right {
        if !voMap[b[left]] {
            left++
        }
        if !voMap[b[right]] {
            right--
        }
        if voMap[b[left]] && voMap[b[right]] {
            b[left], b[right] = b[right], b[left]
            left++
            right--
        }
    }
    return string(b)
}
```

**Line-by-line explanation:**
```go
func revVowel(s string) string {
    // Purpose: Reverse only the vowels in a string
    // Example: "hello" → "holle"
    
    voMap := map[byte]bool{
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
    }
    // Using byte not rune because we only care about ASCII vowels
    
    left := 0
    right := len(s) - 1
    // Two pointer technique: start at both ends, move toward middle
    
    b := []byte(s)
    // CRITICAL: Convert string to byte slice
    // Why? Strings are immutable in Go - can't modify directly
    // []byte is mutable - we can swap elements
    
    for left < right {
        // Continue until pointers meet/cross
        
        if !voMap[b[left]] {
            // Current left is NOT a vowel
            left++
            // Move left pointer right to find a vowel
            // Don't swap, just advance
        }
        
        if !voMap[b[right]] {
            // Current right is NOT a vowel
            right--
            // Move right pointer left to find a vowel
        }
        
        if voMap[b[left]] && voMap[b[right]] {
            // Both pointers are at vowels - ready to swap!
            b[left], b[right] = b[right], b[left]
            // Go's simultaneous assignment - swaps in one line
            
            left++   // Move past swapped vowel
            right--  // Move past swapped vowel
        }
    }
    
    return string(b)
    // Convert byte slice back to string
}
```

**Step-by-step trace:**
```
Input: s = "hello"

Initial:
b = ['h', 'e', 'l', 'l', 'o']
      ↑                   ↑
     left               right

Step 1: left=0, right=4
    b[0]='h' → not vowel → left++
    
Step 2: left=1, right=4
    b[1]='e' → vowel ✓
    b[4]='o' → vowel ✓
    Swap! b = ['h', 'o', 'l', 'l', 'e']
    left++ → 2, right-- → 3

Step 3: left=2, right=3
    b[2]='l' → not vowel → left++
    
Step 4: left=3, right=3
    left == right → loop exits

Result: "holle"
```

**Complexity: O(n) time, O(n) space**
- Time: Single pass with two pointers (each element visited once)
- Space: O(n) for byte slice copy of string

### Example 4: Valid Palindrome
**From solu/main.go:76**

```go
func validPalindrome(s string) bool {
    re := regexp.MustCompile("[^a-zA-Z0-9]+")
    clean := re.ReplaceAllString(s, "")
    s = strings.ToLower(strings.TrimSpace(clean))
    left := 0
    right := len(s) - 1
    b := []byte(s)
    for left < right {
        b[left], b[right] = b[right], b[left]
        left++
        right--
    }
    if string(b) == s {
        return true
    }
    return false
}
```

**Line-by-line explanation:**
```go
func validPalindrome(s string) bool {
    // Purpose: Check if string is palindrome (reads same forwards/backwards)
    // Must ignore non-alphanumeric and case
    
    re := regexp.MustCompile("[^a-zA-Z0-9]+")
    // Compile regex pattern
    // [^...] means "NOT these characters"
    // a-zA-Z0-9 means letters and numbers
    // + means one or more
    // So: "match any non-alphanumeric characters"
    
    clean := re.ReplaceAllString(s, "")
    // Replace all matches with empty string (remove them)
    // "A man, a plan!" → "Amanaplana"
    
    s = strings.ToLower(strings.TrimSpace(clean))
    // TrimSpace removes leading/trailing whitespace
    // ToLower converts to lowercase for case-insensitive comparison
    // "Amanaplana" → "amanaplana"
    
    left := 0
    right := len(s) - 1
    b := []byte(s)
    // Convert to byte slice for mutation
    
    for left < right {
        b[left], b[right] = b[right], b[left]
        // Reverse the string by swapping
        left++
        right--
    }
    
    if string(b) == s {
        // If reversed string equals original, it's a palindrome
        return true
    }
    return false
}
```

**Step-by-step trace:**
```
Input: s = "A man, a plan, a canal: Panama"

Step 1: Remove non-alphanumeric
    clean = "AmanaplanacanalPanama"

Step 2: Lowercase and trim
    s = "amanaplanacanalpanama"

Step 3: Reverse using two pointers
    b = ['a','m','a','n','a','p','l','a','n','a','c','a','n','a','l','p','a','n','a','m','a']
    
    After swapping:
    b = ['a','m','a','n','a','p','l','a','n','a','c','a','n','a','l','p','a','n','a','m','a']
    
Step 4: Compare
    string(b) == s?
    "amanaplanacanalpanama" == "amanaplanacanalpanama"
    true

Return: true
```

**Complexity: O(n) time, O(n) space**
- Time: Regex scan O(n), reverse O(n), comparison O(n) → O(n)
- Space: O(n) for cleaned string, byte slice

---

## Hash Tables

### Real World Analogy
A hash table is like a library with a magical librarian:
- You tell the librarian a book title (key)
- The librarian instantly knows which shelf it's on (hash function calculates index)
- You walk directly to that shelf (O(1) access)
- No need to search through every shelf linearly

### Plain English Explanation
Hash tables (maps in Go) store key-value pairs. A hash function converts keys into array indices. This allows **O(1) average-case** lookup, insertion, and deletion.

### ASCII Diagram
```
Hash Table: map[string]int

How it works internally:
┌─────────────────────────────────────┐
│  Key    │  Hash Function │  Index  │
├─────────────────────────────────────┤
│ "cat"   │  hash("cat")   │   2     │
│ "dog"   │  hash("dog")   │   5     │
│ "bird"  │  hash("bird")  │   1     │
└─────────────────────────────────────┘

Underlying array:
┌───┬────────┬───┬────┬───┬────────┬───┐
│ 0 │   1    │ 2 │ 3  │ 4 │   5    │ 6 │
├───┼────────┼───┼────┼───┼────────┼───┤
│   │ bird:3 │cat:1│   │   │ dog:2  │   │
└───┴────────┴───┴────┴───┴────────┴───┘

To find "cat":
1. hash("cat") = 2
2. Go to index 2
3. Return value: 1
Time: O(1) ✓
```

### map[T]bool vs map[T]int

**Use map[T]bool when:**
- You only care about **existence** (is it in the set?)
- Examples: checking for duplicates, vowel detection, membership testing

**Use map[T]int when:**
- You need to **count occurrences**
- Examples: character frequency, word count, anagram detection

### Example 1: Detect Duplicates
**From hashmaps/main.go:5**

```go
func isSeenMoreThanOnce(arr []int) bool {
    isSeenMap := make(map[int]bool)
    for _, n := range arr {
        if _, ok := isSeenMap[n]; ok {
            return true
        }
        isSeenMap[n] = true
    }
    return false
}
```

**Line-by-line explanation:**
```go
func isSeenMoreThanOnce(arr []int) bool {
    // Purpose: Return true if any element appears more than once
    // Also known as: "Contains duplicate"
    
    isSeenMap := make(map[int]bool)
    // Why map[int]bool? We only need to know "have I seen this?"
    // We don't care HOW MANY times - just yes/no
    // bool is perfect for binary state
    
    for _, n := range arr {
        // Loop through each number
        
        if _, ok := isSeenMap[n]; ok {
            // Check if n already exists in map
            // The "comma ok" idiom:
            //   - First return value (_) is the value (we ignore it)
            //   - Second return value (ok) is boolean: true if key exists
            // If we've seen this number before, it's a duplicate!
            
            return true  // Found duplicate, exit immediately
        }
        
        isSeenMap[n] = true
        // Mark this number as "seen"
        // Value doesn't matter (true is just placeholder)
        // The KEY is what we care about
    }
    
    return false  // Checked all numbers, no duplicates found
}
```

**Step-by-step trace:**
```
Input: arr = [1, 2, 3, 2, 4]

Initial:
isSeenMap = {}

Iteration 1: n = 1
    Is 1 in isSeenMap? No
    Add: isSeenMap = {1: true}

Iteration 2: n = 2
    Is 2 in isSeenMap? No
    Add: isSeenMap = {1: true, 2: true}

Iteration 3: n = 3
    Is 3 in isSeenMap? No
    Add: isSeenMap = {1: true, 2: true, 3: true}

Iteration 4: n = 2
    Is 2 in isSeenMap? Yes! ← Found duplicate
    Return: true

(Never reaches iterations 5, exits early)
```

**Why this beats nested loops:**
```
Naive approach (O(n²)):
for i := 0; i < len(arr); i++ {
    for j := i+1; j < len(arr); j++ {
        if arr[i] == arr[j] {
            return true
        }
    }
}
// For 1000 elements: ~500,000 comparisons

Map approach (O(n)):
// For 1000 elements: ~1000 operations
```

**Complexity: O(n) time, O(n) space**
- Time: Single pass through array, O(1) map operations
- Space: Worst case, all elements unique → map stores all n elements

### Example 2: Character Frequency
**From hashmaps/main.go:15**

```go
func isMoreOften(s string) string {
    isSeenMap := make(map[rune]int)
    maxCount := 0
    maxChar := ' '
    for _, ch := range s {
        isSeenMap[ch]++
    }
    for key, val := range isSeenMap {
        if val >= maxCount {
            maxCount = val
            maxChar = key
        }
    }
    return string(maxChar)
}
```

**Line-by-line explanation:**
```go
func isMoreOften(s string) string {
    // Purpose: Find the character that appears most frequently
    
    isSeenMap := make(map[rune]int)
    // Why map[rune]int? We need to COUNT occurrences
    // map[rune]bool wouldn't work - need actual counts
    
    maxCount := 0
    maxChar := ' '
    // Track the highest frequency seen and which character has it
    
    for _, ch := range s {
        isSeenMap[ch]++
        // If ch not in map, defaults to 0, then 0++ = 1
        // If ch already in map with count=2, then 2++ = 3
        // This is why map[rune]int - we're accumulating counts
    }
    
    for key, val := range isSeenMap {
        // Loop through the map (order is random in Go maps!)
        // key = character, val = how many times it appeared
        
        if val >= maxCount {
            // Found a character with higher (or equal) frequency
            maxCount = val     // Update max frequency
            maxChar = key      // Update which character has it
        }
    }
    
    return string(maxChar)
    // Convert rune back to string for return
}
```

**Step-by-step trace:**
```
Input: s = "aaabbc"

Step 1: Build frequency map
    Iteration 1: ch='a', isSeenMap={'a':1}
    Iteration 2: ch='a', isSeenMap={'a':2}
    Iteration 3: ch='a', isSeenMap={'a':3}
    Iteration 4: ch='b', isSeenMap={'a':3, 'b':1}
    Iteration 5: ch='b', isSeenMap={'a':3, 'b':2}
    Iteration 6: ch='c', isSeenMap={'a':3, 'b':2, 'c':1}

Step 2: Find maximum
    Initial: maxCount=0, maxChar=' '
    
    Check 'a':3
        3 >= 0? Yes
        maxCount=3, maxChar='a'
    
    Check 'b':2
        2 >= 3? No
        (no change)
    
    Check 'c':1
        1 >= 3? No
        (no change)

Return: "a"
```

**Complexity: O(n) time, O(k) space**
- Time: O(n) to build map + O(k) to find max where k = unique chars
- Space: O(k) for map, k ≤ n (unique characters)

### Example 3: Two Sum
**From solu/main.go:26**

```go
func twoSum(arr []int, tar int) []int {
    compMap := make(map[int]int)
    for i, n := range arr {
        comp := tar - n
        if j, ok := compMap[comp]; ok {
            return []int{j, i}
        }
        compMap[n] = i
    }
    return nil
}
```

**Line-by-line explanation:**
```go
func twoSum(arr []int, tar int) []int {
    // Purpose: Find two numbers that add up to target
    // Return their indices
    // Example: arr=[2,7,11,15], tar=9 → [0,1] because 2+7=9
    
    compMap := make(map[int]int)
    // Why map[int]int? 
    // Key = number we saw
    // Value = index where we saw it
    // We need index in the result, so must store it
    
    for i, n := range arr {
        // i = current index
        // n = current number
        
        comp := tar - n
        // Complement: the number that would sum with n to reach target
        // If n=2 and tar=9, we need 7 (because 2+7=9)
        
        if j, ok := compMap[comp]; ok {
            // Check if we've ALREADY seen the complement
            // j = the index where we saw complement
            // ok = true if complement exists in our map
            
            return []int{j, i}
            // Found it! Return both indices
            // j comes first (we saw it earlier)
        }
        
        compMap[n] = i
        // Haven't found pair yet, store current number
        // "Hey future iterations, we saw n at index i"
    }
    
    return nil  // No pair found
}
```

**Step-by-step trace:**
```
Input: arr = [2, 7, 11, 15], tar = 9

Initial:
compMap = {}

Iteration 1: i=0, n=2
    comp = 9 - 2 = 7
    Is 7 in compMap? No
    Add: compMap = {2: 0}
    
Iteration 2: i=1, n=7
    comp = 9 - 7 = 2
    Is 2 in compMap? Yes! At index 0
    Return: [0, 1]

Result: [0, 1] (arr[0]=2, arr[1]=7, and 2+7=9 ✓)
```

**Why map eliminates nested loop:**
```
Brute force O(n²):
for i := 0; i < len(arr); i++ {
    for j := i+1; j < len(arr); j++ {
        if arr[i] + arr[j] == tar {
            return []int{i, j}
        }
    }
}
// Check every pair: n*(n-1)/2 comparisons

Map approach O(n):
// Single pass: for each element, O(1) lookup in map
// Total: n * O(1) = O(n)
```

**Complexity: O(n) time, O(n) space**
- Time: Single pass, O(1) map lookup per element
- Space: Worst case store all elements in map

### Example 4: First Non-Repeating Character
**From solu/main.go:9**

```go
func nonRepeating(s string) string {
    countMap := make(map[rune]int)
    for _, ch := range s {
        _, ok := countMap[ch]
        if ok {
            countMap[ch]++
        } else {
            countMap[ch] = 1
        }
    }
    for _, ch := range s {
        if countMap[ch] == 1 {
            return string(ch)
        }
    }
    return ""
}
```

**Line-by-line explanation:**
```go
func nonRepeating(s string) string {
    // Purpose: Find first character that appears only once
    // Example: "leetcode" → "l" (appears once, comes first)
    // Example: "loveleetcode" → "v" (l,o,e repeat)
    
    countMap := make(map[rune]int)
    // map[rune]int to count frequency of each character
    
    for _, ch := range s {
        // First pass: count all characters
        
        _, ok := countMap[ch]
        if ok {
            countMap[ch]++
            // Character already in map, increment count
        } else {
            countMap[ch] = 1
            // First time seeing this character
        }
        // Could simplify to: countMap[ch]++
        // But explicit if/else shows the logic clearly
    }
    
    for _, ch := range s {
        // CRITICAL: Second pass through ORIGINAL STRING
        // Why not range over map? Maps are unordered in Go!
        // We need to find FIRST non-repeating character
        // Must preserve original order from string
        
        if countMap[ch] == 1 {
            // Look up count in O(1) time
            return string(ch)  // Found first one, return immediately
        }
    }
    
    return ""  // All characters repeat
}
```

**Step-by-step trace:**
```
Input: s = "leetcode"

Pass 1: Build frequency map
    'l': 1
    'e': 3  (e appears 3 times)
    't': 1
    'c': 1
    'o': 1
    'd': 1

Pass 2: Find first non-repeating (scan original string order)
    Check 'l': countMap['l'] = 1 ✓
    Return: "l"

---

Input: s = "loveleetcode"

Pass 1: Build frequency map
    'l': 2
    'o': 2
    'v': 1
    'e': 4
    't': 1
    'c': 1
    'd': 1

Pass 2: Find first non-repeating
    Check 'l': countMap['l'] = 2 (repeats, skip)
    Check 'o': countMap['o'] = 2 (repeats, skip)
    Check 'v': countMap['v'] = 1 ✓
    Return: "v"
```

**Key pattern: Order matters → loop original, not map**
```go
// WRONG - loses order:
for ch, count := range countMap {
    if count == 1 {
        return string(ch)  // Random character with count=1
    }
}

// CORRECT - preserves order:
for _, ch := range s {
    if countMap[ch] == 1 {
        return string(ch)  // First character with count=1
    }
}
```

**Complexity: O(n) time, O(k) space**
- Time: Two passes through string (both O(n))
- Space: O(k) for map where k = unique characters

### Hash Table Complexity Table
| Operation | Average Time | Worst Time | Reason |
|-----------|-------------|------------|---------|
| Lookup | O(1) | O(n) | Direct index; collisions rare |
| Insert | O(1) | O(n) | Direct index; may resize |
| Delete | O(1) | O(n) | Direct index |
| Space | O(n) | O(n) | Store all key-value pairs |

---

## Linked Lists

### Real World Analogy
A linked list is like a treasure hunt:
- Each clue (node) tells you where the next clue is
- You must follow clues in sequence - can't skip to clue #5
- Adding a clue is easy - just redirect one pointer
- Finding the middle requires counting all clues first (or using a trick)

### Plain English Explanation
Linked lists store elements in nodes scattered in memory. Each node contains:
1. **Data** (the value)
2. **Pointer** (memory address of next node)

Unlike arrays, nodes aren't contiguous. You can't jump to index 5 - you must traverse from the head.

### ASCII Diagram
```
Singly Linked List:

Head → [1|•]→[2|•]→[3|•]→[4|•]→[5|nil]
       ────   ────   ────   ────   ────
       Val=1  Val=2  Val=3  Val=4  Val=5
       Next→  Next→  Next→  Next→  Next=nil

Memory layout (NOT contiguous):
┌─────────────────────────────────────────┐
│  Address  │  Value  │  Next Pointer     │
├─────────────────────────────────────────┤
│   1000    │    1    │    2500           │
│   2500    │    2    │    7200           │
│   7200    │    3    │    4100           │
│   4100    │    4    │    9000           │
│   9000    │    5    │    nil            │
└─────────────────────────────────────────┘

Compare to Array (contiguous):
┌──────┬──────┬──────┬──────┬──────┐
│  1   │  2   │  3   │  4   │  5   │
├──────┼──────┼──────┼──────┼──────┤
│ 1000 │ 1004 │ 1008 │ 1012 │ 1016 │
└──────┴──────┴──────┴──────┴──────┘
```

### Node Structure
**From linkedlist/main.go:7**
```go
type Node struct {
    Val  int    // The data
    Next *Node  // Pointer to next node (or nil if last)
}
```

### Example 1: Reverse Linked List
**From linkedlist/main.go:12**

```go
func revLinkedList(head *Node) *Node {
    current := head
    var prev *Node
    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }
    return prev
}
```

**Line-by-line explanation:**
```go
func revLinkedList(head *Node) *Node {
    // Purpose: Reverse direction of all pointers
    // [1]→[2]→[3]→nil  becomes  nil←[1]←[2]←[3]
    //  ↑                                      ↑
    // old head                             new head
    
    current := head
    // Start at the first node
    // current is our "walker" that will visit each node
    
    var prev *Node
    // prev starts as nil (will become the new tail's next)
    // This tracks the node we just processed
    
    for current != nil {
        // Continue until we've processed all nodes
        // When current=nil, we've fallen off the end
        
        next := current.Next
        // CRITICAL: Save the next node BEFORE we break the link
        // We're about to overwrite current.Next
        // If we don't save it, we lose the rest of the list!
        
        current.Next = prev
        // Reverse the pointer!
        // Point current node back to previous node
        // This is the actual "reversal" operation
        
        prev = current
        // Move prev forward (it was behind, now catches up)
        // prev is now the last node we reversed
        
        current = next
        // Move current forward to the next node to process
        // Uses the next pointer we saved earlier
    }
    
    return prev
    // Why prev and not current?
    // When loop ends, current=nil (fell off end)
    // prev is the last node we processed = new head
}
```

**Step-by-step trace:**
```
Input: [1]→[2]→[3]→nil

Initial state:
current = Node1
prev = nil
List: [1]→[2]→[3]→nil

┌─────────────────────────────────────────────────┐
│ Iteration 1: Reverse Node1                      │
└─────────────────────────────────────────────────┘
Before:
    prev    current   (rest of list)
    nil      [1]   →   [2]→[3]→nil

    next = current.Next  →  next = [2]
    
Step by step:
    next = [2]              (save the rest)
    current.Next = prev     ([1]→nil)
    prev = current          (prev now [1])
    current = next          (current now [2])

After:
          prev    current
    nil←─[1]      [2]→[3]→nil

┌─────────────────────────────────────────────────┐
│ Iteration 2: Reverse Node2                      │
└─────────────────────────────────────────────────┘
Before:
          prev    current
    nil←─[1]      [2]→[3]→nil

    next = [3]              (save the rest)
    current.Next = prev     ([2]→[1])
    prev = current          (prev now [2])
    current = next          (current now [3])

After:
                prev    current
    nil←[1]←─[2]        [3]→nil

┌─────────────────────────────────────────────────┐
│ Iteration 3: Reverse Node3                      │
└─────────────────────────────────────────────────┘
Before:
                prev    current
    nil←[1]←[2]         [3]→nil

    next = nil              (no more nodes)
    current.Next = prev     ([3]→[2])
    prev = current          (prev now [3])
    current = next          (current now nil)

After:
                      prev    current
    nil←[1]←[2]←[3]          nil

┌─────────────────────────────────────────────────┐
│ Loop Exit: current == nil                       │
└─────────────────────────────────────────────────┘
Return prev (the new head)

Final result: nil←[1]←[2]←[3]
                           ↑
                        new head
```

**Complexity: O(n) time, O(1) space**
- Time: Visit each node exactly once
- Space: Only 3 pointers (current, prev, next) regardless of list size

### Example 2: Find Middle Element
**From linkedlist/main.go:39**

```go
func middleLinkedList(head *Node) int {
    current := head
    count := 0
    for current != nil {
        count++
        current = current.Next
    }
    mid := count / 2
    return getAt(head, mid)
}

func getAt(head *Node, index int) int {
    current := head
    for i := 0; i < index; i++ {
        current = current.Next
    }
    return current.Val
}
```

**Line-by-line explanation:**
```go
func middleLinkedList(head *Node) int {
    // Purpose: Find the middle element value
    // Method: Count length, then traverse to middle
    
    current := head
    count := 0
    
    for current != nil {
        // First pass: count total nodes
        count++             // Increment counter
        current = current.Next  // Move to next node
    }
    // After loop: count = total length
    
    mid := count / 2
    // Integer division: 5/2 = 2 (0-indexed: positions 0,1,2,3,4)
    // For even length, this gives second middle element
    
    return getAt(head, mid)
    // Helper function to traverse to specific index
}

func getAt(head *Node, index int) int {
    // Purpose: Get value at specific index (0-based)
    // Similar to arr[index] but for linked list
    
    current := head
    // Start at beginning
    
    for i := 0; i < index; i++ {
        // Traverse index times
        // To get index=2: move 0→1, 1→2 (2 moves)
        current = current.Next
    }
    
    return current.Val
    // Return value at target position
}
```

**Step-by-step trace:**
```
Input: [1]→[2]→[3]→[4]→[5]→nil

Pass 1: Count nodes
    Start: current=[1], count=0
    Step 1: count=1, current=[2]
    Step 2: count=2, current=[3]
    Step 3: count=3, current=[4]
    Step 4: count=4, current=[5]
    Step 5: count=5, current=nil (exit loop)
    
    Result: count=5

Calculate middle:
    mid = 5 / 2 = 2 (index 2 is 3rd element)

Pass 2: Get element at index 2
    getAt(head, 2):
        Start: current=[1], i=0
        i=0: current=[2] (moved once)
        i=1: current=[3] (moved twice)
        Loop exits (i=2, not < 2)
        
        Return: current.Val = 3

Return: 3
```

**Complexity: O(n) time, O(1) space**
- Time: O(n) to count + O(n/2) to traverse = O(n)
- Space: Only a few variables

**Alternative: Fast/Slow Pointer (O(n) time, one pass)**
```go
func middleFastSlow(head *Node) int {
    slow := head
    fast := head
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next       // Move 1 step
        fast = fast.Next.Next  // Move 2 steps
    }
    
    return slow.Val  // When fast reaches end, slow is at middle
}
```

### Example 3: Cycle Detection
**From linkedlist/main.go:49**

```go
func hasCycle(head *Node) bool {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

**Line-by-line explanation:**
```go
func hasCycle(head *Node) bool {
    // Purpose: Detect if linked list has a cycle
    // Cycle: a node's Next points back to earlier node
    // Uses Floyd's "Tortoise and Hare" algorithm
    
    slow := head
    fast := head
    // Both pointers start at same position
    // slow moves 1 step per iteration
    // fast moves 2 steps per iteration
    
    for fast != nil && fast.Next != nil {
        // Must check TWO conditions:
        // 1. fast != nil (fast is valid)
        // 2. fast.Next != nil (can take 2 steps)
        // If either is nil, no cycle (reached end)
        
        slow = slow.Next
        // Tortoise: move 1 step
        
        fast = fast.Next.Next
        // Hare: move 2 steps
        // This is why we need fast.Next != nil check
        
        if slow == fast {
            // Pointers met! Must be a cycle
            // Why? In a cycle, fast will eventually "lap" slow
            // Like runners on a track - faster runner catches up
            return true
        }
    }
    
    return false
    // Reached end of list (fast hit nil)
    // No cycle exists
}
```

**Visual trace with cycle:**
```
List with cycle:
[1]→[2]→[3]→[4]→[5]
         ↑         ↓
         └─────────┘

Initial:
slow = [1]
fast = [1]

Iteration 1:
    slow moves 1: slow = [2]
    fast moves 2: fast = [3]
    slow == fast? No

Iteration 2:
    slow moves 1: slow = [3]
    fast moves 2: fast = [5] (from [3]→[4]→[5])
    slow == fast? No

Iteration 3:
    slow moves 1: slow = [4]
    fast moves 2: fast = [3] (from [5]→[2]→[3])
    slow == fast? No

Iteration 4:
    slow moves 1: slow = [5]
    fast moves 2: fast = [5] (from [3]→[4]→[5])
    slow == fast? YES!
    Return: true
```

**Visual trace without cycle:**
```
List without cycle:
[1]→[2]→[3]→[4]→[5]→nil

Initial:
slow = [1]
fast = [1]

Iteration 1:
    slow = [2]
    fast = [3]

Iteration 2:
    slow = [3]
    fast = [5]

Iteration 3:
    slow = [4]
    fast = nil (from [5].Next.Next)
    Loop condition false (fast.Next == nil)
    Exit loop

Return: false
```

**Complexity: O(n) time, O(1) space**
- Time: In cycle, fast catches slow in at most n iterations
- Space: Only 2 pointers

### Example 4: Delete Node by Value
**From linkedlist/main.go:61**

```go
func delNode(head *Node, val int) *Node {
    var prev *Node
    for head != nil && head.Val == val {
        head = head.Next
    }
    current := head
    
    for current != nil {
        if current.Val == val {
            prev.Next = current.Next
        } else {
            prev = current
        }
        current = current.Next
    }
    return head
}
```

**Line-by-line explanation:**
```go
func delNode(head *Node, val int) *Node {
    // Purpose: Delete all nodes with specific value
    // Example: delete 2 from [2]→[1]→[2]→[3] → [1]→[3]
    
    var prev *Node
    // Track previous node for linking around deleted nodes
    
    for head != nil && head.Val == val {
        // CRITICAL: Handle deletion at head
        // Keep moving head forward while it matches val
        // [2]→[2]→[1]→[3] with val=2
        // First loop: head = [2]→[1]→[3]
        // Second loop: head = [1]→[3]
        head = head.Next
    }
    // After this loop, head is either:
    // 1. nil (all nodes matched val)
    // 2. First node that doesn't match val
    
    current := head
    // Start traversal from new head
    
    for current != nil {
        // Process remaining nodes
        
        if current.Val == val {
            // Found node to delete
            prev.Next = current.Next
            // Bridge over current node:
            // prev → current → next
            // becomes:
            // prev → next (skipping current)
            
            // Don't update prev! It stays the same
            // because current is deleted
            
        } else {
            // Keep this node
            prev = current
            // Update prev to current node
            // Ready to potentially delete next node
        }
        
        current = current.Next
        // Move to next node regardless of deletion
    }
    
    return head
    // Return new head (may have changed if original head deleted)
}
```

**Step-by-step trace:**
```
Input: [1]→[2]→[3]→[2]→[4]→nil, val=2

Step 1: Handle head deletions
    head.Val = 1, not equal to 2
    Skip this loop
    head = [1]→[2]→[3]→[2]→[4]

Step 2: Initialize
    current = [1]
    prev = nil

Step 3: Traverse and delete

Iteration 1: current=[1], val=1
    1 != 2 (keep)
    prev = [1]
    current = [2]
    List: [1]→[2]→[3]→[2]→[4]
           ↑    ↑
          prev curr

Iteration 2: current=[2], val=2
    2 == 2 (delete!)
    prev.Next = current.Next
    [1].Next = [3]
    prev stays [1] (don't advance on deletion)
    current = [3]
    List: [1]→[3]→[2]→[4]
           ↑    ↑
          prev curr

Iteration 3: current=[3], val=3
    3 != 2 (keep)
    prev = [3]
    current = [2]
    List: [1]→[3]→[2]→[4]
                ↑    ↑
               prev curr

Iteration 4: current=[2], val=2
    2 == 2 (delete!)
    prev.Next = current.Next
    [3].Next = [4]
    prev stays [3]
    current = [4]
    List: [1]→[3]→[4]
                ↑    ↑
               prev curr

Iteration 5: current=[4], val=4
    4 != 2 (keep)
    prev = [4]
    current = nil
    Exit loop

Return: [1]→[3]→[4]→nil
```

**Edge case: Delete all nodes**
```
Input: [2]→[2]→[2]→nil, val=2

Step 1: Handle head deletions
    Loop 1: head = [2]→[2], val=2 → head = [2]→[2]
    Loop 2: head = [2], val=2 → head = [2]
    Loop 3: head = nil, val=2 → exit (head == nil)

Step 2: current = nil, skip second loop

Return: nil
```

**Complexity: O(n) time, O(1) space**
- Time: Single pass through list
- Space: Only prev and current pointers

### Linked List Complexity Table
| Operation | Time Complexity | Reason |
|-----------|----------------|---------|
| Access by index | O(n) | Must traverse from head |
| Search | O(n) | Must traverse sequentially |
| Insert at head | O(1) | Just update head pointer |
| Insert at tail | O(n) | Must traverse to end* |
| Insert in middle | O(n) | Must traverse to position |
| Delete at head | O(1) | Just update head pointer |
| Delete at tail | O(n) | Must find second-to-last |
| Delete in middle | O(n) | Must traverse to position |
| Reverse | O(n) | Visit each node once |
| Find middle | O(n) | Must traverse (once or twice) |
| Cycle detection | O(n) | Fast/slow pointer technique |

*With tail pointer maintained: O(1)

---

## Key Patterns Summary

### 1. Two Pointer Pattern
**When to use:** Array/string reversal, palindrome check, finding pairs

**Structure:**
```go
left := 0
right := len(arr) - 1
for left < right {
    // Process arr[left] and arr[right]
    left++
    right--
}
```

**Examples from code:**
- Array reversal (arrays/main.go:5)
- Valid palindrome (solu/main.go:76)
- Reverse vowels (solu/main.go:93)

**Complexity:** O(n) time, O(1) space

---

### 2. Fast/Slow Pointer Pattern
**When to use:** Linked list cycle detection, finding middle element

**Structure:**
```go
slow := head
fast := head
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
    if slow == fast {
        // Cycle detected or middle found
    }
}
```

**Examples from code:**
- Cycle detection (linkedlist/main.go:49)

**Complexity:** O(n) time, O(1) space

---

### 3. Prev/Current Pattern for Linked Lists
**When to use:** Linked list reversal, node deletion

**Structure:**
```go
var prev *Node
current := head
for current != nil {
    next := current.Next  // Save next before breaking link
    current.Next = prev   // Reverse pointer
    prev = current        // Move prev forward
    current = next        // Move current forward
}
```

**Examples from code:**
- Reverse linked list (linkedlist/main.go:12)
- Delete node (linkedlist/main.go:61)

**Complexity:** O(n) time, O(1) space

---

### 4. Map to Eliminate Nested Loops
**When to use:** Finding pairs, detecting duplicates, complement search

**Bad (O(n²)):**
```go
for i := 0; i < len(arr); i++ {
    for j := i+1; j < len(arr); j++ {
        if arr[i] + arr[j] == target {
            return []int{i, j}
        }
    }
}
```

**Good (O(n)):**
```go
seen := make(map[int]int)
for i, n := range arr {
    if j, ok := seen[target-n]; ok {
        return []int{j, i}
    }
    seen[n] = i
}
```

**Examples from code:**
- Two sum (solu/main.go:26)
- Duplicate detection (hashmaps/main.go:5)

**Complexity:** O(n²) → O(n) with map

---

### 5. Count with map[T]int vs Existence with map[T]bool

**Use map[T]bool when:**
- Checking membership/existence
- Detecting duplicates
- Set operations

```go
seen := make(map[int]bool)
for _, n := range arr {
    if seen[n] {
        return true  // Duplicate found
    }
    seen[n] = true
}
```

**Use map[T]int when:**
- Counting frequencies
- Finding most common element
- Anagram detection

```go
freq := make(map[rune]int)
for _, ch := range s {
    freq[ch]++  // Auto-initializes to 0, then increments
}
```

**Examples from code:**
- map[int]bool: isSeenMoreThanOnce (hashmaps/main.go:5)
- map[rune]int: isMoreOften (hashmaps/main.go:15)
- map[rune]bool: countVowel (strings/main.go:7)
- map[rune]int: isAnan (strings/main.go:23)

---

### 6. Never Loop Through Map to Search
**Bad (O(n) when you already have O(1)):**
```go
for key, val := range myMap {
    if key == target {
        return val
    }
}
```

**Good (O(1)):**
```go
if val, ok := myMap[target]; ok {
    return val
}
```

**Why:** Maps provide O(1) direct access by key. Looping defeats the purpose!

---

### 7. When Order Matters, Loop Original String Not Map
**Critical for:** First non-repeating character, preserving sequence

**Wrong (loses order):**
```go
freq := make(map[rune]int)
for _, ch := range s {
    freq[ch]++
}
// Maps are UNORDERED in Go!
for ch, count := range freq {
    if count == 1 {
        return ch  // Random character with count=1
    }
}
```

**Correct (preserves order):**
```go
freq := make(map[rune]int)
for _, ch := range s {
    freq[ch]++
}
// Loop ORIGINAL string to preserve order
for _, ch := range s {
    if freq[ch] == 1 {
        return ch  // FIRST character with count=1
    }
}
```

**Examples from code:**
- nonRepeating (solu/main.go:9)

---

### 8. String Immutability Pattern
**Remember:** Strings are immutable in Go

**Pattern:**
```go
s := "hello"
b := []byte(s)      // Convert to mutable byte slice
// Modify b
b[0], b[4] = b[4], b[0]
result := string(b) // Convert back to string
```

**Examples from code:**
- Reverse vowels (solu/main.go:93)
- Valid palindrome (solu/main.go:76)

---

### Quick Reference Table

| Problem Type | Pattern | Time | Space |
|-------------|---------|------|-------|
| Array reversal | Two pointer | O(n) | O(1) |
| Palindrome check | Two pointer | O(n) | O(n)* |
| Find pairs sum | Map complement | O(n) | O(n) |
| Detect duplicate | Map existence | O(n) | O(n) |
| Character frequency | Map count | O(n) | O(k) |
| First unique char | Map + original loop | O(n) | O(k) |
| Reverse linked list | Prev/current | O(n) | O(1) |
| Cycle detection | Fast/slow pointer | O(n) | O(1) |
| Find middle | Fast/slow pointer | O(n) | O(1) |
| Delete node | Prev/current | O(n) | O(1) |

*O(n) space for byte slice conversion in Go

---

## Summary

### Big O Rules of Thumb
1. **Single loop** → O(n)
2. **Nested loops** → O(n²)
3. **Map/hash lookup** → O(1) average
4. **Divide in half repeatedly** → O(log n)
5. **Constant operations** → O(1)

### Data Structure Selection
- **Need index access?** → Array/Slice
- **Frequent insert/delete at beginning?** → Linked List
- **Need fast lookup by key?** → Hash Map
- **String modification?** → Convert to []byte first

### Common Optimizations
- Two loops of O(n) is still O(n), not O(n²)
- Map lookup is O(1), eliminates nested loop
- Two pointers often gives O(1) space instead of O(n)
- Fast/slow pointers solve problems in one pass

### Go-Specific Notes
- `range` on string gives **runes** (Unicode), not bytes
- Maps are **unordered** - preserve sequence by looping original
- Strings are **immutable** - use []byte for modification
- Zero values: int→0, bool→false, pointer→nil
- `make(map[T]V)` auto-initializes on first write

---

*Generated from actual code in: bigo/main.go, arrays/main.go, strings/main.go, hashmaps/main.go, linkedlist/main.go, solu/main.go*

---

## Lesson 6: Stacks

### Real World Analogy
A stack is like a stack of plates in a cafeteria:
- You add plates to the top
- You remove plates from the top
- You can't pull a plate from the middle
- Last plate added is the first plate removed (LIFO)

### Plain English Explanation
A stack is a Last-In-First-Out (LIFO) data structure. Think of it as a vertical collection where you can only add or remove from one end (the top). The most recently added element is the first one to be removed.

### ASCII Diagram
```
    TOP
     ↓
   ┌───┐
   │ 5 │  ← Most recently added (top)
   ├───┤
   │ 4 │
   ├───┤
   │ 3 │
   ├───┤
   │ 2 │
   ├───┤
   │ 1 │  ← First added (bottom)
   └───┘

Operations:
Push(6):  Add to top
Pop():    Remove from top (returns 5)
Peek():   Look at top without removing
```

### Go Implementation

Go doesn't have a built-in stack type, but slices work perfectly:

```go
// Create stack
stack := []int{}

// Push (add to top)
stack = append(stack, value)

// Peek (look at top)
top := stack[len(stack)-1]

// Pop (remove from top)
stack = stack[:len(stack)-1]

// isEmpty
isEmpty := len(stack) == 0
```

### Example 1: Balanced Parentheses
**From stacks/main.go:10**

```go
func balancedBracket(s string) bool {
    stack := []rune{}
    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }
    for _, ch := range s {
        if ch == '(' || ch == '{' || ch == '[' {
            stack = append(stack, ch)
        } else {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if top != pairs[ch] {
                return false
            }
        }
    }
    return len(stack) == 0
}
```

**Step-by-step trace:**
```
Input: s = "{[()]}"

Iteration 1: ch = '{'
  Opening bracket? Yes
  Push: stack = ['{']

Iteration 2: ch = '['
  Opening bracket? Yes
  Push: stack = ['{', '[']

Iteration 3: ch = '('
  Opening bracket? Yes
  Push: stack = ['{', '[', '(']

Iteration 4: ch = ')'
  Closing bracket? Yes
  Stack empty? No
  top = '('
  Pop: stack = ['{', '[']
  Match? '(' == pairs[')'] = '(' ✓

Iteration 5: ch = ']'
  Closing bracket? Yes
  top = '['
  Pop: stack = ['{']
  Match? '[' == pairs[']'] = '[' ✓

Iteration 6: ch = '}'
  Closing bracket? Yes
  top = '{'
  Pop: stack = []
  Match? '{' == pairs['}'] = '{' ✓

Stack empty? Yes → return true
```

**Complexity: O(n) time, O(n) space**

### Example 2: Min Stack
**From stacks/main.go:5**

The challenge: Design a stack that supports getMin() in O(1) time.

**Solution: Two stacks**
- main stack: stores all values
- min stack: stores minimum value at each stack height

```go
type Stack struct {
    main []int
    min  []int
}

func (s *Stack) Push(val int) {
    s.main = append(s.main, val)
    if len(s.min) == 0 || val <= s.min[len(s.min)-1] {
        s.min = append(s.min, val)
    } else {
        s.min = append(s.min, s.min[len(s.min)-1])
    }
}

func (s *Stack) getMin() int {
    return s.min[len(s.min)-1]
}
```

**Step-by-step trace:**
```
Push(3):
  main = [3]
  min empty? Yes → min = [3]

Push(5):
  main = [3, 5]
  5 <= 3? No → min = [3, 3] (keep current min)

Push(2):
  main = [3, 5, 2]
  2 <= 3? Yes → min = [3, 3, 2] (new min)

Push(1):
  main = [3, 5, 2, 1]
  1 <= 2? Yes → min = [3, 3, 2, 1] (new min)

getMin() → min[top] = 1  (O(1) time!)

Pop():
  main = [3, 5, 2]
  min = [3, 3, 2]

getMin() → min[top] = 2  (O(1) time!)
```

**Why this works:** The min stack tracks the minimum at EACH level. When we pop, the previous minimum is automatically revealed.

### Complexity Analysis

| Operation | Time | Reason |
|-----------|------|--------|
| Push | O(1) | Append to slice |
| Pop | O(1) | Slice operation |
| Peek | O(1) | Array access |
| isEmpty | O(1) | Length check |
| getMin (with min stack) | O(1) | Array access |

### When to Use a Stack

✓ **Balanced parentheses/brackets**
✓ **Undo/redo functionality**
✓ **Function call stack (recursion)**
✓ **Reverse a sequence**
✓ **Track min/max in O(1)**
✓ **Expression evaluation (postfix)**
✓ **Backtracking algorithms**

---

## Lesson 7: Queues

### Real World Analogy
A queue is like a line at a coffee shop:
- People join at the back of the line
- People are served from the front of the line
- First person in line is first to be served (FIFO)
- You can't cut in the middle!

### Plain English Explanation
A queue is a First-In-First-Out (FIFO) data structure. Elements are added at the back (rear) and removed from the front. The first element added is the first one to be removed.

### ASCII Diagram
```
  FRONT                       BACK
    ↓                          ↓
  ┌───┬───┬───┬───┬───┐
  │ 1 │ 2 │ 3 │ 4 │ 5 │
  └───┴───┴───┴───┴───┘
   ↑                   ↑
Dequeue             Enqueue
(remove)             (add)
```

### Go Implementation

```go
// Create queue
queue := []int{}

// Enqueue (add to back)
queue = append(queue, value)  // O(1)

// Peek at front
front := queue[0]  // O(1)

// Dequeue (remove from front)
queue = queue[1:]  // O(n) - this is expensive!
```

### Why Dequeue is O(n) Not O(1)

When we do `queue = queue[1:]`, Go must:
1. Create a new slice header pointing to index 1
2. Eventually garbage collect the old elements
3. If underlying array gets too sparse, reallocate

**For true O(1) dequeue, use:**
- Circular buffer
- Linked list
- Two stacks (interview trick)

### Example: Process Queue
**From queues/main.go:7**

```go
func processQueue(tasks []string) {
    queue := []string{}
    for _, t := range tasks {
        queue = append(queue, t)
    }
    for len(queue) > 0 {
        front := queue[0]
        queue = queue[1:]
        fmt.Println("processing", front)
    }
}
```

**Step-by-step trace:**
```
Input: tasks = ["task1", "task2", "task3"]

Enqueue all:
  queue = ["task1", "task2", "task3"]

Dequeue and process:
  Dequeue: front = "task1", queue = ["task2", "task3"]
  Process: "task1"

  Dequeue: front = "task2", queue = ["task3"]
  Process: "task2"

  Dequeue: front = "task3", queue = []
  Process: "task3"

Result: Tasks processed in order received
```

### Stack vs Queue Comparison

| Aspect | Stack (LIFO) | Queue (FIFO) |
|--------|-------------|--------------|
| Order | Last In First Out | First In First Out |
| Add | Push to top | Enqueue to back |
| Remove | Pop from top | Dequeue from front |
| Analogy | Stack of plates | Line of people |
| Use cases | Undo, recursion, reversal | Task processing, BFS, scheduling |

### When to Use a Queue

✓ **Process tasks in order received**
✓ **Breadth-first search (BFS)**
✓ **Print queue, request handling**
✓ **Level-order tree traversal**
✓ **Buffer for streaming data**

### Complexity Analysis

| Operation | Time | Reason |
|-----------|------|--------|
| Enqueue | O(1) | Append to slice |
| Dequeue | O(n) | Slice reslicing* |
| Front/Peek | O(1) | Array access |
| isEmpty | O(1) | Length check |

*With proper circular buffer or linked list: O(1)

---

## Lesson 8: Binary Search

### Real World Analogy
Finding a word in a dictionary:
- You don't start at page 1 and read every page!
- You open to the middle
- If your word comes before the middle word, search the left half
- If your word comes after, search the right half
- Repeat until you find it

### Plain English Explanation
Binary search is an efficient algorithm for finding an element in a **sorted** array. It works by repeatedly dividing the search space in half, eliminating half the remaining elements with each comparison.

**CRITICAL: Array MUST be sorted!**

### ASCII Diagram
```
Array: [1, 3, 5, 7, 9, 11, 13, 15, 17], target = 7

Step 1: left=0, right=8, mid=4
  [1, 3, 5, 7, 9, 11, 13, 15, 17]
             ↑
          arr[4]=9 > 7 → search left

Step 2: left=0, right=3, mid=1
  [1, 3, 5, 7] (right half eliminated)
      ↑
   arr[1]=3 < 7 → search right

Step 3: left=2, right=3, mid=2
  [5, 7] (left half eliminated)
   ↑
 arr[2]=5 < 7 → search right

Step 4: left=3, right=3, mid=3
  [7]
   ↑
 arr[3]=7 == 7 → FOUND!
```

### Example: Binary Search
**From binary_search/main.go:5**

```go
func binarySearch(arr []int, tar int) int {
    left := 0
    right := len(arr) - 1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] == tar {
            return mid
        } else if arr[mid] < tar {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return left
}
```

### Why mid = left + (right-left)/2 NOT (left+right)/2?

**Two reasons:**

1. **Prevents integer overflow**
   ```
   If left and right are large (near max int):
   left + right could overflow!
   
   Example: left = 2^30, right = 2^30
   left + right = 2^31 (overflow!)
   
   But left + (right-left)/2 never overflows
   ```

2. **Mathematical equivalence**
   ```
   left + (right-left)/2
   = left + (right/2 - left/2)
   = left/2 + right/2
   = (left + right)/2
   
   Same result, no overflow risk!
   ```

### When left > right Means Not Found

When the loop exits with `left > right`, the search space is empty:
- We've exhausted all possibilities
- left has "crossed over" right
- The target doesn't exist in the array

### Step-by-step Trace
```
Input: arr = [1, 2, 3, 4], tar = 10

Initial: left=0, right=3

Iteration 1:
  mid = 0 + (3-0)/2 = 1
  arr[1] = 2
  2 < 10 → search right
  left = 2

Iteration 2:
  mid = 2 + (3-2)/2 = 2
  arr[2] = 3
  3 < 10 → search right
  left = 3

Iteration 3:
  mid = 3 + (3-3)/2 = 3
  arr[3] = 4
  4 < 10 → search right
  left = 4

Iteration 4:
  left=4, right=3
  left <= right? NO (4 <= 3 is false)
  Exit loop

Return left = 4 (insertion position)
```

### Search Insert Position Pattern

When binary search doesn't find the target, `left` points to where the element should be inserted to maintain sorted order.

```
Example: arr = [1, 3, 5, 7], tar = 4

Final state: left = 2 (between 3 and 5)
Insert at index 2: [1, 3, 4, 5, 7] ✓
```

### Searching Number Ranges (Not Just Arrays)

Binary search works on any sorted "search space", not just arrays!

**Example: Find square root**
```go
func sqrt(n int) int {
    left, right := 0, n
    for left <= right {
        mid := left + (right-left)/2
        square := mid * mid
        if square == n {
            return mid
        } else if square < n {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return right  // Floor of square root
}
```

You're searching the "imaginary array" [0, 1, 2, ..., n]!

### Complexity Analysis

**Time: O(log n)** - Why?
```
Iteration 1: n elements
Iteration 2: n/2 elements
Iteration 3: n/4 elements
Iteration k: n/(2^k) elements

Stop when n/(2^k) = 1 → k = log₂(n)

Examples:
- 1,000 elements: ~10 iterations
- 1,000,000 elements: ~20 iterations
```

**Space: O(1)** - Only use left, right, mid variables

### When to Use Binary Search

✓ Array is sorted
✓ Need faster than O(n) search
✓ Finding insertion position
✓ Finding first/last occurrence
✓ Searching in ranges (square root, capacity problems)

---

## Lesson 9: Two Pointers

### Real World Analogy
Two people searching for a specific sum in a sorted list of prices:
- One starts at the cheapest item (left)
- One starts at the most expensive item (right)
- If sum too small, left person moves to next expensive item
- If sum too large, right person moves to next cheaper item
- They meet in the middle having checked all possible pairs in one pass!

### Plain English Explanation
Two pointers is a technique using two indices that traverse data in tandem. It eliminates nested loops, reducing time complexity from O(n²) to O(n).

### Three Main Patterns

#### 1. Opposite Ends (Convergence)
```
Start: left=0, right=n-1
Move toward each other

┌───┬───┬───┬───┬───┐
│ 1 │ 2 │ 3 │ 4 │ 5 │
└───┴───┴───┴───┴───┘
  ↑               ↑
 left           right

Use for: palindromes, two sum (sorted), reversal
```

#### 2. Fast/Slow (Different Speeds)
```
Start: both at 0
Slow moves 1 step, fast moves 2 steps

┌───┬───┬───┬───┬───┐
│ 1 │ 2 │ 3 │ 4 │ 5 │
└───┴───┴───┴───┴───┘
  ↑   ↑
slow fast

Use for: cycle detection, finding middle
```

#### 3. Same Direction (Sliding/Expanding)
```
Start: both at 0
Move in same direction at different rates

┌───┬───┬───┬───┬───┐
│ 1 │ 1 │ 2 │ 2 │ 3 │
└───┴───┴───┴───┴───┘
  ↑   ↑
  i   j

Use for: remove duplicates, partitioning
```

### Example 1: Two Sum (Sorted Array)
**From two_pointer/main.go:9**

```go
func twoSum(arr []int, tar int) bool {
    left := 0
    right := len(arr) - 1
    for left <= right {
        sum := arr[left] + arr[right]
        if sum == tar {
            return true
        } else if sum > tar {
            right--
        } else {
            left++
        }
    }
    return false
}
```

**Step-by-step trace:**
```
Input: arr = [1, 2, 3, 4, 5, 6], tar = 9

Iteration 1:
  left=0, right=5
  sum = 1 + 6 = 7
  7 < 9 → sum too small → left++

Iteration 2:
  left=1, right=5
  sum = 2 + 6 = 8
  8 < 9 → sum too small → left++

Iteration 3:
  left=2, right=5
  sum = 3 + 6 = 9
  9 == 9 → FOUND!
  Return true

Result: true (3 + 6 = 9)
```

**When to move left vs right:**
- sum < target: need larger sum → left++ (increase smaller number)
- sum > target: need smaller sum → right-- (decrease larger number)

This works because array is **SORTED**!

### Example 2: Palindrome Check
**From two_pointer/main.go:24**

```go
func palindromeTwoSum(s string) bool {
    re := regexp.MustCompile("[^a-zA-Z0-9]+")
    clean := re.ReplaceAllString(s, "")
    s = strings.ToLower(strings.TrimSpace(clean))
    left := 0
    right := len(s) - 1
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}
```

**Step-by-step trace:**
```
Input: s = "A man, a plan, a canal: Panama"

Step 1: Clean
  Remove non-alphanumeric: "AmanaplanacanalPanama"
  Lowercase: "amanaplanacanalpanama"

Step 2: Two pointer check
  left=0, right=20

  s[0]='a', s[20]='a' → match ✓
  left=1, right=19

  s[1]='m', s[19]='m' → match ✓
  left=2, right=18

  ... (all match)

  left=10, right=10
  left < right? No
  Exit loop

Result: true (is palindrome)
```

### Why Two Pointers Beats Nested Loops

**Brute force (O(n²)):**
```go
for i := 0; i < n; i++ {
    for j := i+1; j < n; j++ {
        // Check pairs: n*(n-1)/2 comparisons
    }
}
```

**Two pointers (O(n)):**
```go
left, right := 0, n-1
for left < right {
    // Each element visited once: n comparisons
}
```

For n=1000: **500,000 operations → 1000 operations (500x faster!)**

### Complexity Analysis

| Operation | Time | Space | Reason |
|-----------|------|-------|--------|
| Two sum (sorted) | O(n) | O(1) | Single pass, two variables |
| Palindrome | O(n) | O(n) | Single pass, cleaned string |
| Reverse array | O(n) | O(1) | Single pass, in-place |

### When to Use Two Pointers

✓ Array/string is sorted (or can be sorted)
✓ Looking for pairs/triplets
✓ Need to eliminate nested loops
✓ Palindrome checking
✓ Removing duplicates in-place

---

## Lesson 10: Sliding Window

### Real World Analogy
Looking through a moving window on a train:
- The window shows a fixed section of the landscape
- As the train moves, new scenery enters the window on one side
- Old scenery leaves the window on the other side
- You don't need to look at the entire landscape again - just track what enters/leaves!

### Plain English Explanation
Sliding window is a technique to efficiently process subarrays or substrings by:
1. Maintaining a "window" of elements
2. Sliding the window by adding new element and removing old element
3. Avoiding recalculation from scratch each time

### Two Types of Sliding Window

#### 1. Fixed Window
Window size is constant. Slide by adding arr[right] and removing arr[right-k].

```
Array: [1, 2, 3, 4, 5], k=3

Window 1: [1, 2, 3] sum=6
           -------
Window 2:    [2, 3, 4] sum=9 (added 4, removed 1)
              -------
Window 3:       [3, 4, 5] sum=12 (added 5, removed 2)
                 -------
```

#### 2. Variable Window
Window size changes. Expand with right++, shrink with left++.

```
String: "abcabcbb", find longest substring without repeating

Window 1: "abc" (length 3) ✓
          ---
Window 2: "bca" (length 3) ← shrunk left to remove duplicate 'a'
           ---
Window 3: "cab" (length 3) ← shrunk left to remove duplicate 'b'
            ---
```

### Example 1: Fixed Window - Maximum Sum
**From sliding_window/main.go:5**

```go
func slidingWindow(arr []int, k int) int {
    sum := 0
    for i := 0; i < k; i++ {
        sum += arr[i]
    }
    max := sum
    for i := k; i < len(arr); i++ {
        sum += arr[i]
        sum -= arr[i-k]
        if sum > max {
            max = sum
        }
    }
    return max
}
```

### Why arr[i-k] is the Leaving Element

When we're at index i with window size k:
```
Window contains: arr[i-k+1], arr[i-k+2], ..., arr[i-1], arr[i]

Example: i=5, k=3
  Window = [arr[3], arr[4], arr[5]]
  
When we move to i=6:
  Add arr[6]
  Remove arr[6-3] = arr[3] ← the leftmost element
```

**Step-by-step trace:**
```
Input: arr = [1, 2, 3, 4, 5], k = 3

Step 1: Build first window
  Window: [1, 2, 3]
  sum = 6
  max = 6

Step 2: Slide to i=3
  Add arr[3]=4: sum = 6 + 4 = 10
  Remove arr[0]=1: sum = 10 - 1 = 9
  Window: [2, 3, 4]
  max = max(6, 9) = 9

Step 3: Slide to i=4
  Add arr[4]=5: sum = 9 + 5 = 14
  Remove arr[1]=2: sum = 14 - 2 = 12
  Window: [3, 4, 5]
  max = max(9, 12) = 12

Result: 12
```

### Example 2: Variable Window - Longest Substring
**From sliding_window/main.go:21**

```go
func longestSubString(s string) int {
    seen := make(map[byte]int)
    left := 0
    max := 0
    for right := 0; right < len(s); right++ {
        ch := s[right]
        if idx, ok := seen[ch]; ok && idx >= left {
            left = idx + 1
        }
        seen[ch] = right
        if right-left+1 > max {
            max = right - left + 1
        }
    }
    return max
}
```

### Key Concepts

**1. Seen map stores LAST INDEX, not just existence**
```
We need to know WHERE we saw the character
to update left pointer correctly
```

**2. idx >= left check**
```
Character might be in map but OUTSIDE current window!

Example: s = "abba"
When we see second 'a':
  seen['a'] = 0, but left = 2
  0 < 2, so 'a' is outside current window
  Don't move left backwards!
```

**3. Window size formula: right - left + 1**
```
Example: left=2, right=5
Indices: 2, 3, 4, 5
Count: 5 - 2 + 1 = 4 elements ✓
```

**Step-by-step trace:**
```
Input: s = "abcabcbb"

right=0, ch='a':
  'a' in seen? No
  seen = {'a':0}
  Window: "a" (length 1)
  max = 1

right=1, ch='b':
  'b' in seen? No
  seen = {'a':0, 'b':1}
  Window: "ab" (length 2)
  max = 2

right=2, ch='c':
  'c' in seen? No
  seen = {'a':0, 'b':1, 'c':2}
  Window: "abc" (length 3)
  max = 3

right=3, ch='a':
  'a' in seen AND idx>=left? Yes (0>=0)
  DUPLICATE! left = 0 + 1 = 1
  seen['a'] = 3
  Window: "bca" (length 3)
  max = 3

right=4, ch='b':
  'b' in seen AND idx>=left? Yes (1>=1)
  DUPLICATE! left = 1 + 1 = 2
  seen['b'] = 4
  Window: "cab" (length 3)
  max = 3

right=5, ch='c':
  'c' in seen AND idx>=left? Yes (2>=2)
  DUPLICATE! left = 2 + 1 = 3
  seen['c'] = 5
  Window: "abc" (length 3)
  max = 3

right=6, ch='b':
  'b' in seen AND idx>=left? Yes (4>=3)
  DUPLICATE! left = 4 + 1 = 5
  seen['b'] = 6
  Window: "cb" (length 2)
  max = 3

right=7, ch='b':
  'b' in seen AND idx>=left? Yes (6>=5)
  DUPLICATE! left = 6 + 1 = 7
  seen['b'] = 7
  Window: "b" (length 1)
  max = 3

Result: 3
```

### Complexity Analysis

| Type | Time | Space | Reason |
|------|------|-------|--------|
| Fixed window | O(n) | O(1) | Build first window + slide |
| Variable window | O(n) | O(k) | Each char visited at most twice, k=unique chars |

### When to Use Sliding Window

**Fixed Window:**
✓ Maximum/minimum sum of k elements
✓ Average of k elements
✓ First negative in every window of size k

**Variable Window:**
✓ Longest substring without repeating characters
✓ Minimum window substring
✓ Longest substring with at most k distinct characters
✓ Max consecutive ones after flipping k zeros

### Why Sliding Window?

**Brute force (O(n*k) or O(n²)):**
```go
for i := 0; i <= n-k; i++ {
    sum := 0
    for j := i; j < i+k; j++ {
        sum += arr[j]  // Recalculate sum each time!
    }
}
```

**Sliding window (O(n)):**
```go
sum := sum of first k elements
for i := k; i < n; i++ {
    sum += arr[i]      // Add new element
    sum -= arr[i-k]    // Remove old element
}
```

For n=1000, k=100: **100,000 operations → 1000 operations (100x faster!)**

---

## Updated Summary

### Lessons 6-10 Quick Reference

| Data Structure/Algorithm | Key Pattern | Time | Space | When to Use |
|-------------------------|-------------|------|-------|-------------|
| Stack (balanced brackets) | LIFO with map of pairs | O(n) | O(n) | Matching pairs, undo/redo |
| Stack (min stack) | Two stacks track min | O(1) | O(n) | Track min in constant time |
| Queue (FIFO) | Enqueue/dequeue | O(n)* | O(n) | Process in order, BFS |
| Binary Search | Divide and conquer | O(log n) | O(1) | Sorted array, find in range |
| Two Pointers (opposite) | Converge to middle | O(n) | O(1) | Sorted pairs, palindrome |
| Two Pointers (fast/slow) | Different speeds | O(n) | O(1) | Cycle detection, find middle |
| Sliding Window (fixed) | Add new, remove old | O(n) | O(1) | Max sum of k elements |
| Sliding Window (variable) | Expand/shrink with map | O(n) | O(k) | Longest substring conditions |

*Queue dequeue is O(n) with slice; O(1) with proper implementation

### Advanced Patterns Covered

1. **Map of pairs for matching** (balanced brackets)
2. **Auxiliary data structure for O(1) operations** (min stack)
3. **Mid calculation to prevent overflow** (binary search)
4. **Search insert position** (binary search returns left)
5. **Last index in map vs existence** (sliding window)
6. **Window boundary conditions** (idx >= left check)
7. **When to move pointers** (two sum logic)

---

*Lessons 6-10 cover: stacks/main.go, queues/main.go, binary_search/main.go, two_pointer/main.go, sliding_window/main.go*
