package weekone

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		// TODO: Add test cases.
		{"test 1", []int{9}, []int{1, 0}},
		{"test 2", []int{1, 2, 3}, []int{1, 2, 4}},
		{"test 3", []int{}, []int{}},
		{"test 4", []int{9, 9, 9, 9}, []int{1, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
