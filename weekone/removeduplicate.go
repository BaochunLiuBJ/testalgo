package weekone

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0 
	}
	n := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[n] = nums[i]
			n++
		}
	}

	return n 
}
