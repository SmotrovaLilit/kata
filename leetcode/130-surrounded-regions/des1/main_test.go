package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "case1",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "case2",
			args: args{
				board: [][]byte{
					{'X'},
				},
			},
			want: [][]byte{
				{'X'},
			},
		},
		{
			name: "case4",
			args: args{
				board: [][]byte{
					{'X', 'O', 'X', 'O', 'X', 'O'},
					{'O', 'X', 'O', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'O', 'X', 'O'},
					{'O', 'X', 'O', 'X', 'O', 'X'},
				},
			},
			want: [][]byte{
				{'X', 'O', 'X', 'O', 'X', 'O'},
				{'O', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'X'},
			},
		},
		{
			name: "case5",
			args: args{
				board: [][]byte{
					{'O', 'O', 'O', 'O', 'X', 'X'},
					{'O', 'O', 'O', 'O', 'O', 'O'},
					{'O', 'X', 'O', 'X', 'O', 'O'},
					{'O', 'X', 'O', 'O', 'X', 'O'},
					{'O', 'X', 'O', 'X', 'O', 'O'},
					{'O', 'X', 'O', 'O', 'O', 'O'},
				},
			},
			want: [][]byte{
				{'O', 'O', 'O', 'O', 'X', 'X'},
				{'O', 'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'O', 'O'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			require.Equal(t, tt.want, tt.args.board)
		})
	}
}
