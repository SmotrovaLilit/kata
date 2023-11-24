package helper

func BuildList(values []int) *ListNode {
	var head, node *ListNode
	if len(values) == 0 {
		return nil
	}
	head = &ListNode{Val: values[0]}
	node = head
	for i := 1; i < len(values); i++ {
		node.Next = &ListNode{Val: values[i]}
		node = node.Next
	}
	return head
}

func ListToSlice(node *ListNode) []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}
