package des1

import (
	"reflect"
	"testing"
)

func TestGetSubList(t *testing.T) {
	type args struct {
		start int
		list  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty list",
			args: args{start: 0, list: []int{}},
			want: []int{},
		},
		{
			name: "start from 0",
			args: args{start: 0, list: []int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "start from 1",
			args: args{start: 1, list: []int{1, 2, 3, 4, 5}},
			want: []int{2, 3, 4, 5},
		},
		{
			name: "start from last",
			args: args{start: 4, list: []int{1, 2, 3, 4, 5}},
			want: []int{5},
		},
		{
			name: "start from out of range",
			args: args{start: 5, list: []int{1, 2, 3, 4, 5}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		got := GetSubList(tt.args.start, CreateFromSlice(tt.args.list))
		if !reflect.DeepEqual(got.ToSlice(), tt.want) {
			t.Errorf("GetSubList() = %v, want %v", got.ToSlice(), tt.want)
		}
	}

}
