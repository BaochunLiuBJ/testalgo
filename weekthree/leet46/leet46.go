package leet46 

var (
	used []bool
	resultSet [][]int
	chosenSet []int
)

func permute(nums []int) [][]int {
	resultSet = make([][]int, 0)
	chosenSet = make([]int, 0)
	used = make([]bool, len(nums))
	recur46(nums, 0)
	return resultSet
}

func recur46(nums []int, i int) {
	if len(chosenSet) == len(nums) {
		tmp := make([]int, len(chosenSet))
		copy(tmp, chosenSet)
		resultSet = append(resultSet, tmp)
	}

	for p := 0; p < len(nums); p++ {
		if !used[p] {
			chosenSet = append(chosenSet, nums[p])
			used[p] = true
			recur46(nums, i+1)
			l := len(chosenSet)
			chosenSet = chosenSet[:l-1]
			used[p] = false 
		}
	}
}
