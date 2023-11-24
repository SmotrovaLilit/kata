package des1

import (
	"github.com/stretchr/testify/require"
	. "leetcode/leetcode/pkg/helper"

	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "empty",
			args: args{head: nil, k: 0},
			want: nil,
		},
		{
			name: "Example 1",
			args: args{head: BuildList([]int{1, 2, 3, 4, 5}), k: 2},
			want: BuildList([]int{4, 5, 1, 2, 3}),
		},
		{
			name: "Example 2",
			args: args{head: BuildList([]int{0, 1, 2}), k: 4},
			want: BuildList([]int{2, 0, 1}),
		},
		{
			name: "Example 3",
			args: args{head: BuildList([]int{1, 2}), k: 1},
			want: BuildList([]int{2, 1}),
		},
		{
			name: "Example 3:",
			args: args{head: BuildList([]int{1, 2}), k: 0},
			want: BuildList([]int{1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateRight(tt.args.head, tt.args.k)
			require.Equal(t, ListToSlice(tt.want), ListToSlice(got))
		})
	}
}
