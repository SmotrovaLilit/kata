package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canConstruct(tt.args.ransomNote, tt.args.magazine)
			require.Equal(t, tt.want, got)
		})
	}
}
