package des1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	st := stack{}
	st.push(root, targetSum)
	for !st.isEmpty() {
		n, sum := st.pop()
		sum -= n.Val
		if sum == 0 && n.Left == nil && n.Right == nil {
			return true
		}
		if n.Left != nil {
			st.push(n.Left, sum)
		}
		if n.Right != nil {
			st.push(n.Right, sum)
		}
	}
	return false
}

type nodeStack struct {
	node *TreeNode
	sum  int
	next *nodeStack
}

type stack struct {
	top *nodeStack
}

func (s *stack) push(node *TreeNode, sum int) {
	s.top = &nodeStack{
		node: node,
		sum:  sum,
		next: s.top,
	}
}

func (s *stack) pop() (*TreeNode, int) {
	if s.isEmpty() {
		panic("stack is empty")
	}
	n := s.top
	s.top = n.next
	return n.node, n.sum
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
