package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_romanToInt(t *testing.T) {
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
				s: "III",
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				s: "IV",
			},
			want: 4,
		},
		{
			name: "case3",
			args: args{
				s: "IX",
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}
