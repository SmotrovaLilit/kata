package des1

import (
	"fmt"
	"github.com/stretchr/testify/require"
	. "leetcode/leetcode/pkg/tree"
	"math"
	"reflect"
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "t1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: []float64{3, 14.5, 11},
		},
		{
			name: "t2",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
					Right: &TreeNode{
						Val: 20,
					},
				},
			},
			want: []float64{3, 14.5, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := averageOfLevels(tt.args.root)
			require.True(t, reflect.DeepEqual(got, tt.want))
		})
	}
}

func Benchmark_averageOfLevels(b *testing.B) {
	for j := 1; j <= 20; j++ {
		nodeCount := int(math.Pow(float64(2), float64(j)))
		tree := GenerateTree(nodeCount)
		b.Run(fmt.Sprintf("nodes/%d", nodeCount), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				averageOfLevels(tree)
			}
		})
	}
}
