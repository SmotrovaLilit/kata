package des1

import (
	. "leetcode/leetcode/pkg/tree"
) /**
 * Definition for a binary tree tree.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func averageOfLevels(root *TreeNode) []float64 {
	levels := make(map[int]*levelInformation)
	stack := stack{}
	stack.push(root, 0)
	for !stack.isEmpty() {
		tree, level := stack.pop()
		if _, ok := levels[level]; !ok {
			levels[level] = &levelInformation{
				count: 1,
				sum:   tree.Val,
			}
		} else {
			levels[level].count++
			levels[level].sum += tree.Val
		}
		if tree.Left != nil {
			stack.push(tree.Left, level+1)
		}
		if tree.Right != nil {
			stack.push(tree.Right, level+1)
		}
	}
	res := make([]float64, 0, len(levels))
	for _, level := range levels {
		res = append(res, float64(level.sum)/float64(level.count))
	}
	return res
}

type levelInformation struct {
	count int
	sum   int
}

type stackNode struct {
	tree  *TreeNode
	level int
	next  *stackNode
}

type stack struct {
	top *stackNode
}

func (s *stack) push(node *TreeNode, level int) {
	s.top = &stackNode{
		tree:  node,
		level: level,
		next:  s.top,
	}
}

func (s *stack) pop() (*TreeNode, int) {
	if s.isEmpty() {
		panic("stack is empty")
	}
	n := s.top
	s.top = n.next
	return n.tree, n.level
}
func (s *stack) isEmpty() bool {
	return s.top == nil
}
