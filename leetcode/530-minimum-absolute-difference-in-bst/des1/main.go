package des0

import (
	. "leetcode/leetcode/pkg/helper"
	"math"
	"sort"
)

func getMinimumDifference(root *TreeNode) int {
	values := make([]int, 0)
	dfs(root, &values)
	sort.Ints(values)
	minValue := math.MaxInt
	for i := 1; i < len(values); i++ {
		minValue = min(minValue, distance(values[i], values[i-1]))
	}
	return minValue
}

func dfs(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, values)
	*values = append(*values, root.Val)
	dfs(root.Right, values)
}

func distance(a, b int) int {
	dist := a - b
	if dist < 0 {
		dist *= -1
	}
	return dist
}
