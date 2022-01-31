package leet1245

import "testing"

func Test_treeDiameter(t *testing.T) {
	tests := []struct {
		name string
		edges [][]int
		want int
	}{
		// TODO: Add test cases.
		{ name: "2 width", edges: [][]int{{0,1},{0,2}}, want:2 }, 
        {name: "4 width", edges: [][]int{{0,1},{1,2},{2,3},{1,4},{4,5}}, want:4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeDiameter(tt.edges); got != tt.want {
				t.Errorf("treeDiameter() = %v, want %v", got, tt.want)
			}
		})
	}
}
