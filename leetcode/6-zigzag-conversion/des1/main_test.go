package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "case2",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "case3",
			args: args{
				s:       "A",
				numRows: 1,
			},
			want: "A",
		},
		{
			name: "case4",
			args: args{
				s:       "PAHNAPLSIIGYIR",
				numRows: 1,
			},
			want: "PAHNAPLSIIGYIR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convert(tt.args.s, tt.args.numRows)
			require.Equal(t, tt.want, got)
		})
	}
}
