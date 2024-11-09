package des1

import (
	"reflect"
	"testing"
)

func Test_sum(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{
			name:  "empty lists",
			list1: []int{},
			list2: []int{},
			want:  []int{},
		},
		{
			name:  "empty list 1",
			list1: []int{},
			list2: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty list 2",
			list1: []int{1, 2, 3},
			list2: []int{},
			want:  []int{1, 2, 3},
		},
		{
			name:  "[7, 1 , 6] + [5, 9, 2]  = [2, 1, 9]",
			list1: []int{7, 1, 6},
			list2: []int{5, 9, 2},
			want:  []int{2, 1, 9},
		},
		{
			name:  "[7, 1 , 6, 1] + [5, 9, 2]  = [2, 1, 9, 1]",
			list1: []int{7, 1, 6, 1},
			list2: []int{5, 9, 2},
			want:  []int{2, 1, 9, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sum(CreateFromSlice(tt.list1), CreateFromSlice(tt.list2))
			if !reflect.DeepEqual(got.ToSlice(), tt.want) {
				t.Errorf("sum() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func Test_sumForward(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{
			name:  "empty lists",
			list1: []int{},
			list2: []int{},
			want:  []int{},
		},
		{
			name:  "empty list 1",
			list1: []int{},
			list2: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty list 2",
			list1: []int{1, 2, 3},
			list2: []int{},
			want:  []int{1, 2, 3},
		},
		{
			name:  "[6, 1, 7] + [2, 9, 5]  = [9, 1, 2]",
			list1: []int{6, 1, 7},
			list2: []int{2, 9, 5},
			want:  []int{9, 1, 2},
		},
		{
			name:  "[1->6->1->7] + [2->9->5]  = [1->9->1->2]",
			list1: []int{1, 6, 1, 7},
			list2: []int{2, 9, 5},
			want:  []int{1, 9, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumForward(CreateFromSlice(tt.list1), CreateFromSlice(tt.list2))
			if !reflect.DeepEqual(got.ToSlice(), tt.want) {
				t.Errorf("sumForward() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}
