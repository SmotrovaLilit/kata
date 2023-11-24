package des1

import (
	"github.com/stretchr/testify/require"
	. "leetcode/leetcode/pkg/helper"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				head:  BuildList([]int{1, 2, 3, 4, 5}),
				left:  2,
				right: 4,
			},
			want: BuildList([]int{1, 4, 3, 2, 5}),
		},
		{
			name: "Example 2",
			args: args{
				head:  BuildList([]int{5}),
				left:  1,
				right: 1,
			},
			want: BuildList([]int{5}),
		},
		{
			name: "Example 3",
			args: args{
				head:  BuildList([]int{3, 5}),
				left:  1,
				right: 2,
			},
			want: BuildList([]int{5, 3}),
		},
		{
			name: "Example 4",
			args: args{
				head:  BuildList([]int{1, 2, 3, 4, 5}),
				left:  1,
				right: 5,
			},
			want: BuildList([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "Example 5",
			args: args{
				head:  BuildList([]int{1, 2, 3, 4, 5, 6}),
				left:  1,
				right: 5,
			},
			want: BuildList([]int{5, 4, 3, 2, 1, 6}),
		},
		{
			name: "Example 6",
			args: args{
				head:  BuildList([]int{1, 2, 3, 4, 5}),
				left:  2,
				right: 5,
			},
			want: BuildList([]int{1, 5, 4, 3, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween(tt.args.head, tt.args.left, tt.args.right)
			require.Equal(t, ListToSlice(tt.want), ListToSlice(got))
		})
	}
}
