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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}
	node := head
	length := 0
	for node != nil {
		node = node.Next
		length++
	}
	node = head
	if n > length {
		panic("invalid n")
	}
	if n == length { // remove first element
		return head.Next
	}
	i := 0
	for node != nil {
		if i == length-n-1 {
			node.Next = node.Next.Next
			break
		}
		node = node.Next
		i++
	}
	return head
}
