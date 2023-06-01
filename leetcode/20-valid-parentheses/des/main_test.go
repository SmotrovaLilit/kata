package des

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isValid(t *testing.T) {
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
				s: "{{}}",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				s: "]][[",
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "case5",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "case6",
			args: args{
				s: "[(])",
			},
			want: false,
		},
		{
			name: "case7",
			args: args{
				s: "{}{}",
			},
			want: true,
		},
		{
			name: "case8",
			args: args{
				s: "[",
			},
			want: false,
		},
		{
			name: "case9",
			args: args{
				s: "[{}]",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}
