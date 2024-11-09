package des1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }.
 */
func invertTree(root *TreeNode) *TreeNode {
	stack := stack{}
	stack.push(root)
	for !stack.isEmpty() {
		n := stack.pop()
		if n == nil {
			continue
		}
		n.Left, n.Right = n.Right, n.Left
		stack.push(n.Left)
		stack.push(n.Right)
	}
	return root
}

type nodeStack struct {
	val  *TreeNode
	next *nodeStack
}

type stack struct {
	top *nodeStack
}

func (s *stack) push(val *TreeNode) {
	s.top = &nodeStack{
		val:  val,
		next: s.top,
	}
}

func (s *stack) pop() *TreeNode {
	if s.isEmpty() {
		panic("stack is empty")
	}
	n := s.top
	s.top = n.next
	return n.val
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
