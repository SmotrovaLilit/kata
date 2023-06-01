package des1_bad

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
	currentList1 := list1
	currentList2 := list2
	var result *ListNode
	var minNode *ListNode
	var resultCurrent *ListNode
	for {
		minNode = nil
		if currentList1 == nil && currentList2 == nil {
			break
		}
		if currentList1 != nil && currentList2 == nil {
			minNode = currentList1
			if resultCurrent != nil {
				resultCurrent.Next = minNode
			}
			if result == nil {
				result = minNode
			}
			break
		}

		if currentList1 == nil && currentList2 != nil {
			minNode = currentList2
			if resultCurrent != nil {
				resultCurrent.Next = minNode
			}
			if result == nil {
				result = minNode
			}
			break
		}

		if currentList2.Val < currentList1.Val {
			minNode = currentList2
			currentList2 = currentList2.Next
		} else {
			minNode = currentList1
			currentList1 = currentList1.Next
		}

		if resultCurrent != nil {
			resultCurrent.Next = minNode
		}
		if result == nil {
			result = minNode
		}

		resultCurrent = minNode
	}
	return result
}
