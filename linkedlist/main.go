package main

import (
	"fmt"
)

// ============================================================================
// LINKED LISTS - Non-Contiguous Memory & Pointer Traversal
// ============================================================================
//
// WHAT IS A LINKED LIST?
// A collection of nodes where each node contains:
// 1. Data (the value)
// 2. Pointer to next node (memory address)
//
// Unlike arrays, nodes are NOT in contiguous memory.
// You can't jump to index 5 - must traverse from head.
//
// VISUAL:
// Head → [1|•]→[2|•]→[3|•]→[4|•]→[5|nil]
//        ────   ────   ────   ────   ────
//        Val=1  Val=2  Val=3  Val=4  Val=5
//        Next→  Next→  Next→  Next→  Next=nil
//
// MEMORY LAYOUT (scattered, NOT contiguous):
// ┌─────────────────────────────────────────┐
// │  Address  │  Value  │  Next Pointer     │
// ├─────────────────────────────────────────┤
// │   1000    │    1    │    2500           │  Node 1
// │   2500    │    2    │    7200           │  Node 2
// │   7200    │    3    │    4100           │  Node 3
// │   4100    │    4    │    9000           │  Node 4
// │   9000    │    5    │    nil            │  Node 5
// └─────────────────────────────────────────┘
//
// Compare to Array (contiguous):
// ┌──────┬──────┬──────┬──────┬──────┐
// │  1   │  2   │  3   │  4   │  5   │
// ├──────┼──────┼──────┼──────┼──────┤
// │ 1000 │ 1004 │ 1008 │ 1012 │ 1016 │
// └──────┴──────┴──────┴──────┴──────┘
//
// ============================================================================
// LINKED LIST OPERATIONS:
// ============================================================================
// Access by index:     O(n) - must traverse from head
// Search:              O(n) - must check each node
// Insert at head:      O(1) - just update head pointer
// Insert at tail:      O(n) - must traverse to end
// Insert in middle:    O(n) - must traverse to position
// Delete at head:      O(1) - just update head pointer
// Delete at tail:      O(n) - must find second-to-last node
// Delete in middle:    O(n) - must traverse to position
//
// WHEN TO USE LINKED LIST:
// ✓ Frequent insertions/deletions at beginning
// ✓ Don't need random access by index
// ✓ Size unknown or changes frequently
//
// WHEN TO USE ARRAY:
// ✓ Need random access (arr[i])
// ✓ Know size in advance
// ✓ Mostly reading, rare insertions
//
// ============================================================================

// Node represents a single element in the linked list
type Node struct {
	Val  int   // The data stored in this node
	Next *Node // Pointer to next node (or nil if this is the last node)
}

// ============================================================================

