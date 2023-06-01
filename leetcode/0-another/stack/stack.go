package stack

type Stack interface {
	Push(int32)
	Pop() int32
	IsEmpty() bool
}

type ListNode struct {
	val  int32
	next *ListNode
}
type stack struct {
	head *ListNode
}

func (s *stack) Push(val int32) {
	node := &ListNode{
		val:  val,
		next: s.head,
	}
	if s.head == nil {
		s.head = node
	} else {
		s.head = node
	}

}

func (s *stack) Pop() int32 {
	if s.IsEmpty() {
		panic("empty stack")
	}
	res := s.head.val
	s.head = s.head.next
	return res
}

func (s *stack) IsEmpty() bool {
	return s.head == nil
}
