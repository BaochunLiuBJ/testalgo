func moveZeroes(nums []int) {
	n := 0
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n] = nums[i]
			n++
		}
	}
	for ; n < len(nums); n++ {
		nums[n] = 0
	}
	return

}