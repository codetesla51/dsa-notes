# Data Structures & Algorithms - Complete Learning Guide

A comprehensive DSA curriculum with detailed Go implementations and extensive study notes. Every algorithm includes line-by-line comments, execution traces, complexity analysis, and real-world examples.

## What's Inside

- **10 Complete Lessons** covering fundamental to intermediate DSA concepts
- **4000+ lines of study notes** in `DSA_NOTES.md` with analogies, diagrams, and traces
- **Fully commented Go code** with step-by-step execution examples
- **ASCII diagrams** for visualizing data structures and algorithms
- **Complexity analysis** with detailed reasoning for every operation

## Full Curriculum

| Lesson | Topic | Key Concepts | Example Problems |
|--------|-------|--------------|------------------|
| 1 | **Big O Notation** | Time/Space complexity, Growth rates | Linear search, Constant vs Linear vs Quadratic |
| 2 | **Arrays** | Index access, Two pointers, In-place operations | Reverse array, Multiply by index, Remove duplicates |
| 3 | **Strings** | Immutability, Character operations, Iteration | Count vowels, Reverse string, Anagram check |
| 4 | **Hash Tables** | O(1) lookup, Frequency counting, Deduplication | Find duplicates, Two sum, Character frequency |
| 5 | **Linked Lists** | Node structure, Pointers, Traversal patterns | Reverse list, Detect cycle, Find middle, Delete node |
| 6 | **Stacks** | LIFO, Push/Pop, Backtracking | Balanced brackets, Min stack, Reverse string |
| 7 | **Queues** | FIFO, Enqueue/Dequeue, Order preservation | Process queue, Reverse queue, Even/odd separation |
| 8 | **Binary Search** | O(log n) search, Divide & conquer, Overflow prevention | Search sorted array, Find insert position |
| 9 | **Two Pointers** | Opposite ends, Fast/slow, Same direction | Two sum sorted, Palindrome check, Remove duplicates |
| 10 | **Sliding Window** | Fixed/variable windows, Optimization | Max sum k elements, Longest substring no repeats |

## Topics Covered

### Data Structures
- Arrays & Dynamic Arrays
- Strings & String Manipulation
- Hash Tables (Maps & Sets)
- Singly Linked Lists
- Stacks (LIFO)
- Queues (FIFO)

### Algorithms & Techniques
- Big O Analysis (Time & Space Complexity)
- Two Pointer Technique (3 patterns)
- Sliding Window (Fixed & Variable)
- Binary Search (O(log n) search)
- Fast & Slow Pointers (Cycle detection)
- Frequency Counting with Maps
- In-place Array Operations

### Problem-Solving Patterns
- Map to eliminate nested loops (O(n²) → O(n))
- Convert string to []byte for mutations
- Use map[T]int for counting, map[T]bool for existence
- Store last index in maps for variable windows
- Binary search: left + (right-left)/2 prevents overflow
- Queue with slices is O(n) dequeue, not O(1)

## Example Problems Solved

### Lesson 1: Big O Notation
- Linear search O(n)
- Constant time access O(1)
- Understanding growth rates

### Lesson 2: Arrays
- Reverse array in-place (two pointers)
- Multiply elements by their index
- Remove duplicates from sorted array

### Lesson 3: Strings
- Count vowels in string
- Reverse string using []byte
- Check if two strings are anagrams

### Lesson 4: Hash Tables
- Find first duplicate in array
- Count character frequency
- Two sum problem (O(n) solution)
- Find first non-repeating character

