package leet207

import "testing"

func Test_canFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		// TODO: Add test cases.
		{"false case 1", 2, [][]int{{1, 0}, {0, 1}}, false},
		{"false case 2", 20, [][]int{{0,10},{3,18},{5,5},{6,11},{11,14},{13,1},{15,1},{17,4}}, false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.numCourses, tt.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
