package des1

import "testing"

func Test_getSumOfTarget(t *testing.T) {
	tests := []struct {
		name   string
		list   []int
		target int
		want1  int
		want2  int
	}{
		{
			name:   "case1",
			list:   []int{1, 2, 3, 4, 5},
			target: 5,
			want1:  0,
			want2:  3,
		},
		{
			name:   "case2",
			list:   []int{1, 2, 3, 4, 5},
			target: 6,
			want1:  0,
			want2:  4,
		},
		{
			name:   "case3",
			list:   []int{2, 7, 11, 15},
			target: 9,
			want1:  0,
			want2:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := getSumOfTarget(tt.list, tt.target)
			if got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("expected: %v, %v, actual: %v, %v", tt.want1, tt.want2, got1, got2)
			}
		})
	}
}
