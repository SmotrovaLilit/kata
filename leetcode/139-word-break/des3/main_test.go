package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{s: "leetcode", wordDict: []string{"leet", "code"}},
			want: true,
		},
		{
			name: "Example 2",
			args: args{s: "applepenapple", wordDict: []string{"apple", "pen"}},
			want: true,
		},
		{
			name: "Example 3",
			args: args{s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}},
			want: false,
		},
		{
			name: "Example 4",
			args: args{s: "cars", wordDict: []string{"car", "ca", "rs"}},
			want: true,
		},
		{
			name: "Example 5",
			args: args{s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
				wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.args.s, tt.args.wordDict)
			require.Equal(t, tt.want, got)
		})
	}
}
