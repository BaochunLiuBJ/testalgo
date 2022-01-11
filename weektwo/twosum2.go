package weektwo

func twoSumOrder(numbers []int, target int) []int {
	j := len(numbers) - 1
	for i := 0; i != j; i++ {
		for i < j && (numbers[i]+numbers[j] > target) {
			j--
		}
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}
