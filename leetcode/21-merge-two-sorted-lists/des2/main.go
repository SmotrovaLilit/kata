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
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
