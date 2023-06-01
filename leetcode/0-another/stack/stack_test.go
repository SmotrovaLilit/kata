package stack

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
	"testing"
)

func Test_stack_Push(t *testing.T) {
	type fields struct {
		head *ListNode
	}
	type args struct {
		val int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "case1",
			fields: fields{
				head: nil,
			},
			args: args{
				val: 5,
			},
			want: "#5",
		},
		{
			name: "case2",
			fields: fields{
				head: &ListNode{
					val: 8,
				},
			},
			args: args{
				val: 5,
			},
			want: "#5#8",
		},
		{
			name: "case3",
			fields: fields{
				head: &ListNode{
					val: 8,
					next: &ListNode{
						val:  9,
						next: nil,
					},
				},
			},
			args: args{
				val: 5,
			},
			want: "#5#8#9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				head: tt.fields.head,
			}
			s.Push(tt.args.val)
			require.Equal(t, tt.want, listToString(s.head))
		})
	}
}

func listToString(node *ListNode) string {
	buf := strings.Builder{}
	for node != nil {
		buf.WriteString("#")
		buf.WriteString(strconv.Itoa(int(node.val)))
		node = node.next
	}
	return buf.String()
}

func Test_stack_Pop(t *testing.T) {
	type fields struct {
		head *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		{
			name: "case1",
			fields: fields{
				head: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				head: tt.fields.head,
			}
			got := s.Pop()
			require.Equal(t, tt.want, got)
		})
	}
}
