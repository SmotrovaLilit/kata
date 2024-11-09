package des2

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		list     LinkedList
		expected []int
	}{
		{
			name:     "empty list",
			list:     CreateFromSlice([]int{}),
			expected: []int{},
		},
		{
			name:     "no duplicates",
			list:     CreateFromSlice([]int{1, 2, 3}),
			expected: []int{1, 2, 3},
		},
		{
			name:     "duplicates in the middle",
			list:     CreateFromSlice([]int{1, 2, 4, 0, 2, 3}),
			expected: []int{1, 2, 4, 0, 3},
		},
		{
			name:     "duplicates in the middle near each other",
			list:     CreateFromSlice([]int{1, 2, 2, 4, 3}),
			expected: []int{1, 2, 4, 3},
		},
		{
			name:     "duplicates at the end",
			list:     CreateFromSlice([]int{1, 2, 3, 3}),
			expected: []int{1, 2, 3},
		},
		{
			name:     "duplicates at the beginning",
			list:     CreateFromSlice([]int{1, 1, 2, 3}),
			expected: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.list.ToSlice()
			RemoveDuplicates(tt.list)
			if !reflect.DeepEqual(tt.expected, tt.list.ToSlice()) {
				t.Errorf("%v: expected %v, got %v", input, tt.expected, tt.list.ToSlice())
			}
		})
	}
}
