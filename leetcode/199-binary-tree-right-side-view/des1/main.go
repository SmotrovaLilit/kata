package des1

import (
	. "leetcode/leetcode/pkg/tree"
)

func rightSideView(root *TreeNode) []int {
	st := &stack{}
	st.push(root, 0)
	var res []int
	for !st.isEmpty() {
		v, level := st.pop()
		if v == nil {
			continue
		}
		if len(res) == level {
			res = append(res, v.Val)
		}
		st.push(v.Left, level+1)
		st.push(v.Right, level+1)
	}
	return res
}

type node struct {
	val   *TreeNode
	level int
	next  *node
}

type stack struct {
	top *node
}

func (s *stack) push(val *TreeNode, level int) {
	s.top = &node{
		val:   val,
		level: level,
		next:  s.top,
	}
}

func (s *stack) pop() (*TreeNode, int) {
	if s.isEmpty() {
		panic("stack is empty")
	}
	val := s.top.val
	level := s.top.level
	s.top = s.top.next
	return val, level
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}
