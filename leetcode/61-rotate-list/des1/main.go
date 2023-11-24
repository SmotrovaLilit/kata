package des1

import (
	. "leetcode/leetcode/pkg/helper"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length := 0
	node := head
	for node != nil && node.Next != nil {
		length++
		node = node.Next
	}
	length++
	tail := node
	k = k % length
	if k == 0 {
		return head
	}
	node = head
	for i := 0; i < length-k-1; i++ {
		node = node.Next
	}
	newHead := node.Next
	node.Next = nil
	tail.Next = head
	return newHead
}
