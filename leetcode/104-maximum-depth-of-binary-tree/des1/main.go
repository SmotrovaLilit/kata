package des1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	stack := &treeNodeDepthStack{}
	stack.Push(root, 1)
	var node *TreeNode
	var maxDepth, depth int
	for !stack.Empty() {
		node, depth = stack.Pop()
		if depth > maxDepth {
			maxDepth = depth
		}
		if node.Right != nil {
			stack.Push(node.Right, depth+1)
		}
		if node.Left != nil {
			stack.Push(node.Left, depth+1)
		}
	}
	return maxDepth
}

type treeDepthNode struct {
	depth int
	node  *TreeNode
	next  *treeDepthNode
}

type treeNodeDepthStack struct {
	top *treeDepthNode
}

func (s *treeNodeDepthStack) Push(node *TreeNode, depth int) {
	s.top = &treeDepthNode{
		depth: depth,
		node:  node,
		next:  s.top,
	}
}

func (s *treeNodeDepthStack) Pop() (node *TreeNode, depth int) {
	if s.Empty() {
		panic("stack is empty")
	}
	node = s.top.node
	depth = s.top.depth
	s.top = s.top.next
	return node, depth
}

func (s *treeNodeDepthStack) Empty() bool {
	return s.top == nil
}
