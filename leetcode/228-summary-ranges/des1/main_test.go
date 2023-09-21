package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{nums: []int{0, 1, 2, 4, 5, 7}},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			name: "test2",
			args: args{nums: []int{0, 2, 3, 4, 6, 8, 9}},
			want: []string{"0", "2->4", "6", "8->9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := summaryRanges(tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
