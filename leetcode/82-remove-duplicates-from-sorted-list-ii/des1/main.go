package des1

import (
	. "leetcode/leetcode/pkg/helper"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }.
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := head
	dummy := &ListNode{Next: head}
	prev := dummy
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			node = findNext(node)
			prev.Next = node
			continue
		}
		// move prev only when node.Val != node.Next.Val
		prev = node
		node = node.Next
	}
	return dummy.Next
}

func findNext(node *ListNode) *ListNode {
	for node.Next != nil && node.Val == node.Next.Val {
		node = node.Next
	}
	return node.Next
}
