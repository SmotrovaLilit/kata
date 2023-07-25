package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "case2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			want: []int{1},
		},
		{
			name: "case3",
			args: args{
				nums1: []int{},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums1 := make([]int, tt.args.m+tt.args.n)
			copy(nums1, tt.args.nums1)
			merge(nums1, tt.args.m, tt.args.nums2, tt.args.n)
			require.Equal(t, tt.want, nums1)
		})
	}
}
