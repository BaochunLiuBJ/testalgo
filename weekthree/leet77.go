package weekthree

func combine(n int, k int) [][]int {
	resultSet := make([][]int, 0)
	chosenSet := make([]int, 0)
	recur(0, n, k)
	return resultSet
}

func recur(i int, n int, k int) {
	if len(chosenSet) == k {
		tmp := make([]int, len(chosenSet))
		copy(tmp, chosenSet)
		resultSet = append(resultSet, tmp)
		return
	}
	if len(chosenSet) > k || len(chosenSet)+n-i+1 < k || i == n {
		return
	}
	recur(i+1, n, k)
	chosenSet = append(chosenSet, i+1)
	recur(i+1, n, k)
	l := len(chosenSet)
	chosenSet = chosenSet[:l-1]
}
