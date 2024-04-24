package main

import "fmt"

// ListNode defines a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		val1, val2 := 0, 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10
		newNode := &ListNode{Val: sum % 10}
		current.Next = newNode
		current = newNode
	}

	return dummyHead.Next
}

// Helper function to print list nodes
func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
	fmt.Println()
}

// main function to test the code
func main() {
	// Create first list: 2 -> 4 -> 3
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	// Create second list: 5 -> 6 -> 4
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	result := addTwoNumbers(l1, l2)
	printList(result) // Output should be 7 -> 0 -> 8
}
