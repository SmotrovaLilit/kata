package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "Example 1",
			args: args{
				num: 0b00000010100101000001111010011100,
			},
			want: 0b00111001011110000010100101000000,
		},
		{
			name: "Example 2",
			args: args{
				num: 0b1111111111,
			},
			want: 0b11111111110000000000000000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBits(tt.args.num)
			require.Equal(t, tt.want, got)
		})
	}
}
