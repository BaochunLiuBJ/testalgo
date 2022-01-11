package weektwo

func maxSubArray(nums []int) int {
	l := len(nums)
	sums := make([]int, l+1)
	preMins := make([]int, l+1)

	sums[0] = 0
	for i := 1; i <= l; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	preMins[0] = 0
	for i := 1; i <= l; i++ {
		if preMins[i-1] < sums[i] {
			preMins[i] = preMins[i-1]
		} else {
			preMins[i] = sums[i]
		}
	}

	result := int(-10000)
	for i := 1; i <= l; i++ {
		if result < sums[i]-preMins[i-1] {
			result = sums[i] - preMins[i-1]
		}
	}
	return result 	
}
