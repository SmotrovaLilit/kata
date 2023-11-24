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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left > right {
		panic("left must be less than right")
	}
	if left == right {
		return head
	}
	node := head
	var prev, beforeFirstSub, firstSub *ListNode
	i := 1
	for node != nil {
		next := node.Next
		if i >= left && i <= right {
			node.Next = prev
			prev = node
			if i == left {
				firstSub = node
			}
			if i == right {
				if beforeFirstSub != nil { // subsequence is not at the beginning
					beforeFirstSub.Next.Next = next
					beforeFirstSub.Next = node
					break
				}
				if next == nil { //subsequence is the whole list
					head = node
					break
				}
				// subsequence is at the beginning of the list and does not include the whole element
				head = node
				firstSub.Next = next
				break
			}
		}
		if i == left-1 {
			beforeFirstSub = node
		}
		node = next
		i++
	}
	return head
}
