package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
		{
			name: "case3",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "case4",
			args: args{
				nums: []int{4, 1, 2, 1, 2, 4},
			},
			want: 0,
		},
		{
			name: "case5",
			args: args{
				nums: []int{0, 1, 2, 1, 2},
			},
			want: 0,
		},
		{
			name: "case6",
			args: args{
				nums: []int{-2, -1, 2, -1, 2},
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
