package weektwo

func twoSum(nums []int, target int) []int {
	numHash := make(map[int]int)
	r := make([]int, 0)

	for i, n := range nums {
		j, ok := numHash[target-n]
		if ok {
			r = append(r, i)
			r = append(r, j)
			break 
		} else {
			numHash[n] = i 
		}
	}
	return r 
}
