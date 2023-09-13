package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				s: "  ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}
