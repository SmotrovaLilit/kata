package des1

import (
	"github.com/stretchr/testify/require"
	. "leetcode/leetcode/pkg/helper"

	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "empty",
			args: args{head: nil},
			want: nil,
		},
		{
			name: "one",
			args: args{head: BuildList([]int{1})},
			want: BuildList([]int{1}),
		},
		{
			name: "two",
			args: args{head: BuildList([]int{1, 1})},
			want: nil,
		},
		{
			name: "Example 1:",
			args: args{head: BuildList([]int{1, 2, 3, 3, 4, 4, 5})},
			want: BuildList([]int{1, 2, 5}),
		},
		{
			name: "Example 2:",
			args: args{head: BuildList([]int{1, 1, 1, 2, 3})},
			want: BuildList([]int{2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.args.head)
			require.Equal(t, ListToSlice(tt.want), ListToSlice(got))
		})
	}
}
