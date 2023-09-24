package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		{
			name: "Example 2",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addBinary(tt.args.a, tt.args.b)
			require.Equal(t, tt.want, got)
		})
	}
}
