package des

import "fmt"

type listNode struct {
	val  int32
	next *listNode
}
type stack struct {
	head *listNode
}

func (s *stack) Push(val int32) {
	node := &listNode{
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

func isValid(s string) bool {
	stMain := &stack{}
	for _, v := range s {
		if isClosedBracket(v) {
			if stMain.IsEmpty() {
				return false
			}
			if stMain.Pop() != getOpenBracketCode(v) {
				return false
			}
		} else {
			stMain.Push(v)
		}
	}

	return stMain.IsEmpty()
}

func getOpenBracketCode(val int32) int32 {
	switch val {
	case []rune("]")[0]:
		return []rune("[")[0]
	case []rune("}")[0]:
		return []rune("{")[0]
	case []rune(")")[0]:
		return []rune("(")[0]
	}
	panic(fmt.Sprintf("incorrect symbol %d", val))
}

func isClosedBracket(val int32) bool {
	if val == []rune("]")[0] {
		return true
	}
	if val == []rune("}")[0] {
		return true
	}
	if val == []rune(")")[0] {
		return true
	}
	return false
}
