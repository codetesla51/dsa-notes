package main

import "fmt"

// ============================================================================
// STACKS - Last In First Out (LIFO)
// ============================================================================
//
// WHAT IS A STACK?
// Think of a stack of plates:
// - You add plates to the TOP
// - You remove plates from the TOP
// - You can't pull a plate from the middle
// - Last plate added is first plate removed (LIFO)
//
// VISUAL:
//     TOP
//      ↓
//    ┌───┐
//    │ 5 │  ← Most recently added (top)
//    ├───┤
//    │ 4 │
//    ├───┤
//    │ 3 │
//    ├───┤
//    │ 2 │
//    ├───┤
//    │ 1 │  ← First added (bottom)
//    └───┘
//
// CORE OPERATIONS:
// - Push:    Add element to top       → O(1)
// - Pop:     Remove element from top  → O(1)
// - Peek:    Look at top without removing → O(1)
// - isEmpty: Check if stack is empty  → O(1)
//
// GO IMPLEMENTATION:
// Go doesn't have built-in stack, use slice:
//   stack := []int{}
//   stack = append(stack, x)        // Push
//   top := stack[len(stack)-1]      // Peek
//   stack = stack[:len(stack)-1]    // Pop
//
// WHEN TO USE STACK:
// ✓ Balanced parentheses/brackets
// ✓ Undo/redo functionality
// ✓ Function call stack (recursion)
// ✓ Reverse a sequence
// ✓ Track min/max in O(1) with auxiliary stack
//
// ============================================================================

// Stack implements a min-stack using two slices
// Tracks both main values and minimum value at each level
//
// WHY TWO STACKS?
// - main: stores all values
// - min: stores minimum value at each stack height
// - This allows getMin() in O(1) time
type Stack struct {
	main []int // Primary stack storing all values
	min  []int // Auxiliary stack tracking minimum at each level
}

// ============================================================================

// balancedBracket checks if all brackets/parentheses are properly matched
//
// COMPLEXITY:
// Time:  O(n) - single pass through string
// Space: O(n) - worst case, all opening brackets on stack
//
// PATTERN: Stack for matching pairs
// - Opening bracket: push onto stack
// - Closing bracket: pop and verify it matches
// - Empty stack at end means balanced
//
// EXAMPLE TRACE:
// Input: s = "{[()]}"
//
// Initial:
// stack = []
// pairs = {')':'(', '}':'{', ']':'['}
//
// Iteration 1: ch = '{'
//
//	Is opening bracket? Yes
//	Push: stack = ['{']
//
// Iteration 2: ch = '['
//
//	Is opening bracket? Yes
//	Push: stack = ['{', '[']
//
// Iteration 3: ch = '('
//
//	Is opening bracket? Yes
//	Push: stack = ['{', '[', '(']
//
// Iteration 4: ch = ')'
//
//	Is closing bracket? Yes
//	Stack empty? No
//	top = '(' (peek)
//	Pop: stack = ['{', '[']
//	Does top match pairs[')']? '(' == '(' ✓
//
// Iteration 5: ch = ']'
//
//	Is closing bracket? Yes
//	Stack empty? No
//	top = '[' (peek)
//	Pop: stack = ['{']
//	Does top match pairs[']']? '[' == '[' ✓
//
// Iteration 6: ch = '}'
//
//	Is closing bracket? Yes
//	Stack empty? No
//	top = '{' (peek)
//	Pop: stack = []
//	Does top match pairs['}']? '{' == '{' ✓
//
// After loop:
//
//	Is stack empty? Yes → return true
//
// Result: true (all brackets balanced)
//
// COMMON ERRORS:
// - Too many closing: ")(" → stack empty when trying to pop
// - Too many opening: "(((" → stack not empty at end
// - Mismatched: "([)]" → wrong bracket type when popping
func balancedBracket(s string) bool {
	// Stack to track opening brackets
	// Use rune slice to handle any Unicode brackets
	stack := []rune{}

	// Map closing brackets to their matching opening brackets
	// Key: closing bracket, Value: expected opening bracket
	// This makes verification O(1) instead of multiple if checks
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Process each character in the string
	for _, ch := range s {
		// Check if it's an opening bracket
		if ch == '(' || ch == '{' || ch == '[' {
			// PUSH: Add opening bracket to stack
			// We'll match it with a closing bracket later
			stack = append(stack, ch)

		} else {
			// It's a closing bracket

			// ERROR CHECK: If stack is empty, we have too many closing brackets
			// Example: ")" with no opening bracket before it
			if len(stack) == 0 {
				return false
			}

			// PEEK: Look at the most recent opening bracket
			// Don't remove it yet - we need to verify it matches
			top := stack[len(stack)-1]

			// POP: Remove the top element from stack
			// Slice from start to second-to-last element
			stack = stack[:len(stack)-1]

			// VERIFY: Does the opening bracket match this closing bracket?
			// pairs[ch] gives us the expected opening bracket
			// Example: pairs[')'] = '(', so we check if top == '('
			if top != pairs[ch] {
				return false // Mismatched brackets like "([)]"
			}
		}
	}

	// Final check: All brackets should be matched
	// If stack is empty, every opening had a matching closing
	// If stack has elements, we have unclosed opening brackets
	return len(stack) == 0
}

