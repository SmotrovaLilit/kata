package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{n: 19},
			want: true,
		},
		{
			name: "ex2",
			args: args{n: 2},
			want: false,
		},
		{
			name: "ex3",
			args: args{n: 7},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isHappy(tt.args.n)
			require.Equal(t, tt.want, got)
		})
	}
}
