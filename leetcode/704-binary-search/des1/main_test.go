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
		want int
	}{
		{
			name: "case1",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "case2",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12, 15},
				target: 9,
			},
			want: 4,
		},
		{
			name: "case3",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.args.nums, tt.args.target)
			require.Equal(t, tt.want, got)
		})
	}
}
