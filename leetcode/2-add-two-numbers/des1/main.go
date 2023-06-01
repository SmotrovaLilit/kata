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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := l1
	var sum, carry int
	for {
		sum = carry + l1.Val + l2.Val
		l1.Val = sum % 10
		carry = sum / 10
		if l2.Next == nil || l1.Next == nil {
			if l1.Next == nil && l2.Next != nil {
				l1.Next = l2.Next
			}
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if carry == 0 {
		return result
	}

	for l1.Next != nil {
		sum = carry + l1.Next.Val
		l1.Next.Val = sum % 10
		carry = sum / 10
		l1 = l1.Next
	}
	if carry != 0 {
		l1.Next = &ListNode{
			Val: carry,
		}
	}

	return result
}
