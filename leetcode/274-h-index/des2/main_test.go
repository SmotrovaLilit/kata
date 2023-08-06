package des1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case0",
			args: args{citations: []int{1, 3, 2, 3, 100}},
			want: 3,
		},
		{
			name: "case1",
			args: args{citations: []int{3, 0, 6, 1, 5}},
			want: 3,
		},
		{
			name: "case2",
			args: args{citations: []int{1, 3, 1}},
			want: 1,
		},
		{
			name: "case3",
			args: args{citations: []int{1, 1, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hIndex(tt.args.citations)
			require.Equal(t, tt.want, got)
		})
	}
}
