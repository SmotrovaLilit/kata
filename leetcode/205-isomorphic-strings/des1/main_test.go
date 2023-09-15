package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
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
				s: "egg",
				t: "add",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				s: "badc",
				t: "baba",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isIsomorphic(tt.args.s, tt.args.t)
			require.Equal(t, tt.want, got)
		})
	}
}
