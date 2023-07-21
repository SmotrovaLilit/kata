package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				prices: []int{
					7, 1, 5, 3, 6, 4,
				},
			},
			want: 5,
		},
		{
			name: "case2",
			args: args{
				prices: []int{
					7, 6, 4, 3, 1,
				},
			},
			want: 0,
		},
		{
			name: "case3",
			args: args{
				prices: []int{
					7, 5, 5, 3, 6, 1,
				},
			},
			want: 3,
		},
		{
			name: "case4",
			args: args{
				prices: []int{
					7, 5, 5, 10, 1, 2,
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.args.prices)
			require.Equal(t, tt.want, got)
		})
	}
}
