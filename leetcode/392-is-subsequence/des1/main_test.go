package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubsequence(tt.args.s, tt.args.t)
			require.Equal(t, tt.want, got)
		})
	}
}
