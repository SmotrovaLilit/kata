package des1

func LoopDetection(l LinkedList) int {
	slow := l.head
	fast := l.head
	found := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			found = true
			break
		}
	}
	if !found {
		return -1
	}
	slow = l.head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Value
}
