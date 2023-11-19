package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{nums: []int{0, 1, 0, 3, 2, 3}},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{nums: []int{7, 7, 7, 7, 7, 7, 7}},
			want: 1,
		},
		{
			name: "Example 4",
			args: args{nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLIS(tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
