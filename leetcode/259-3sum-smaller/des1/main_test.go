package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_threeSumSmaller(t *testing.T) {
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
				nums:   []int{-2, 0, 1, 3},
				target: 2,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: 0,
		},
		{
			name: "case3",
			args: args{
				nums:   []int{0},
				target: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSumSmaller(tt.args.nums, tt.args.target)
			require.Equal(t, tt.want, got)
		})
	}
}
