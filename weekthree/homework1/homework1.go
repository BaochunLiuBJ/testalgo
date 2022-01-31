package homework1

import "sort"

var (
	used      []bool
	resultSet [][]int
	chosenSet []int
)

func permuteUnique(nums []int) [][]int {
	resultSet = make([][]int, 0)
	chosenSet = make([]int, 0)
	used = make([]bool, len(nums))
	if len(nums) == 0 {
		return resultSet
	}
	sort.Ints(nums)
	recur47(nums, 0)
	return resultSet
}

func recur47(nums []int, i int) {
	if len(chosenSet) == len(nums) {
		tmp := make([]int, len(chosenSet))
		copy(tmp, chosenSet)
		resultSet = append(resultSet, tmp)
	}

	for p := 0; p < len(nums); p++ {
		if used[p] || p > 0 && nums[p] == nums[p-1] && !used[p-1] {
			continue
		}
		chosenSet = append(chosenSet, nums[p])
		used[p] = true
		recur47(nums, i+1)
		l := len(chosenSet)
		chosenSet = chosenSet[:l-1]
		used[p] = false
	}
}
