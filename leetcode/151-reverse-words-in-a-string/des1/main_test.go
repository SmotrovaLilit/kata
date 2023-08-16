package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "case2",
			args: args{
				s: "  hello world  ",
			},
			want: "world hello",
		},
		{
			name: "case3",
			args: args{
				s: "a good   example",
			},
			want: "example good a",
		},
		{
			name: "case4",
			args: args{
				s: "  Bob    Loves  Alice   ",
			},
			want: "Alice Loves Bob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}
