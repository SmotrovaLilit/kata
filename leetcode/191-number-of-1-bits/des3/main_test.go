package des3

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				num: 11,
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				num: 128,
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				num: 4294967293,
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hammingWeight(tt.args.num)
			require.Equal(t, tt.want, got)
		})
	}
}
