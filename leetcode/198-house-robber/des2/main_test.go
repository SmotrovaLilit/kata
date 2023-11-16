package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rob(t *testing.T) {
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
			args: args{nums: []int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{nums: []int{2, 7, 9, 3, 1}},
			want: 12,
		},
		{
			name: "Example 3",
			args: args{nums: []int{1, 1}},
			want: 1,
		},
		{
			name: "Example 4",
			args: args{nums: []int{2, 3, 2}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
