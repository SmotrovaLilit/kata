package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				x: 8,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mySqrt(tt.args.x)
			require.Equal(t, tt.want, got)
		})
	}
}
