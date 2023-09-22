package des1

/**
 * Definition for a binary tree stackNode.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	st := newStack()
	st.push(p, q)
	for !st.isEmpty() {
		p, q = st.pop()
		if (p == nil) && (q == nil) {
			continue
		}
		if (p == nil) || (q == nil) {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		st.push(p.Left, q.Left)
		st.push(p.Right, q.Right)
	}
	return true
}

type stackNode struct {
	p    *TreeNode
	q    *TreeNode
	next *stackNode
}

type stack struct {
	top *stackNode
}

func (s *stack) push(p *TreeNode, q *TreeNode) {
	s.top = &stackNode{
		p:    p,
		q:    q,
		next: s.top,
	}
}

func (s *stack) pop() (*TreeNode, *TreeNode) {
	if s.isEmpty() {
		panic("stack is empty")
	}
	n := s.top
	s.top = n.next
	return n.p, n.q
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

func newStack() stack {
	return stack{}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
