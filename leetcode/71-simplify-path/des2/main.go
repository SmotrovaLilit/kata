package des1

import "strings"

func simplifyPath(path string) string {
	components := strings.Split(path, "/")
	st := &dequeue{}
	for i := 0; i < len(components); i++ {
		directory := components[i]
		if directory == "." || directory == "" {
			continue
		}
		if directory == ".." {
			if !st.isEmpty() {
				st.popBack()
			}
			continue
		}
		st.push(directory)
	}
	bi := strings.Builder{}
	for !st.isEmpty() {
		bi.WriteString("/")
		bi.WriteString(st.popFront())
	}
	res := bi.String()
	if res == "" {
		return "/"
	}
	return res
}

type node struct {
	value    string
	previous *node
	next     *node
}

type dequeue struct {
	tail *node
	head *node
}

func (s *dequeue) isEmpty() bool {
	return s.tail == nil && s.head == nil
}

func (s *dequeue) push(value string) {
	if s.isEmpty() {
		s.tail = &node{
			value: value,
		}
		s.head = s.tail
		return
	}
	s.tail = &node{
		value:    value,
		previous: s.tail,
	}
	s.tail.previous.next = s.tail
}

func (s *dequeue) popBack() string {
	if s.isEmpty() {
		panic("dequeue is empty")
	}
	value := s.tail.value
	s.tail = s.tail.previous
	if s.tail == nil {
		s.head = nil
	} else {
		s.tail.next = nil
	}
	return value
}

func (s *dequeue) popFront() string {
	if s.isEmpty() {
		panic("dequeue is empty")
	}
	value := s.head.value
	s.head = s.head.next
	if s.head == nil {
		s.tail = nil
	} else {
		s.head.previous = nil
	}
	return value
}
