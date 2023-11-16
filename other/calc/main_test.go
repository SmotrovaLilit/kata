package calc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_calc(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example 1",
			args: args{str: "1 + 1"},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{str: "5 - 1 - 1"},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{str: "84 / 1 / 2"},
			want: 42,
		},
		{
			name: "Example 4",
			args: args{str: "5/2 - 2*3"},
			want: -3.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc(tt.args.str)
			require.Equal(t, tt.want, got)
		})
	}
}