// revLinkedList reverses a linked list in-place
//
// COMPLEXITY:
// Time:  O(n) - visit each node exactly once
// Space: O(1) - only use 3 pointers regardless of list size
//
// PATTERN: prev/current/next for reversal
// - prev: tracks the node we just processed
// - current: the node we're currently processing
// - next: saves the next node before we break the link
//
// WHY SAVE next?
// We're about to change current.Next to point backwards.
// If we don't save the original next, we lose the rest of the list!
//
// EXAMPLE TRACE:
// Input: [1]→[2]→[3]→nil
//
// Initial state:
// current = [1]
// prev = nil
//
// ┌─────────────────────────────────────────────────┐
// │ Iteration 1: Reverse Node 1                     │
// └─────────────────────────────────────────────────┘
// Before:
//
//	prev    current   (rest of list)
//	nil      [1]   →   [2]→[3]→nil
//
// Step by step:
//
//	next = current.Next  →  next = [2]    (save the rest!)
//	current.Next = prev  →  [1]→nil       (reverse pointer)
//	prev = current       →  prev = [1]    (move prev forward)
//	current = next       →  current = [2] (move current forward)
//
// After:
//
//	       prev    current
//	nil←──[1]      [2]→[3]→nil
//
// ┌─────────────────────────────────────────────────┐
// │ Iteration 2: Reverse Node 2                     │
// └─────────────────────────────────────────────────┘
// Before:
//
//	       prev    current
//	nil←──[1]      [2]→[3]→nil
//
// Step by step:
//
//	next = [3]              (save the rest)
//	current.Next = prev     ([2]→[1])
//	prev = current          (prev now [2])
//	current = next          (current now [3])
//
// After:
//
//	             prev    current
//	nil←[1]←──[2]        [3]→nil
//
// ┌─────────────────────────────────────────────────┐
// │ Iteration 3: Reverse Node 3                     │
// └─────────────────────────────────────────────────┘
// Before:
//
//	             prev    current
//	nil←[1]←[2]         [3]→nil
//
// Step by step:
//
//	next = nil              (no more nodes)
//	current.Next = prev     ([3]→[2])
//	prev = current          (prev now [3])
//	current = next          (current now nil)
//
// After:
//
//	                   prev    current
//	nil←[1]←[2]←[3]          nil
//
// Loop exits (current == nil)
// Return prev (the new head)
//
// Result: nil←[1]←[2]←[3]
//
//	   ↑
//	new head
func revLinkedList(head *Node) *Node {
	// Start at the beginning
	current := head

	// prev starts as nil (will become the new tail's Next pointer)
	var prev *Node

	// Process each node until we reach the end
	for current != nil {
		// CRITICAL: Save the next node before we break the link
		// We're about to overwrite current.Next
		// Without this, we'd lose the rest of the list!
		next := current.Next

		// Reverse the pointer!
		// Point current node back to previous node
		// This is the actual "reversal" operation
		current.Next = prev

		// Move prev forward to current node
		// prev is now the last node we reversed
		prev = current

		// Move current forward to next node
		// Uses the next pointer we saved earlier
		current = next
	}

	// When loop ends:
	// - current is nil (fell off the end)
	// - prev is the last node we processed (new head)
	return prev
}

// ============================================================================

// printList is a helper to visualize the linked list
func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " → ")
		current = current.Next
	}
	fmt.Println("nil")
}

// ============================================================================

// getAt returns the value at a specific index (0-based)
// Similar to arr[index] but requires traversal
//
// COMPLEXITY:
// Time:  O(n) - worst case, traverse to end
// Space: O(1) - only use one pointer
//
// EXAMPLE TRACE:
// Input: [1]→[2]→[3]→[4]→[5]→nil, index = 2
//
// Initial: current = [1], i = 0
//
// Iteration 1: i=0, i < 2? Yes
//
//	current = current.Next → current = [2]
//	i++ → i = 1
//
// Iteration 2: i=1, i < 2? Yes
//
//	current = current.Next → current = [3]
//	i++ → i = 2
//
// Iteration 3: i=2, i < 2? No
//
//	Exit loop
//
// Return: current.Val = 3
func getAt(head *Node, index int) int {
	current := head

	// Move forward 'index' times
	// To get index 2: start at 0, move to 1, move to 2 (2 moves)
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Val
}

// ============================================================================

// middleLinkedList finds the middle element value
// Method: Count length, then traverse to middle
//
// COMPLEXITY:
// Time:  O(n) - two passes (count + traverse to middle)
// Space: O(1) - only a few variables
//
// EXAMPLE TRACE:
// Input: [1]→[2]→[3]→[4]→[5]→nil
//
// Pass 1: Count total nodes
//
//	current = [1], count = 0
//	Iteration 1: count = 1, current = [2]
//	Iteration 2: count = 2, current = [3]
//	Iteration 3: count = 3, current = [4]
//	Iteration 4: count = 4, current = [5]
//	Iteration 5: count = 5, current = nil
//	Exit loop
//	Result: count = 5
//
// Calculate middle index:
//
//	mid = 5 / 2 = 2 (integer division)
//
// Pass 2: Get element at index 2 (using getAt)
//
//	Start at [1], move twice → [3]
//	Return: 3
//
// Result: 3 (the middle element)
//
// ALTERNATIVE: Fast/Slow Pointer (one pass)
// slow moves 1 step, fast moves 2 steps
// When fast reaches end, slow is at middle
// See hasCycle() for fast/slow pattern
func middleLinkedList(head *Node) int {
	current := head
	count := 0

	// First pass: count total nodes
	for current != nil {
		count++                // Increment counter
		current = current.Next // Move to next node
	}

	// Calculate middle index
	// For odd length (5): 5/2 = 2 (positions: 0,1,2,3,4)
	// For even length (4): 4/2 = 2 (positions: 0,1,2,3)
	mid := count / 2

	// Second pass: traverse to middle index
	return getAt(head, mid)
}

