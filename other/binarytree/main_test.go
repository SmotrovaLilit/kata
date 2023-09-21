package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"golang.org/x/tour/tree"
	"testing"
)

func TestWalk(t *testing.T) {
	tre := tree.New(1)
	ch := make(chan int)
	go func() {
		Walk(tre, ch)
		close(ch)
	}()
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := make([]int, 0, len(expected))
	for v := range ch {
		fmt.Println(v)
		got = append(got, v)
	}

	require.ElementsMatch(t, expected, got)
}

func TestSame(t *testing.T) {
	type args struct {
		t1 *tree.Tree
		t2 *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same",
			args: args{
				t1: tree.New(1),
				t2: tree.New(1),
			},
			want: true,
		},
		{
			name: "not same",
			args: args{
				t1: tree.New(1),
				t2: tree.New(2),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Same(tt.args.t1, tt.args.t2)
			require.Equal(t, tt.want, got)
		})
	}
}
