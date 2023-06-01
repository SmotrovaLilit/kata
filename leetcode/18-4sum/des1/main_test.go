package des1

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			want: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			name: "case2",
			args: args{
				nums:   []int{2, 2, 2, 2, 2},
				target: 8,
			},
			want: [][]int{
				{2, 2, 2, 2},
			},
		},
		{
			name: "case3",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.args.nums, tt.args.target)
			require.Equal(t, len(got), len(tt.want), fmt.Sprintf("expected: %v, actual: %v", tt.want, got))
			for _, v := range got {
				if !exists(v, tt.want) {
					t.Errorf("%v not exist in expected lists", v)
					return
				}
			}
		})
	}
}

func exists(got []int, want [][]int) bool {
	for _, expList := range want {
		ex := true
		for _, gotValue := range got {
			if !existValue(gotValue, expList) {
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

func existValue(target int, list []int) bool {
	for _, v := range list {
		if target == v {
			return true
		}
	}
	return false
}