// ============================================================================

// hasCycle detects if linked list has a cycle
// Uses Floyd's "Tortoise and Hare" algorithm
//
// COMPLEXITY:
// Time:  O(n) - fast pointer travels at most 2n steps
// Space: O(1) - only two pointers
//
// PATTERN: Fast/Slow Pointer
// - Slow moves 1 step per iteration
// - Fast moves 2 steps per iteration
// - If there's a cycle, fast will eventually catch slow
// - Like runners on a track - faster runner laps slower runner
//
// EXAMPLE TRACE 1: List with cycle
// [1]→[2]→[3]→[4]→[5]
//
//	↑         ↓
//	└─────────┘
//
// Initial:
//
//	slow = [1], fast = [1]
//
// Iteration 1:
//
//	slow moves 1: slow = [2]
//	fast moves 2: fast = [3]
//	slow == fast? No
//
// Iteration 2:
//
//	slow moves 1: slow = [3]
//	fast moves 2: fast = [5] (from [3]→[4]→[5])
//	slow == fast? No
//
// Iteration 3:
//
//	slow moves 1: slow = [4]
//	fast moves 2: fast = [3] (from [5]→[2]→[3] via cycle)
//	slow == fast? No
//
// Iteration 4:
//
//	slow moves 1: slow = [5]
//	fast moves 2: fast = [5] (from [3]→[4]→[5])
//	slow == fast? YES!
//	Return: true (cycle detected)
//
// EXAMPLE TRACE 2: List without cycle
// [1]→[2]→[3]→[4]→[5]→nil
//
// Initial:
//
//	slow = [1], fast = [1]
//
// Iteration 1:
//
//	slow = [2], fast = [3]
//
// Iteration 2:
//
//	slow = [3], fast = [5]
//
// Iteration 3:
//
//	fast.Next = nil
//	Loop condition false (fast.Next == nil)
//	Exit loop
//	Return: false (no cycle)
func hasCycle(head *Node) bool {
	// Both pointers start at head
	slow := head
	fast := head

	// Continue while fast can move 2 steps
	// Must check TWO conditions:
	// 1. fast != nil (fast exists)
	// 2. fast.Next != nil (can take second step)
	for fast != nil && fast.Next != nil {
		// Tortoise: move 1 step
		slow = slow.Next

		// Hare: move 2 steps
		// This is why we check fast.Next != nil
		fast = fast.Next.Next

		// Did they meet?
		if slow == fast {
			// Pointers met! Must be a cycle
			// In a cycle, fast will eventually "lap" slow
			return true
		}
	}

	// Reached end of list (fast hit nil)
	// No cycle exists
	return false
}

// ============================================================================

