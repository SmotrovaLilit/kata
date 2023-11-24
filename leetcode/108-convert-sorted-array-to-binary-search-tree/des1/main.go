package des1

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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	rootIndex := len(nums) / 2
	return &TreeNode{
		Val:   nums[rootIndex],
		Left:  sortedArrayToBST(nums[:rootIndex]),
		Right: sortedArrayToBST(nums[rootIndex+1:]),
	}
}
