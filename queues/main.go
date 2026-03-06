package main

import (
	"fmt"
)

// ============================================================================
// QUEUES - First In First Out (FIFO)
// ============================================================================
//
// WHAT IS A QUEUE?
// Think of a line at a coffee shop:
// - People join at the BACK of the line
// - People are served from the FRONT of the line
// - First person in line is first to be served (FIFO)
// - Can't cut in the middle!
//
// VISUAL:
//   FRONT                       BACK
//     ↓                          ↓
//   ┌───┬───┬───┬───┬───┐
//   │ 1 │ 2 │ 3 │ 4 │ 5 │
//   └───┴───┴───┴───┴───┘
//    ↑                   ↑
//  Dequeue             Enqueue
//  (remove)             (add)
//
// CORE OPERATIONS:
// - Enqueue: Add element to back    → O(1) with slice
// - Dequeue: Remove element from front → O(n) with slice*
// - Front/Peek: Look at front element → O(1)
// - isEmpty: Check if queue is empty → O(1)
//
// *WHY DEQUEUE IS O(n) NOT O(1):
// When we do queue = queue[1:], Go must:
// 1. Create a new slice header pointing to index 1
// 2. Eventually, garbage collect the old elements
// 3. If underlying array gets too sparse, reallocate
//
// For true O(1) dequeue, use:
// - Circular buffer
// - Linked list
// - Two stacks (for interview trick questions)
//
// GO IMPLEMENTATION:
//   queue := []int{}
//   queue = append(queue, x)  // Enqueue (add to back)
//   front := queue[0]         // Peek at front
//   queue = queue[1:]         // Dequeue (remove from front) - O(n)!
//
// WHEN TO USE QUEUE:
// ✓ Process tasks in order received (job queue)
// ✓ Breadth-first search (BFS)
// ✓ Print queue, request handling
// ✓ Level-order tree traversal
//
// STACK vs QUEUE:
// Stack: LIFO (Last In First Out)  - like a stack of plates
// Queue: FIFO (First In First Out) - like a line of people
//
// ============================================================================

// processQueue processes tasks in FIFO order
//
// COMPLEXITY:
// Time:  O(n²) - n enqueues O(1) + n dequeues O(n each)
// Space: O(n) - queue stores all tasks at peak
//
// WHY O(n²)?
// Each dequeue with queue = queue[1:] is O(n)
// We dequeue n times, so n * O(n) = O(n²)
//
// EXAMPLE TRACE:
// Input: tasks = ["task1", "task2", "task3"]
//
// Step 1: Enqueue all tasks
//
//	Enqueue "task1": queue = ["task1"]
//	Enqueue "task2": queue = ["task1", "task2"]
//	Enqueue "task3": queue = ["task1", "task2", "task3"]
//
// Step 2: Dequeue and process (FIFO order)
//
//	Dequeue: front = "task1", queue = ["task2", "task3"]
//	Process: "task1"
//
//	Dequeue: front = "task2", queue = ["task3"]
//	Process: "task2"
//
//	Dequeue: front = "task3", queue = []
//	Process: "task3"
//
// Result: Tasks processed in order received
func processQueue(tasks []string) {
	// Create empty queue
	queue := []string{}

	// ENQUEUE: Add all tasks to queue (back of line)
	for _, t := range tasks {
		// append adds to end of slice → back of queue
		// This is O(1) amortized
		queue = append(queue, t)
	}

	// DEQUEUE: Process tasks in FIFO order (from front)
	for len(queue) > 0 {
		// PEEK: Get front element (first in line)
		// queue[0] is the front of the queue
		front := queue[0]

		// DEQUEUE: Remove front element
		// queue[1:] creates new slice starting from index 1
		// This is O(n) because Go may need to shift elements
		queue = queue[1:]

		// Process the task
		fmt.Println("processing", front)
	}
}

// ============================================================================

// revQueue reverses a queue using a stack
// Demonstrates the difference between LIFO and FIFO
//
// COMPLEXITY:
// Time:  O(n²) - dequeue is O(n), done n times
// Space: O(n) - stack stores all elements
//
// EXAMPLE TRACE:
// Input: queue = [1, 2, 3, 4, 5]
//
// Step 1: Transfer queue to stack (using FIFO)
//
//	Dequeue 1: stack = [1], queue = [2,3,4,5]
//	Dequeue 2: stack = [1,2], queue = [3,4,5]
//	Dequeue 3: stack = [1,2,3], queue = [4,5]
//	Dequeue 4: stack = [1,2,3,4], queue = [5]
//	Dequeue 5: stack = [1,2,3,4,5], queue = []
//
// Step 2: Pop from stack to result (using LIFO)
//
//	Pop: top=5, res=[5], stack=[1,2,3,4]
//	Pop: top=4, res=[5,4], stack=[1,2,3]
//	Pop: top=3, res=[5,4,3], stack=[1,2]
//	Pop: top=2, res=[5,4,3,2], stack=[1]
//	Pop: top=1, res=[5,4,3,2,1], stack=[]
//
// Result: [5, 4, 3, 2, 1]
//
// WHY THIS WORKS:
// Queue (FIFO): 1→2→3→4→5 goes into stack as [1,2,3,4,5]
// Stack (LIFO): pops as 5→4→3→2→1
func revQueue(queue []int) []int {
	// Use a stack to reverse
	stack := []int{}

	// Dequeue all elements and push to stack
	for len(queue) > 0 {
		// Get front of queue
		front := queue[0]

		// Remove from queue (O(n) operation!)
		queue = queue[1:]

		// Push to stack (O(1) operation)
		stack = append(stack, front)
	}

	// Result slice
	res := []int{}

	// Pop all elements from stack (LIFO gives us reverse)
	for len(stack) > 0 {
		// Peek at top
		top := stack[len(stack)-1]

		// Pop from stack
		stack = stack[:len(stack)-1]

		// Add to result
		res = append(res, top)
	}

	return res
}

