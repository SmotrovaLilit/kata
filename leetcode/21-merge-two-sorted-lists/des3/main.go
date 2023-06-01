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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	prev := &ListNode{}
	resultHead := prev
	for {
		if list1 == nil {
			prev.Next = list2
			break
		}
		if list2 == nil {
			prev.Next = list1
			break
		}
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}

	return resultHead.Next
}
