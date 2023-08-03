package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
			},
			want:     5,
			wantNums: []int{1, 1, 2, 2, 3},
		},
		{
			name: "case2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want:     7,
			wantNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantNums, tt.args.nums[:got])
		})
	}
}
