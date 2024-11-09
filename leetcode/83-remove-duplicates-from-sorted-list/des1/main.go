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
	node := head
	var next *ListNode
	for node != nil {
		if node.Next == nil {
			break
		}
		if node.Val == node.Next.Val {
			next = node.Next
			for next != nil && next.Val == node.Val {
				next = next.Next
			}
			node.Next = next
		}
		node = node.Next
	}
	return head
}
