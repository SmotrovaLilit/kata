package des1

import (
	"fmt"
	"github.com/stretchr/testify/require"
	. "leetcode/leetcode/pkg/helper"
	"math"
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -10,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
					Right: nil,
				},
			},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1},
			},
			want: &TreeNode{
				Val: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedArrayToBST(tt.args.nums)
			require.True(t, reflect.DeepEqual(tt.want, got))
		})
	}
}

func Benchmark_sortedArrayToBST(b *testing.B) {
	for j := 1; j <= 20; j++ {
		n := int(math.Pow(2, float64(j)))
		nums := generateSortedArray(n)
		b.Run(fmt.Sprintf("n/%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sortedArrayToBST(nums)
			}
		})
	}
}

func generateSortedArray(n int) []int {
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, i)
	}
	return res
}