// ============================================================================

// revSlice reverses a slice using a stack
//
// COMPLEXITY:
// Time:  O(n) - push all elements + pop all elements
// Space: O(n) - stack stores all n elements
//
// DEMONSTRATES: LIFO reversal property of stacks
//
// EXAMPLE TRACE:
// Input: arr = [1, 2, 3, 4, 5]
//
// Step 1: Push all elements onto stack
//
//	Push 1: stack = [1]
//	Push 2: stack = [1, 2]
//	Push 3: stack = [1, 2, 3]
//	Push 4: stack = [1, 2, 3, 4]
//	Push 5: stack = [1, 2, 3, 4, 5]
//
// Step 2: Pop all elements (LIFO order)
//
//	Pop: top=5, result=[5], stack=[1,2,3,4]
//	Pop: top=4, result=[5,4], stack=[1,2,3]
//	Pop: top=3, result=[5,4,3], stack=[1,2]
//	Pop: top=2, result=[5,4,3,2], stack=[1]
//	Pop: top=1, result=[5,4,3,2,1], stack=[]
//
// Result: [5, 4, 3, 2, 1]
//
// WHY THIS WORKS:
// Stack is LIFO: last in, first out
// Elements come out in reverse order of insertion
func revSlice(arr []int) []int {
	// Create empty stack
	stack := []int{}

	// Push all elements onto stack
	for _, n := range arr {
		// PUSH operation: append to end of slice
		stack = append(stack, n)
	}

	// Result slice to collect reversed elements
	var result []int

	// Pop all elements (they come out in reverse order)
	for len(stack) > 0 {
		// PEEK: Look at top element
		top := stack[len(stack)-1]

		// POP: Remove top element
		// Take slice from start to second-to-last element
		stack = stack[:len(stack)-1]

		// Add popped element to result
		result = append(result, top)
	}

	return result
}

// ============================================================================

// validParenOnlyParne checks if parentheses are balanced (only '(' and ')')
// Simplified version of balancedBracket for single bracket type
//
// COMPLEXITY:
// Time:  O(n) - single pass
// Space: O(n) - worst case stack size
//
// EXAMPLE TRACE:
// Input: s = "(())"
//
// Iteration 1: ch='('
//
//	Push: stack = ['(']
//
// Iteration 2: ch='('
//
//	Push: stack = ['(', '(']
//
// Iteration 3: ch=')'
//
//	Stack empty? No
//	top='(', pop: stack = ['(']
//	Match? '(' == pair[')'] = '(' ✓
//
// Iteration 4: ch=')'
//
//	Stack empty? No
//	top='(', pop: stack = []
//	Match? '(' == pair[')'] = '(' ✓
//
// Stack empty? Yes → return true
func validParenOnlyParne(s string) bool {
	// Stack to track opening parentheses
	stack := []rune{}

	// Map for matching (simpler than balancedBracket - only one pair)
	pair := map[rune]rune{
		')': '(',
	}

	for _, ch := range s {
		if ch == '(' {
			// Opening parenthesis - push onto stack
			stack = append(stack, ch)
		} else {
			// Closing parenthesis

			// Check if we have an opening to match
			if len(stack) == 0 {
				return false // Too many closing
			}

			// Get and remove top element
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Verify it matches
			if top != pair[ch] {
				return false
			}
		}
	}

	// All parentheses should be matched
	return len(stack) == 0
}

// ============================================================================
// MIN STACK IMPLEMENTATION
// ============================================================================
//
// CHALLENGE: Design a stack that supports getMin() in O(1) time
//
// NAIVE APPROACH (doesn't work):
//   - Loop through stack to find min → O(n) time
//
// CLEVER SOLUTION: Two stacks
//   - main stack: stores all values
//   - min stack: stores minimum value at each height
//
// EXAMPLE:
//   Push(3): main=[3],     min=[3]     (3 is min)
//   Push(5): main=[3,5],   min=[3,3]   (3 still min)
//   Push(2): main=[3,5,2], min=[3,3,2] (2 is new min)
//   Push(1): main=[3,5,2,1], min=[3,3,2,1] (1 is new min)
//
//   getMin() → min[top] = 1 (O(1) time!)
//
//   Pop(): main=[3,5,2], min=[3,3,2]
//   getMin() → min[top] = 2 (O(1) time!)
//
// KEY INSIGHT:
// min stack tracks minimum at EACH LEVEL
// When we pop, the previous minimum is revealed
//
// ============================================================================

