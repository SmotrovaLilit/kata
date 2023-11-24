package des1

import (
	"github.com/stretchr/testify/require"
	"testing"

	. "leetcode/leetcode/pkg/helper"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "empty",
			args: args{head: nil, n: 0},
			want: nil,
		},
		{
			name: "one",
			args: args{head: BuildList([]int{1}), n: 1},
			want: nil,
		},
		{
			name: "two",
			args: args{head: BuildList([]int{1, 2}), n: 1},
			want: BuildList([]int{1}),
		},
		{
			name: "remove first element",
			args: args{head: BuildList([]int{1, 2}), n: 2},
			want: BuildList([]int{2}),
		},
		{
			name: "three",
			args: args{head: BuildList([]int{1, 2, 3}), n: 2},
			want: BuildList([]int{1, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.args.head, tt.args.n)
			require.Equal(t, tt.want, got)
		})
	}
}
