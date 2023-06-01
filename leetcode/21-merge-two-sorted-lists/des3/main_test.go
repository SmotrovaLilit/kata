package des1

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				list1: NewListNode(1, 2, 4),
				list2: NewListNode(1, 3, 4),
			},
			want: NewListNode(1, 1, 2, 3, 4, 4),
		},
		{
			name: "case2",
			args: args{
				list1: NewListNode(),
				list2: NewListNode(),
			},
			want: NewListNode(),
		},
		{
			name: "case3",
			args: args{
				list1: NewListNode(),
				list2: NewListNode(0),
			},
			want: NewListNode(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
