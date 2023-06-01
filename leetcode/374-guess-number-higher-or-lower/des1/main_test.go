package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_guessNumber(t *testing.T) {
	type args struct {
		n    int
		pick int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n:    10,
				pick: 6,
			},
			want: 6,
		},
		{
			name: "case2",
			args: args{
				n:    1,
				pick: 1,
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				n:    2,
				pick: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			guess = testGuess(tt.args.pick)
			got := guessNumber(tt.args.n)
			require.Equal(t, tt.want, got)
		})
	}
}
