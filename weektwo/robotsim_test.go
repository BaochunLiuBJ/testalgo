package weektwo

import "testing"

func Test_robotSim(t *testing.T) {
	type args struct {
		commands  []int
		obstacles [][]int
	}
	tests := []struct {
		name string
		commands []int
		obstacles [][]int 
		want int
	}{
		// TODO: Add test cases.
		{"case1", []int{4, -1, 3}, [][]int{{100, 100}}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotSim(tt.commands, tt.obstacles); got != tt.want {
				t.Errorf("robotSim() = %v, want %v", got, tt.want)
			}
		})
	}
}
