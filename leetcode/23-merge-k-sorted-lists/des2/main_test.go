package main

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				lists: []*ListNode{
					NewListNode(1, 4, 5),
					NewListNode(1, 3, 4),
					NewListNode(2, 6),
				},
			},
			want: NewListNode(1, 1, 2, 3, 4, 4, 5, 6),
		},
		{
			name: "case1_1",
			args: args{
				lists: []*ListNode{
					NewListNode(1, 3, 4, 6, 8, 9, 12),
					NewListNode(1, 2, 5, 7, 11, 21, 24),
					NewListNode(-4, 0, 4, 7, 10, 14, 22, 29),
				},
			},
			want: NewListNode(-4, 0, 1, 1, 2, 3, 4, 4, 5, 6, 7, 7, 8, 9, 10, 11, 12, 14, 21, 22, 24, 29),
		},
		{
			name: "case2",
			args: args{
				lists: []*ListNode{},
			},
			want: NewListNode(),
		},
		{
			name: "case3",
			args: args{
				lists: []*ListNode{
					NewListNode(),
				},
			},
			want: NewListNode(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMergeKLists(b *testing.B) {
	//b.ReportAllocs()
	lists := []*ListNode{
		NewListNode(1, 3, 4, 6, 8, 9, 12),
		NewListNode(1, 2, 5, 7, 11, 21, 24),
		NewListNode(-4, 0, 4, 7, 10, 14, 22, 29),
	}
	for i := 0; i < b.N; i++ {
		mergeKLists(lists)
	}
}
