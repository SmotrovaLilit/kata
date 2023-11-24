package des2

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

type levelInformation struct {
	count int
	sum   int
}
type levels struct {
	data []levelInformation
}

func (l *levels) averageHelper(t *TreeNode, level int) {
	if t == nil {
		return
	}
	if len(l.data) <= level {
		l.data = append(l.data, levelInformation{
			count: 1,
			sum:   t.Val,
		})
	} else {
		l.data[level].sum += t.Val
		l.data[level].count++
	}

	l.averageHelper(t.Left, level+1)
	l.averageHelper(t.Right, level+1)
}

func averageOfLevels(root *TreeNode) []float64 {
	l := levels{}
	l.averageHelper(root, 0)
	res := make([]float64, len(l.data))
	for i := 0; i < len(res); i++ {
		res[i] = float64(l.data[i].sum) / float64(l.data[i].count)
	}
	return res
}
