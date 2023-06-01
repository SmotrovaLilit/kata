package main

func main() {

}

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
func mergeKLists(lists []*ListNode) *ListNode {
	resultHead := &ListNode{}
	resultCurrent := resultHead
	var minListNumber int
	for {
		minListNumber = -1
		min := 10_001
		for listNumber, list := range lists {
			if list == nil {
				continue
			}
			if list.Val <= min {
				min = list.Val
				minListNumber = listNumber
			}
		}
		if minListNumber == -1 {
			break
		}
		resultCurrent.Next = lists[minListNumber]
		lists[minListNumber] = lists[minListNumber].Next
		resultCurrent = resultCurrent.Next
	}
	return resultHead.Next
}
