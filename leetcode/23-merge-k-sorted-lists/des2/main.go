package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(values ...int) *ListNode {
	var head, current *ListNode
	for _, value := range values {
		node := &ListNode{
			Val:  value,
			Next: nil,
		}
		if head == nil {
			head = node
		} else {
			current.Next = node
		}
		current = node
	}
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }.
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	result := lists[0]
	for i := 1; i < len(lists); i++ {
		prev := &ListNode{}
		headOfTwo := prev
		for {
			if result == nil {
				prev.Next = lists[i]
				break
			}
			if lists[i] == nil {
				prev.Next = result
				break
			}
			if result.Val <= lists[i].Val {
				prev.Next = result
				result = result.Next
			} else {
				prev.Next = lists[i]
				lists[i] = lists[i].Next
			}
			prev = prev.Next
		}

		result = headOfTwo.Next
	}
	return result
}
