package des2

import (
	"reflect"
	"testing"
)

func TestMultiplyExceptCurrentIndex(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []int
	}{
		{
			name:       "empty list",
			args:       args{input: []int{}},
			wantOutput: []int{},
		},
		{
			name:       "[1, 2, 3, 4, 5] => [120, 60, 40, 30, 24]",
			args:       args{input: []int{1, 2, 3, 4, 5}},
			wantOutput: []int{120, 60, 40, 30, 24},
		},
		{
			name:       "[3, 2, 1] => [2, 3, 6]",
			args:       args{input: []int{3, 2, 1}},
			wantOutput: []int{2, 3, 6},
		},
		{
			name:       "[3, 2, 1, 0] => [0, 0, 0, 6]",
			args:       args{input: []int{3, 2, 1, 0}},
			wantOutput: []int{0, 0, 0, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MultiplyExceptCurrentIndex(tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MultiplyExceptCurrentIndex() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
