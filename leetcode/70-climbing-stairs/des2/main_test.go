package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "case3",
			args: args{
				n: 4,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.args.n)
			require.Equal(t, tt.want, got)
		})
	}
}
