package des1

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				l1: NewListNode(2, 4, 3),
				l2: NewListNode(5, 6, 4),
			},
			want: NewListNode(7, 0, 8),
		},
		{
			name: "case1_1",
			args: args{
				l1: NewListNode(2, 4, 3),
				l2: NewListNode(2, 6, 4),
			},
			want: NewListNode(4, 0, 8),
		},
		{
			name: "case2",
			args: args{
				l1: NewListNode(0),
				l2: NewListNode(0),
			},
			want: NewListNode(0),
		},
		{
			name: "case3",
			args: args{
				l1: NewListNode(9, 9, 9, 9),
				l2: NewListNode(9, 9, 9, 9),
			},
			want: NewListNode(8, 9, 9, 9, 1),
		},
		{
			name: "case4",
			args: args{
				l1: NewListNode(9, 9, 9, 9),
				l2: NewListNode(9, 9, 9),
			},
			want: NewListNode(8, 9, 9, 0, 1),
		},
		{
			name: "case5",
			args: args{
				l1: NewListNode(9, 9, 9),
				l2: NewListNode(9, 9, 9, 9),
			},
			want: NewListNode(8, 9, 9, 0, 1),
		},
		{
			name: "case6",
			args: args{
				l1: NewListNode(9, 9, 9, 9, 9, 9, 9),
				l2: NewListNode(9, 9, 9, 9),
			},
			want: NewListNode(8, 9, 9, 9, 0, 0, 0, 1),
		},
		{
			name: "case7",
			args: args{
				l1: NewListNode(9, 9, 9, 9, 9, 9, 9),
				l2: NewListNode(),
			},
			want: NewListNode(9, 9, 9, 9, 9, 9, 9),
		},
		{
			name: "case8",
			args: args{
				l1: NewListNode(),
				l2: NewListNode(9, 9, 9, 9, 9, 9, 9),
			},
			want: NewListNode(9, 9, 9, 9, 9, 9, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
