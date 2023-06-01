package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "case1_1",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "case1_2",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 7,
			},
			want: true,
		},
		{
			name: "case1_3",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 20,
			},
			want: true,
		},
		{
			name: "case1_4",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 16,
			},
			want: true,
		},
		{
			name: "case1_5",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 1,
			},
			want: true,
		},
		{
			name: "case1_6",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 23,
			},
			want: true,
		},
		{
			name: "case1_7",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 60,
			},
			want: true,
		},
		{
			name: "case1_8",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 34,
			},
			want: true,
		},
		{
			name: "case1_9",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7, 8},
					{10, 11, 16, 20, 21},
					{23, 30, 34, 60, 61},
				},
				target: 34,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 13,
			},
			want: false,
		},
		{
			name: "case2_1",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
				},
				target: 13,
			},
			want: false,
		},
		{
			name: "case2_2",
			args: args{
				matrix: [][]int{},
				target: 13,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchMatrix(tt.args.matrix, tt.args.target)
			require.Equal(t, tt.want, got)
		})
	}
}
