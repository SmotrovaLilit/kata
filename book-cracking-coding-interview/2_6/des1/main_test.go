package des1

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		want bool
		list []int
	}{
		{
			name: "empty list",
			list: []int{},
			want: true,
		},
		{
			name: "one element",
			list: []int{1},
			want: true,
		},
		{
			name: "two elements",
			list: []int{1, 1},
			want: true,
		},
		{
			name: "two elements different",
			list: []int{1, 2},
			want: false,
		},
		{
			name: "three elements",
			list: []int{1, 2, 1},
			want: true,
		},
		{
			name: "three elements different",
			list: []int{1, 2, 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := CreateFromSlice(tt.list)
			if got := Palindrome(l); got != tt.want {
				t.Errorf("Palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
