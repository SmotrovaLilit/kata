package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_jump(t *testing.T) {
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
				nums: []int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				nums: []int{3, 2, 1, 0, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jump(tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