// Push adds a value to the stack and updates min stack
//
// COMPLEXITY:
// Time:  O(1) - constant time operations
// Space: O(1) - adds one element to each stack
//
// EXAMPLE TRACE:
// Start: main=[], min=[]
//
// Push(3):
//
//	main = [3]
//	min empty? Yes → min = [3]
//	State: main=[3], min=[3]
//
// Push(5):
//
//	main = [3, 5]
//	5 <= 3? No → min = [3, 3] (keep current min)
//	State: main=[3,5], min=[3,3]
//
// Push(2):
//
//	main = [3, 5, 2]
//	2 <= 3? Yes → min = [3, 3, 2] (new min)
//	State: main=[3,5,2], min=[3,3,2]
//
// Push(1):
//
//	main = [3, 5, 2, 1]
//	1 <= 2? Yes → min = [3, 3, 2, 1] (new min)
//	State: main=[3,5,2,1], min=[3,3,2,1]
func (s *Stack) Push(val int) {
	// Add value to main stack
	s.main = append(s.main, val)

	// Update min stack
	if len(s.min) == 0 || val <= s.min[len(s.min)-1] {
		// First element OR new minimum found
		// Push the new minimum
		s.min = append(s.min, val)
	} else {
		// val is NOT a new minimum
		// Duplicate the current minimum to maintain stack height
		// This ensures min stack has same length as main stack
		s.min = append(s.min, s.min[len(s.min)-1])
	}
}

// ============================================================================

// Pop removes the top element from both stacks
//
// COMPLEXITY: O(1)
//
// WHY POP FROM BOTH?
// min stack must stay in sync with main stack
// Otherwise heights won't match and getMin() breaks
func (s *Stack) Pop() {
	// Remove top from main stack
	s.main = s.main[:len(s.main)-1]

	// Remove top from min stack (keep them in sync!)
	s.min = s.min[:len(s.min)-1]
}

// ============================================================================

// Peek returns the top element without removing it
//
// COMPLEXITY: O(1)
func (s *Stack) Peek() int {
	// Access last element (top of stack)
	top := s.main[len(s.main)-1]
	return top
}

// ============================================================================

// getMin returns the minimum element in O(1) time
//
// COMPLEXITY: O(1) - this is the magic!
//
// HOW IT WORKS:
// min stack always has current minimum at its top
// No need to search through main stack
func (s *Stack) getMin() int {
	// Top of min stack is always the current minimum
	min := s.min[len(s.min)-1]
	return min
}

// ============================================================================

func main() {
	// Demonstrate balanced brackets
	fmt.Println("=== Balanced Brackets ===")
	fmt.Println("{[()]}", balancedBracket("{[()]}")) // true
	fmt.Println("([)]", balancedBracket("([)]"))     // false (mismatched)
	fmt.Println("((", balancedBracket("(("))         // false (unclosed)
	fmt.Println("))", balancedBracket("))"))         // false (too many closing)

	// Demonstrate stack reversal
	fmt.Println("\n=== Reverse with Stack ===")
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", arr)
	fmt.Println("Reversed:", revSlice(arr))

	// Demonstrate min stack
	fmt.Println("\n=== Min Stack ===")
	NewStack := &Stack{
		main: []int{},
		min:  []int{},
	}

	fmt.Println("Push(1):")
	NewStack.Push(1)
	fmt.Println("  main:", NewStack.main, "min:", NewStack.min)
	fmt.Println("  getMin():", NewStack.getMin())

	fmt.Println("\nPush(3):")
	NewStack.Push(3)
	fmt.Println("  main:", NewStack.main, "min:", NewStack.min)
	fmt.Println("  getMin():", NewStack.getMin())

	fmt.Println("\nPush(4):")
	NewStack.Push(4)
	fmt.Println("  main:", NewStack.main, "min:", NewStack.min)
	fmt.Println("  getMin():", NewStack.getMin())

	fmt.Println("\nPush(5):")
	NewStack.Push(5)
	fmt.Println("  main:", NewStack.main, "min:", NewStack.min)
	fmt.Println("  getMin():", NewStack.getMin())

	fmt.Println("\nPop():")
	NewStack.Pop()
	fmt.Println("  main:", NewStack.main, "min:", NewStack.min)
	fmt.Println("  getMin():", NewStack.getMin())

	// Demonstrate parentheses
	fmt.Println("\n=== Valid Parentheses ===")
	fmt.Println("(())", validParenOnlyParne("(())")) // true
	fmt.Println("()()", validParenOnlyParne("()()")) // true
	fmt.Println("(()", validParenOnlyParne("(()"))   // false
}
