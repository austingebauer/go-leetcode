// Package structures provides structures that can be used to create various data structures.
package structures

// TreeNode is a node in a binary tree data structure.
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// ListNode is a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
