package weekthree


func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	aw := make([]string, 0)
	for i := 1; i <= n; i++ {
		A := generateParenthesis(i - 1)
		B := generateParenthesis(n - i)
		for _, a := range A {
			for _, b := range B {
				r := "(" + a + ")" + b
				aw = append(aw, r)
			}
		}
	}
	return aw
}
