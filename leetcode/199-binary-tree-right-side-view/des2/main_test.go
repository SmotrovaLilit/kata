package des2

import (
	"github.com/stretchr/testify/require"
	. "leetcode/leetcode/pkg/helper"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
						},
					},
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			want: []int{1, 3, 4},
		},
		{
			name: "case2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: []int{1, 3, 4},
		},
		{
			name: "case3",
			args: args{
				root: nil,
			},
			want: []int{},
		},
		{
			name: "case4",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			want: []int{1, 2},
		},
		{
			name: "case5",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: []int{1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rightSideView(tt.args.root)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