// delNode deletes all nodes with a specific value
//
// COMPLEXITY:
// Time:  O(n) - single pass through list
// Space: O(1) - only prev and current pointers
//
// PATTERN: prev/current for deletion
// - prev tracks previous node (for linking around deleted node)
// - When deleting: prev.Next = current.Next (skip current)
// - When keeping: prev = current (advance prev)
//
// CRITICAL: Handle head deletion separately!
// If head matches val, keep moving head forward until it doesn't
//
// EXAMPLE TRACE:
// Input: [1]→[2]→[3]→[2]→[4]→nil, val = 2
//
// Step 1: Handle head deletions
//
//	head.Val = 1, not equal to 2
//	Skip this loop
//
// Step 2: Initialize
//
//	current = [1]
//	prev = nil
//
// Step 3: Traverse and delete
//
// Iteration 1: current=[1]
//
//	1 != 2 (keep this node)
//	prev = [1]
//	current = [2]
//	List: [1]→[2]→[3]→[2]→[4]
//	       ↑    ↑
//	      prev curr
//
// Iteration 2: current=[2]
//
//	2 == 2 (delete this node!)
//	prev.Next = current.Next
//	[1].Next = [3]
//	prev stays [1] (DON'T advance on deletion!)
//	current = [3]
//	List: [1]→[3]→[2]→[4]
//	       ↑    ↑
//	      prev curr
//
// Iteration 3: current=[3]
//
//	3 != 2 (keep this node)
//	prev = [3]
//	current = [2]
//	List: [1]→[3]→[2]→[4]
//	            ↑    ↑
//	           prev curr
//
// Iteration 4: current=[2]
//
//	2 == 2 (delete this node!)
//	prev.Next = current.Next
//	[3].Next = [4]
//	prev stays [3]
//	current = [4]
//	List: [1]→[3]→[4]
//	            ↑    ↑
//	           prev curr
//
// Iteration 5: current=[4]
//
//	4 != 2 (keep this node)
//	prev = [4]
//	current = nil
//	Exit loop
//
// Result: [1]→[3]→[4]→nil
//
// WHY DON'T WE ADVANCE prev ON DELETION?
// When we delete current, prev should still point to the same node.
// That node now points to current.Next (skipping deleted node).
// Advancing prev would skip checking the next node!
func delNode(head *Node, val int) *Node {
	var prev *Node

	// CRITICAL: Handle deletion at head
	// Keep moving head forward while it matches val
	// Example: [2]→[2]→[1] with val=2 → head becomes [1]
	for head != nil && head.Val == val {
		head = head.Next
	}

	// After this loop, head is either:
	// 1. nil (all nodes matched val)
	// 2. First node that doesn't match val

	current := head

	// Process remaining nodes
	for current != nil {
		if current.Val == val {
			// Found node to delete
			// Bridge over it: prev → current → next
			//        becomes: prev → next (skip current)
			prev.Next = current.Next

			// DON'T update prev!
			// prev stays pointing to same node
			// because current is being deleted
		} else {
			// Keep this node
			// Update prev to current node
			prev = current
		}

		// Move to next node regardless of deletion
		current = current.Next
	}

	// Return new head (may have changed if original head was deleted)
	return head
}

// ============================================================================

func main() {
	// Create a linked list: [1]→[2]→[3]→[4]→[5]→nil
	n5 := &Node{Val: 5}
	n4 := &Node{Val: 4, Next: n5}
	n3 := &Node{Val: 3, Next: n4}
	n2 := &Node{Val: 2, Next: n3}
	n1 := &Node{Val: 1, Next: n2}

	fmt.Println("=== Original List ===")
	printList(n1)

	// Test reversal
	fmt.Println("\n=== Reverse List ===")
	reversed := revLinkedList(n1)
	printList(reversed)

	// Reverse back to original
	original := revLinkedList(reversed)
	printList(original)

	// Test middle
	fmt.Println("\n=== Find Middle ===")
	fmt.Println("Middle value:", middleLinkedList(original))

	// Test getAt
	fmt.Println("\n=== Access by Index ===")
	for i := 0; i < 5; i++ {
		fmt.Printf("Index %d: %d\n", i, getAt(original, i))
	}

	// Test deletion
	fmt.Println("\n=== Delete Node ===")
	fmt.Println("Before deletion:")
	printList(original)
	fmt.Println("Deleting value 2...")
	afterDelete := delNode(original, 2)
	fmt.Println("After deletion:")
	printList(afterDelete)

	// Test cycle detection
	fmt.Println("\n=== Cycle Detection ===")
	// Create list without cycle
	a := &Node{Val: 1}
	b := &Node{Val: 2}
	c := &Node{Val: 3}
	a.Next = b
	b.Next = c
	fmt.Println("List without cycle:", hasCycle(a))

	// Create cycle: [1]→[2]→[3]→[2] (3 points back to 2)
	c.Next = b
	fmt.Println("List with cycle:", hasCycle(a))
}
