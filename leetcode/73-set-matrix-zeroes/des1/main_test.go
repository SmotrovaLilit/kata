package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	var tests = []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				matrix: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{-1, 1, 1},
				},
			},
			want: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{-1, 0, 1},
			},
		},
		{
			name: "case2",
			args: args{
				matrix: [][]int{
					{0, 1, 2, 0},
					{3, 4, 5, 2},
					{1, 3, 1, 5},
				},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			name: "case3",
			args: args{
				matrix: [][]int{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 0},
				},
			},
			want: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 0},
			},
		},
		{
			name: "case4",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4},
					{5, 0, 7, 8},
					{0, 10, 11, 12},
					{13, 14, 15, 0},
				},
			},
			want: [][]int{
				{0, 0, 3, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			require.Equal(t, tt.want, tt.args.matrix)
		})
	}
}
