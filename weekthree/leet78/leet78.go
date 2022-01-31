package leet78 

var (
	resultSet [][]int
	chosenSet []int
)

func subsets(nums []int) [][]int {
	resultSet = make([][]int, 0)
	chosenSet = make([]int, 0)
	choose(nums, 0)
	return resultSet
}

func choose(nums []int, i int) {
	if i == len(nums) {
		tmp := make([]int, len(chosenSet))
		copy(tmp, chosenSet)
		resultSet = append(resultSet, tmp)
		return
	}
	choose(nums, i+1)
	chosenSet = append(chosenSet, nums[i])
	choose(nums, i+1)
	l := len(chosenSet)
	chosenSet = chosenSet[:l-1]
}
