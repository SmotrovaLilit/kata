package des3

import (
	. "leetcode/leetcode/pkg/helper"
)

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := &queue{}
	q.push(root)
	var res []int
	for !q.isEmpty() {
		levelLength := q.len()
		for i := 0; i < levelLength; i++ {
			v := q.pop()
			if i == levelLength-1 { // if it is the last node in this level
				res = append(res, v.Val)
			}
			if v.Left != nil {
				q.push(v.Left)
			}
			if v.Right != nil {
				q.push(v.Right)
			}
		}

	}
	return res
}

type node struct {
	val  *TreeNode
	next *node
}

type queue struct {
	last *node
	head *node
	size int
}

func (s *queue) len() int {
	return s.size
}

func (s *queue) push(val *TreeNode) {
	node := &node{
		val: val,
	}
	s.size++
	if s.last == nil {
		s.last = node
		s.head = node
		return
	}
	s.last.next = node
	s.last = node

}

func (s *queue) pop() *TreeNode {
	if s.isEmpty() {
		panic("queue is empty")
	}
	s.size--
	val := s.head.val
	s.head = s.head.next
	if s.head == nil {
		s.last = nil
	}
	return val
}

func (s *queue) isEmpty() bool {
	return s.last == nil && s.head == nil
}
