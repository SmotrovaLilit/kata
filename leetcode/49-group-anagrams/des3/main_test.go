package des1

import (
	"reflect"
	"testing"
)

func exists(got []int, want [][]int) bool {
	for _, expList := range want {
		ex := true
		for _, gotValue := range got {
			if !existValue(gotValue, expList) {
				ex = false
				break
			}
		}
		if ex {
			return true
		}
	}
	return false
}

func existValue(target int, list []int) bool {
	for _, v := range list {
		if target == v {
			return true
		}
	}
	return false
}

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "case1",
			args: args{
				strs: []string{
					"eat", "tea", "tan", "ate", "nat", "bat",
				},
			},
			want: [][]string{
				{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"},
			},
		},
		{
			name: "case2",
			args: args{
				strs: []string{""},
			},
			want: [][]string{
				{""},
			},
		},
		{
			name: "case3",
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{
				{"a"},
			},
		},
		{
			name: "case4",
			args: args{
				strs: []string{"ddddddddddg", "dgggggggggg"},
			},
			want: [][]string{
				{
					"ddddddddddg",
				},
				{
					"dgggggggggg",
				},
			},
		},
		{
			name: "case5",
			args: args{
				strs: []string{"bdddddddddd", "bbbbbbbbbbc"},
			},
			want: [][]string{
				{
					"bdddddddddd",
				},
				{
					"bbbbbbbbbbc",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
