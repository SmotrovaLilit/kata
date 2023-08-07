package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"[1,2,3,4]",
			args{[]int{1, 2, 3, 4}},
			[]int{24, 12, 8, 6},
		},
		{
			"[1,2,3,4,5]",
			args{[]int{1, 2, 3, 4, 5}},
			[]int{120, 60, 40, 30, 24},
		},
		{
			"[1,2,3,4,5,6]",
			args{[]int{1, 2, 3, 4, 5, 6}},
			[]int{720, 360, 240, 180, 144, 120},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
