package fib

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{n: 0},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "Example 4",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "Example 5",
			args: args{n: 10},
			want: 89,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fib(tt.args.n)
			require.Equal(t, tt.want, got)
		})
	}
}
