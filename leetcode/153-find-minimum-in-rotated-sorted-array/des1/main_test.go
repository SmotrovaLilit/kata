package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_findMin(t *testing.T) {
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
				nums: []int{3, 4, 5, 1, 2},
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 0,
		},
		{
			name: "case3",
			args: args{
				nums: []int{11, 13, 15, 17},
			},
			want: 11,
		},
		{
			name: "case4",
			args: args{
				nums: []int{17, 11, 13, 15},
			},
			want: 11,
		},
		{
			name: "case5",
			args: args{
				nums: []int{13, 15, 17, 11},
			},
			want: 11,
		},
		{
			name: "case6",
			args: args{
				nums: []int{13},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMin(tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
