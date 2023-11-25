package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[100,4,200,1,3,2]",
			args: args{nums: []int{100, 4, 200, 1, 3, 2}},
			want: 4,
		},
		{
			name: "[0,3,7,2,5,8,4,6,0,1]",
			args: args{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
			want: 9,
		},
		{
			name: "[0,3,7,2,5,8,4,6,0,1,9]",
			args: args{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1, 9}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
