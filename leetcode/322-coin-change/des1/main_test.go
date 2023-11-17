package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{coins: []int{1, 2, 5}, amount: 11},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{coins: []int{2}, amount: 3},
			want: -1,
		},
		{
			name: "Example 3",
			args: args{coins: []int{1}, amount: 0},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{coins: []int{1}, amount: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coinChange(tt.args.coins, tt.args.amount)
			require.Equal(t, tt.want, got)
		})
	}
}
