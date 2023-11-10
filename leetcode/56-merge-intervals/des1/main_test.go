package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test-1",
			args: args{
				intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "test-2",
			args: args{
				intervals: [][]int{{1, 4}, {4, 5}},
			},
			want: [][]int{{1, 5}},
		},
		{
			name: "test-3",
			args: args{
				intervals: [][]int{{1, 4}, {0, 4}},
			},
			want: [][]int{{0, 4}},
		},
		{
			name: "test-4",
			args: args{
				intervals: [][]int{{1, 4}, {0, 2}, {3, 5}},
			},
			want: [][]int{{0, 5}},
		},
		{
			name: "test-5",
			args: args{
				intervals: [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			},
			want: [][]int{{1, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.args.intervals)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestMakeNonOverlappingIntervals(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test-1",
			args: args{
				a: []int{1, 3},
				b: []int{3, 6},
			},
			want: [][]int{{1, 6}},
		},
		{
			name: "test-2",
			args: args{
				a: []int{3, 6},
				b: []int{3, 8},
			},
			want: [][]int{{3, 8}},
		},
		{
			name: "test-2",
			args: args{
				a: []int{2, 8},
				b: []int{3, 8},
			},
			want: [][]int{{2, 8}},
		},
		{
			name: "test-4",
			args: args{
				a: []int{3, 8},
				b: []int{1, 3},
			},
			want: [][]int{{1, 8}},
		},
		{
			name: "test-5",
			args: args{
				a: []int{1, 3},
				b: []int{2, 8},
			},
			want: [][]int{{1, 8}},
		},
		{
			name: "test-6",
			args: args{
				a: []int{2, 8},
				b: []int{1, 3},
			},
			want: [][]int{{1, 8}},
		},
		{
			name: "test-7",
			args: args{
				a: []int{2, 3},
				b: []int{1, 8},
			},
			want: [][]int{{1, 8}},
		},
		{
			name: "test-8",
			args: args{
				a: []int{1, 8},
				b: []int{2, 3},
			},
			want: [][]int{{1, 8}},
		},
		{
			name: "test-9",
			args: args{
				a: []int{1, 8},
				b: []int{9, 12},
			},
			want: [][]int{{1, 8}, {9, 12}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeNonOverlappingIntervals(tt.args.a, tt.args.b)
			require.Equal(t, tt.want, got)
		})
	}
}
