package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "case4",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "case5",
			args: args{
				s: "abba",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}
