package biginteger

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_sum(t *testing.T) {
	type args struct {
		num1 []int
		num2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				num1: []int{1, 2, 3},
				num2: []int{4, 5, 6},
			},
			want: []int{5, 7, 9},
		},
		{
			name: "case2",
			args: args{
				num1: []int{1, 2, 3},
				num2: []int{4, 5, 6, 7},
			},
			want: []int{4, 6, 9, 0},
		},
		{
			name: "big",
			args: args{
				num1: []int{9, 9, 9},
				num2: []int{9, 9, 9, 9},
			},
			want: []int{
				1, 0, 9, 9, 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sum(tt.args.num1, tt.args.num2)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_fff(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{
					0, 2, 3, 1, 2, 3, 3, 3, 3, 2, 8, 2, 4, 5, 6, 7, 7,
				},
				k: 3,
			},
			want: []int{3, 2, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fff(tt.args.nums, tt.args.k)
			require.Equal(t, tt.want, got)
		})
	}
}