// ============================================================================

// howManyProcessesd counts how many elements were processed from queue
//
// COMPLEXITY:
// Time:  O(n²) - n dequeues at O(n) each
// Space: O(n) - queue stores all elements
//
// EXAMPLE TRACE:
// Input: nums = [10, 20, 30]
//
// Enqueue phase:
//
//	queue = [10, 20, 30]
//
// Dequeue and count:
//
//	Dequeue 10: count=1, queue=[20,30]
//	Dequeue 20: count=2, queue=[30]
//	Dequeue 30: count=3, queue=[]
//
// Result: 3
func howManyProcessesd(nums []int) int {
	// Create queue
	queue := []int{}

	// Enqueue all numbers
	for _, n := range nums {
		queue = append(queue, n)
	}

	// Counter for processed elements
	count := 0

	// Dequeue and count each element
	for len(queue) > 0 {
		// Get front
		front := queue[0]

		// Dequeue (O(n))
		queue = queue[1:]

		// Process
		fmt.Println(front)

		// Increment counter
		count++
	}

	return count
}

// ============================================================================

// evenQueue separates odd and even numbers using two queues
// Processes odds first, then evens
//
// COMPLEXITY:
// Time:  O(n²) - dequeuing is expensive
// Space: O(n) - two queues store all elements
//
// EXAMPLE TRACE:
// Input: nums = [1, 2, 3, 4, 5, 6]
//
// Step 1: Separate into two queues
//
//	1 is odd → queue2 = [1]
//	2 is even → queue1 = [2]
//	3 is odd → queue2 = [1, 3]
//	4 is even → queue1 = [2, 4]
//	5 is odd → queue2 = [1, 3, 5]
//	6 is even → queue1 = [2, 4, 6]
//
// Step 2: Process odds first (queue2)
//
//	Dequeue 1: res = [1]
//	Dequeue 3: res = [1, 3]
//	Dequeue 5: res = [1, 3, 5]
//
// Step 3: Process evens (queue1)
//
//	Dequeue 2: res = [1, 3, 5, 2]
//	Dequeue 4: res = [1, 3, 5, 2, 4]
//	Dequeue 6: res = [1, 3, 5, 2, 4, 6]
//
// Result: [1, 3, 5, 2, 4, 6]
func evenQueue(nums []int) []int {
	// Two queues: one for even, one for odd
	queue1 := []int{} // Even numbers
	queue2 := []int{} // Odd numbers
	res := []int{}

	// Separate numbers into appropriate queues
	for _, n := range nums {
		if n%2 == 0 {
			// Even number → queue1
			queue1 = append(queue1, n)
		} else {
			// Odd number → queue2
			queue2 = append(queue2, n)
		}
	}

	// Process odd queue first (queue2)
	for len(queue2) > 0 {
		// Dequeue from front
		front := queue2[0]
		queue2 = queue2[1:]

		// Add to result
		res = append(res, front)
	}

	// Then process even queue (queue1)
	for len(queue1) > 0 {
		// Dequeue from front
		front := queue1[0]
		queue1 = queue1[1:]

		// Add to result
		res = append(res, front)
	}

	return res
}

// ============================================================================

func main() {
	// Demonstrate process queue
	fmt.Println("=== Process Queue ===")
	tasks := []string{"task1", "task2", "task3"}
	processQueue(tasks)

	// Demonstrate reverse queue
	fmt.Println("\n=== Reverse Queue ===")
	queue := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Original:", queue)
	fmt.Println("Reversed:", revQueue(queue))

	// Demonstrate count processed
	fmt.Println("\n=== Count Processed ===")
	nums := []int{10, 20, 30, 40}
	count := howManyProcessesd(nums)
	fmt.Println("Total processed:", count)

	// Demonstrate even/odd separation
	fmt.Println("\n=== Separate Even/Odd ===")
	mixed := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Input:", mixed)
	fmt.Println("Odds first, then evens:", evenQueue(mixed))
}
