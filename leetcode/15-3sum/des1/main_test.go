package des1

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{-1, -1, 0, 1, 2, -4},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "case2",
			args: args{
				nums: []int{0, 1, 1},
			},
			want: [][]int{},
		},
		{
			name: "case3",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "case4",
			args: args{
				nums: []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0},
			},
			want: [][]int{
				{3, -4, 1}, {0, -4, 4}, {1, -2, 1}, {-2, -2, 4}, {4, 1, -5}, {0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.args.nums)
			require.Equal(t, len(got), len(tt.want), fmt.Sprintf("expected: %v, actual: %v", tt.want, got))
			for _, v := range got {
				if !existsT(v, tt.want) {
					t.Errorf("%v not exist in expected lists", v)
					return
				}
			}
		})
	}
}

func existsT(got []int, want [][]int) bool {
	for _, expList := range want {
		ex := true
		for _, gotValue := range got {
			if !existValueT(gotValue, expList) {
				ex = false
				break
			}
		}
		if ex {
			return true
		}
	}
	return false
}

func existValueT(target int, list []int) bool {
	for _, v := range list {
		if target == v {
			return true
		}
	}
	return false
}
