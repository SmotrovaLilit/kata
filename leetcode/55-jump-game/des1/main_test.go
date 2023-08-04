package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				nums: []int{3, 2, 1, 0, 4},
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				nums: []int{2, 5, 0, 0},
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				nums: []int{2, 5, 1, 0, 1},
			},
			want: true,
		},
		{
			name: "case5",
			args: args{
				nums: []int{0},
			},
			want: true,
		},
		{
			name: "case6",
			args: args{
				nums: []int{1, 3, 0, 0, 4},
			},
			want: true,
		},
		{
			name: "case7",
			args: args{
				nums: []int{1, 2, 0, 0, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canJump(tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
