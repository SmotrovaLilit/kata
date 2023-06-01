package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case4",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{1, 2},
		},
		{
			name: "case5",
			args: args{
				nums:   []int{2, 3, 4},
				target: 6,
			},
			want: []int{1, 3},
		},
		{
			name: "case6",
			args: args{
				nums:   []int{-1, 0},
				target: -1,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			require.Equal(t, tt.want, got)
		})
	}
}
