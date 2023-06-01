package des1

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
 * }
 */
func reverseList(head *ListNode) *ListNode {
	current := head
	var prev *ListNode
	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	return prev
}
