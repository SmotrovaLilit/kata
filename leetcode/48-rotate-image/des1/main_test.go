package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			name: "Example 1",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "Example 2",
			input: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			want: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{
			name: "Example 3",
			input: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			want: [][]int{
				{21, 16, 11, 6, 1},
				{22, 17, 12, 7, 2},
				{23, 18, 13, 8, 3},
				{24, 19, 14, 9, 4},
				{25, 20, 15, 10, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.input)
			require.Equal(t, tt.want, tt.input)
		})
	}
}

func Test_nextPointToRotate(t *testing.T) {
	type args struct {
		row        int
		col        int
		firstIndex int
		lastIndex  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "corner case 1",
			args: args{
				row:        1,
				col:        1,
				firstIndex: 1,
				lastIndex:  7,
			},
			want: []int{1, 7},
		},
		{
			name: "middle case 1",
			args: args{
				row:        1,
				col:        2,
				firstIndex: 1,
				lastIndex:  7,
			},
			want: []int{2, 7},
		},
		{
			name: "middle case 2",
			args: args{
				row:        1,
				col:        3,
				firstIndex: 1,
				lastIndex:  7,
			},
			want: []int{3, 7},
		},
		{
			name: "corner case 2",
			args: args{
				row:        1,
				col:        7,
				firstIndex: 1,
				lastIndex:  7,
			},
			want: []int{7, 7},
		},
		{
			name: "mod case 3",
			args: args{
				row:        4,
				col:        7,
				firstIndex: 1,
				lastIndex:  7,
			},
			want: []int{7, 4},
		},
		{
			name: "corner case 3",
			args: args{
				row:        7,
				col:        7,
				firstIndex: 1,
				lastIndex:  7,
			},
			want: []int{7, 1},
		},
		{
			name: "mid case 5",
			args: args{
				row:        7,
				col:        4,
				firstIndex: 1,
				lastIndex:  7,
			},
			want: []int{4, 1},
		},
		{
			name: "corner case 4",
			args: args{
				row:        7,
				col:        1,
				firstIndex: 1,
				lastIndex:  7,
			},
			want: []int{1, 1},
		},
		{
			name: "mid case 6",
			args: args{
				row:        3,
				col:        1,
				firstIndex: 1,
				lastIndex:  7,
			},
			want: []int{1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := nextPointToRotate(tt.args.row, tt.args.col, tt.args.firstIndex, tt.args.lastIndex)
			require.Equal(t, tt.want, []int{got, got1})
		})
	}
}
