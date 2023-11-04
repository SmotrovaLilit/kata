package des1

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math"
	"math/rand"
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				target: 4,
				nums:   []int{1, 4, 4},
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				target: 11,
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: 0,
		},
		{
			name: "case4",
			args: args{
				target: 213,
				nums:   []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSubArrayLen(tt.args.target, tt.args.nums)
			require.Equal(t, tt.want, got)
		})
	}
}

func Benchmark_minSubArrayLen(b *testing.B) {
	for j := 0; j < 20; j++ {
		count := int(math.Pow(2, float64(j)))
		b.Run(fmt.Sprintf("n=%d", count), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				minSubArrayLen(rand.Intn(count), rand.Perm(count))
			}
		})
	}

}
