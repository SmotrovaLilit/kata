package des1

import (
	"reflect"
	"testing"
)

func TestLinkedList_Remove(t *testing.T) {
	tests := []struct {
		name   string
		list   []int
		remove int
		want   []int
	}{
		{
			name:   "empty list",
			list:   []int{},
			remove: 1,
			want:   []int{},
		},
		{
			name:   "remove in the middle",
			list:   []int{1, 2, 3, 4, 5},
			remove: 3,
			want:   []int{1, 2, 4, 5},
		},
		{
			name:   "remove first",
			list:   []int{1, 2, 3, 4, 5},
			remove: 1,
			want:   []int{2, 3, 4, 5},
		},
		{
			name:   "remove last",
			list:   []int{1, 2, 3, 4, 5},
			remove: 5,
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "remove several times",
			list:   []int{3, 3, 2, 3, 4, 5, 3, 3},
			remove: 3,
			want:   []int{2, 4, 5},
		},
		{
			name:   "remove all",
			list:   []int{3, 3, 3, 3, 3, 3, 3},
			remove: 3,
			want:   []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := CreateFromSlice(tt.list)
			l.Remove(tt.remove)
			if !reflect.DeepEqual(l.ToSlice(), tt.want) {
				t.Errorf("Remove() = %v, want %v", l.ToSlice(), tt.want)
			}
		})
	}
}
