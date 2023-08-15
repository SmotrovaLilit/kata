package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				strs: []string{
					"flower", "flow", "flight",
				},
			},
			want: "fl",
		},
		{
			name: "case2",
			args: args{
				strs: []string{
					"dog", "racecar", "car",
				},
			},
			want: "",
		},
		{
			name: "case3",
			args: args{
				strs: []string{
					"dog", "dog", "dog",
				},
			},
			want: "dog",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.args.strs)
			require.Equal(t, tt.want, got)
		})
	}
}
