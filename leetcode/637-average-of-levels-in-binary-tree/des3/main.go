package des3

import (
	. "leetcode/leetcode/pkg/helper"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	mainStack := stack{}
	mainStack.push(root)
	res := make([]float64, 0)
	for !mainStack.isEmpty() {
		sum, count := 0, 0
		tempStack := stack{}
		for !mainStack.isEmpty() {
			tree := mainStack.pop()
			sum += tree.Val
			count++
			if tree.Left != nil {
				tempStack.push(tree.Left)
			}
			if tree.Right != nil {
				tempStack.push(tree.Right)
			}
		}
		mainStack = tempStack
		res = append(res, float64(sum)/float64(count))
	}
	return res
}

type stackNode struct {
	tree *TreeNode
	next *stackNode
}

type stack struct {
	top *stackNode
}

func (s *stack) push(node *TreeNode) {
	s.top = &stackNode{
		tree: node,
		next: s.top,
	}
}

func (s *stack) pop() *TreeNode {
	if s.isEmpty() {
		panic("stack is empty")
	}
	n := s.top
	s.top = n.next
	return n.tree
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

type stack2 struct {
	data  []*TreeNode
	index int
}

func newStack2() *stack2 {
	return &stack2{
		data:  make([]*TreeNode, 0),
		index: -1,
	}
}

func (s *stack2) push(node *TreeNode) {
	s.index++
	s.data = append(s.data, node)
}

func (s *stack2) pop() *TreeNode {
	if s.isEmpty() {
		panic("stack is empty")
	}
	v := s.data[s.index]
	s.index--
	return v
}

func (s *stack2) isEmpty() bool {
	return s.index < 0
}
