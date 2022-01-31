package weekthree

func myPow(x float64, n int) float64 {
	var negative bool
	t := 0
	if n == 0 {
		return 1
	} else if n < 0 {
		t = 0 - n
		negative = true
	} else {
		t = n
	}

	if t%2 == 0 {
		r := myPow(x, t/2)
		r = r * r
		if negative {
			return 1 / r
		} else {
			return r
		}
	} else {
		r := myPow(x, (t-1)/2)
		r = r * r * x
		if negative {
			return 1 / r
		} else {
			return r
		}
	}
}
