package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{nums: []int{1, 2, 3, 1}, k: 3},
			want: true,
		},
		{
			name: "ex2",
			args: args{nums: []int{1, 0, 1, 1}, k: 1},
			want: true,
		},
		{
			name: "ex3",
			args: args{nums: []int{1, 2, 3, 1, 2, 3}, k: 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsNearbyDuplicate(tt.args.nums, tt.args.k)
			require.Equal(t, tt.want, got)
		})
	}
}
