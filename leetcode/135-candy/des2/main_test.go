package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				ratings: []int{1, 0, 2},
			},
			want: 5,
		},
		{
			name: "case2",
			args: args{
				ratings: []int{1, 2, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := candy(tt.args.ratings)
			require.Equal(t, tt.want, got)
		})
	}
}
