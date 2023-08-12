package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			name: "case2",
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			name: "case3",
			args: args{
				num: 9,
			},
			want: "IX",
		},
		{
			name: "case4",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "case5",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
		{
			name: "case6",
			args: args{
				num: 20,
			},
			want: "XX",
		},
		{
			name: "case7",
			args: args{
				num: 60,
			},
			want: "LX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intToRoman(tt.args.num)
			require.Equal(t, tt.want, got)
		})
	}
}
