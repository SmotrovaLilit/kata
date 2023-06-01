package main

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				head: NewCycledListNode(1, 3, 2, 0, -4),
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				head: NewCycledListNode(0, 1, 2),
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				head: NewCycledListNode(-1, 1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func NewCycledListNode(pos int, values ...int) *ListNode {
	head := NewListNode(values...)
	if pos < 0 || pos >= len(values) {
		return head
	}
	var tail *ListNode
	for p := head; p != nil; p = p.Next {
		tail = p
	}
	posNode := head
	for i := 0; i < pos; i++ {
		posNode = posNode.Next
	}
	tail.Next = posNode
	return head
}

func NewListNode(values ...int) *ListNode {
	var head, current *ListNode
	for _, value := range values {
		node := &ListNode{
			Val:  value,
			Next: nil,
		}
		if head == nil {
			head = node
		} else {
			current.Next = node
		}
		current = node
	}
	return head
}
