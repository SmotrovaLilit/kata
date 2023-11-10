package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-1",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "test-2",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "test-3",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := simplifyPath(tt.args.path)
			require.Equal(t, tt.want, got)
		})
	}
}
