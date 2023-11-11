package des2

import (
	. "leetcode/leetcode/pkg/tree"
)

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	helper(root, 0, &res)
	return res
}

func helper(root *TreeNode, level int, res *[]int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, root.Val)
	}
	if root.Right != nil {
		helper(root.Right, level+1, res)
	}
	if root.Left != nil {
		helper(root.Left, level+1, res)
	}
}
