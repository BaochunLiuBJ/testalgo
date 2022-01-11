package weektwo

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		input []string  
		want [][]string
	}{
		// TODO: Add test cases.
		{"testOne",[]string{"eat","tea","tan","ate","nat","bat"}, [][]string{{"bat"},{"nat","tan"},{"ate","eat","tea"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
