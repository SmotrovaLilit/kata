package des1

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) Add(value int) {
	end := Node{
		Value: value,
	}
	if l.IsEmpty() {
		l.head = &end
		return
	}
	n := l.head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = &end
}

func (l *LinkedList) Reverse() LinkedList {
	var newHead *Node
	n := l.head
	for n != nil {
		newNode := &Node{Value: n.Value}
		newNode.Next = newHead
		newHead = newNode
		n = n.Next
	}

	return LinkedList{head: newHead}
}

func (l *LinkedList) Remove(value int) {
	n := l.head
	// Remove first elements
	for n != nil && n.Value == value {
		n = n.Next
	}
	l.head = n
	for n != nil && n.Next != nil {
		if n.Next.Value == value {
			n.Next = n.Next.Next
			continue
		}
		n = n.Next
	}
}

func CreateFromSlice(values []int) LinkedList {
	if len(values) == 0 {
		return LinkedList{}
	}
	list := LinkedList{
		head: &Node{Value: values[0]},
	}
	end := list.head
	for i := 1; i < len(values); i++ {
		end.Next = &Node{Value: values[i]}
		end = end.Next
	}
	return list
}

func (l *LinkedList) ToSlice() []int {
	n := l.head
	result := make([]int, 0)
	for n != nil {
		result = append(result, n.Value)
		n = n.Next
	}
	return result
}
