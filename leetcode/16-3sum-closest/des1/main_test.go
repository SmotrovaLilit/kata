package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				nums:   []int{0, 0, 0},
				target: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSumClosest(tt.args.nums, tt.args.target)
			require.Equal(t, tt.want, got)
		})
	}
}
