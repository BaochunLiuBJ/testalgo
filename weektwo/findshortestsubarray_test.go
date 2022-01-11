package weektwo

import "testing"

func Test_findShortestSubArray(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		// TODO: Add test cases.
		{"test one", []int{1, 2, 2, 3, 1}, 2},
		{"test two", []int{1, 2, 2, 3, 1, 4, 2}, 6},
		{"test 3 ", []int{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray(tt.nums); got != tt.want {
				t.Errorf("findShortestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
