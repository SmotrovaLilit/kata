package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
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
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 0,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 3,
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				nums:   []int{11, 13, 15, 17},
				target: 1,
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				nums:   []int{11, 13, 15, 17},
				target: 11,
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				nums:   []int{1, 0, 1, 1, 1},
				target: 0,
			},
			want: true,
		},
		{
			name: "case5",
			args: args{
				nums:   []int{1, 1, 1, 0, 1},
				target: 0,
			},
			want: true,
		},
		{
			name: "case6",
			args: args{
				nums:   []int{1, 0, 1, 1, 1},
				target: 2,
			},
			want: false,
		},
		{
			name: "case7",
			args: args{
				nums:   []int{1, 0, 1, 1, 1},
				target: 1,
			},
			want: true,
		},
		{
			name: "case8",
			args: args{
				nums:   []int{2, 0, 2, 2, 2},
				target: 1,
			},
			want: false,
		},
		{
			name: "case9",
			args: args{
				nums:   []int{2, 2, 2, 2, 2},
				target: 1,
			},
			want: false,
		},
		{
			name: "case10",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 1,
			},
			want: true,
		},
		{
			name: "case11",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 2,
			},
			want: true,
		},
		{
			name: "case12",
			args: args{
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1},
				target: 2,
			},
			want: true,
		},
		{
			name: "case13-both in second, mid > target",
			args: args{
				nums:   []int{14, 15, -1, 0, 1, 2, 3, 4, 5},
				target: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.args.nums, tt.args.target)
			require.Equal(t, tt.want, got)
		})
	}
}
