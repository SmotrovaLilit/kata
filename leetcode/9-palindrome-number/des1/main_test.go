package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case121",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "case-121",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "case10",
			args: args{
				x: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.args.x)
			require.Equal(t, tt.want, got)
		})
	}
}
