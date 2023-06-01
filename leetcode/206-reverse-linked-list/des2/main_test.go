package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				head: NewListNode(1, 2, 3, 4, 5),
			},
			want: NewListNode(5, 4, 3, 2, 1),
		},
		{
			name: "case2",
			args: args{
				head: NewListNode(1, 2),
			},
			want: NewListNode(2, 1),
		},
		{
			name: "case3",
			args: args{
				head: NewListNode(),
			},
			want: NewListNode(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.args.head)
			require.Equal(t, tt.want, got)
		})
	}
}
