package des1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	count := 0
	st := stack{}
	st.push(root)
	for !st.isEmpty() {
		n := st.pop()
		if n == nil {
			continue
		}
		count++
		st.push(n.Left)
		st.push(n.Right)
	}
	return count
}

type stackNode struct {
	node *TreeNode
	next *stackNode
}

type stack struct {
	top *stackNode
}

func (s *stack) push(node *TreeNode) {
	s.top = &stackNode{
		node: node,
		next: s.top,
	}
}

func (s *stack) pop() *TreeNode {
	if s.isEmpty() {
		panic("stack is empty")
	}
	n := s.top
	s.top = n.next
	return n.node
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