### Lesson 5: Linked Lists
- Reverse linked list iteratively
- Detect cycle (Floyd's algorithm)
- Find middle node (fast/slow pointers)
- Delete node at position

### Lesson 6: Stacks
- Check balanced brackets/parentheses
- Implement min stack with O(1) getMin
- Reverse string using stack

### Lesson 7: Queues
- Process tasks in FIFO order
- Reverse queue using stack
- Separate even/odd numbers maintaining order

### Lesson 8: Binary Search
- Search in sorted array O(log n)
- Find insertion position for target
- Prevent integer overflow with mid calculation

### Lesson 9: Two Pointers
- Two sum in sorted array
- Check if string is palindrome
- Remove duplicates in-place

### Lesson 10: Sliding Window
- Maximum sum of k consecutive elements (fixed window)
- Longest substring without repeating characters (variable window)
- Optimize from O(n²) to O(n)

## Repository Structure

```
algorithms/
├── DSA_NOTES.md          # Complete study guide (4000+ lines)
├── README.md             # This file
├── bigo/main.go          # Lesson 1: Big O examples
├── arrays/main.go        # Lesson 2: Array operations
├── strings/main.go       # Lesson 3: String manipulation
├── hashmaps/main.go      # Lesson 4: Hash table patterns
├── linkedlist/main.go    # Lesson 5: Linked list operations
├── stacks/main.go        # Lesson 6: Stack problems
├── queues/main.go        # Lesson 7: Queue problems
├── binary_search/main.go # Lesson 8: Binary search
├── two_pointer/main.go   # Lesson 9: Two pointer technique
├── sliding_window/main.go# Lesson 10: Sliding window
└── solu/main.go          # Additional practice problems
```

## How to Run

Each directory contains a standalone `main.go` file with multiple examples and test cases.

### Run a specific lesson:
```bash
cd arrays
go run main.go
```

### Run all tests in a directory:
```bash
cd hashmaps
go test -v
```

### Run all examples:
```bash
# From the root directory
for dir in bigo arrays strings hashmaps linkedlist stacks queues binary_search two_pointer sliding_window solu; do
  echo "Running $dir..."
  cd $dir && go run main.go && cd ..
done
```

## Learning Path

### Beginner Track (Start here)
1. Big O Notation - Understand complexity analysis
2. Arrays - Master index-based operations
3. Strings - Learn immutability and manipulation
4. Hash Tables - Unlock O(n) solutions

### Intermediate Track
5. Linked Lists - Practice pointer manipulation
6. Stacks - Understand LIFO and backtracking
7. Queues - Master FIFO and ordering
8. Binary Search - Learn divide & conquer

### Advanced Techniques
9. Two Pointers - Optimize array problems
10. Sliding Window - Master subarray optimization

## Code Documentation Style

Every implementation includes:
- **Real-world analogy** to explain the concept
- **Line-by-line comments** explaining WHY, not just what
- **Step-by-step execution trace** with actual values
- **ASCII diagrams** showing data structure state
- **Complexity analysis** with detailed reasoning
- **Multiple test cases** in main() function
- **Edge cases** and gotchas highlighted

## Key Insights & Patterns

### Time Complexity Optimizations
- Nested loops O(n²) → Map lookup O(n)
- Linear search O(n) → Binary search O(log n)
- Fixed window avoids recalculating sum
- Variable window with seen map tracks last index

### Go-Specific Tips
- Strings are immutable → convert to []byte for mutations
- Queue dequeue with slices is O(n), not O(1)
- Use `left + (right-left)/2` to prevent overflow
- Maps: map[T]int for counting, map[T]bool for existence
- Never loop through map to search, use direct indexing

### Problem-Solving Strategies
- Draw the state at each step
- Start with brute force, then optimize
- Ask: "Can a map help here?"
- Two pointers: opposite ends vs same direction
- Sliding window: fixed size vs variable size
- Always check edge cases: empty, single element, all same

## Study Notes

The `DSA_NOTES.md` file contains the complete curriculum with:
- Real-world analogies for every concept
- ASCII diagrams for visualization
- Step-by-step traces showing algorithm execution
- Complexity tables and analysis
- Code examples from actual implementations
- When to use / when not to use each pattern
- Common pitfalls and gotchas

## Contributing

This is a personal learning repository, but feel free to:
- Open issues for corrections or improvements
- Suggest additional problems or patterns
- Share your own solutions and optimizations

## Resources & References

- Introduction to Algorithms (CLRS)
- Cracking the Coding Interview
- LeetCode problem patterns
- Go documentation and best practices

## License

MIT License - Feel free to use this for your own learning journey!

---

**Happy Learning!** Start with `DSA_NOTES.md` for theory, then practice with the code examples.
