package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			want: 6,
		},
		{
			name: "t2",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: 1,
		},
		{
			name: "t3",
			args: args{
				root: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countNodes(tt.args.root)
			require.Equal(t, tt.want, got)
		})
	}
}
