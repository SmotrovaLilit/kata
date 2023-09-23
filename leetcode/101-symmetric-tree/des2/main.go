package des1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	st := stack{}
	st.push(root.Left, root.Right)
	for !st.isEmpty() {
		n1, n2 := st.pop()
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		st.push(n1.Left, n2.Right)
		st.push(n1.Right, n2.Left)
	}
	return true
}

type nodeStack struct {
	val1 *TreeNode
	val2 *TreeNode
	next *nodeStack
}

type stack struct {
	top *nodeStack
}

func (s *stack) push(val1 *TreeNode, val2 *TreeNode) {
	s.top = &nodeStack{
		val1: val1,
		val2: val2,
		next: s.top,
	}
}

func (s *stack) pop() (*TreeNode, *TreeNode) {
	if s.isEmpty() {
		panic("stack is empty")
	}
	n := s.top
	s.top = n.next
	return n.val1, n.val2
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
