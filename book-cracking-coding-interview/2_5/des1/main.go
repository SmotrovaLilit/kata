package des1

func sum(list1, list2 LinkedList) LinkedList {
	n1 := list1.head
	n2 := list2.head
	resEnd := &Node{}
	resStart := resEnd

	c := 0
	for n1 != nil && n2 != nil {
		value := n1.Value + n2.Value + c
		c = value / 10
		value = value % 10
		resEnd.Next = &Node{
			Value: value,
		}

		resEnd = resEnd.Next
		n1 = n1.Next
		n2 = n2.Next
	}
	for n2 != nil {
		resEnd.Next = &Node{
			Value: n2.Value,
		}
		resEnd = resEnd.Next
		n2 = n2.Next
	}
	for n1 != nil {
		resEnd.Next = &Node{
			Value: n1.Value,
		}
		resEnd = resEnd.Next
		n1 = n1.Next
	}
	return LinkedList{head: resStart.Next}
}

func sumForward(list1, list2 LinkedList) LinkedList {
	reversed1 := list1.Reverse()
	reversed2 := list2.Reverse()
	s := sum(reversed1, reversed2)
	return s.Reverse()
}
