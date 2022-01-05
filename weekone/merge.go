package weekone

func merge(nums1 []int, m int, nums2 []int, n int) {
	p := m
	q := n
	i := m + n
	for p > 0 && q > 0 {
		if nums1[p-1] > nums2[q-1] {
			nums1[i-1] = nums1[p-1]
			p--
			i--
		} else {
			nums1[i-1] = nums2[q-1]
			q--
			i--
		}
	}

	if p == 0 {
		for q > 0 {
			nums1[i-1] = nums2[q-1]
			q--
			i--
		}
	}
	if q == 0 {
		for p > 0 {
			nums1[i-1] = nums1[p-1]
			p--
			i--
		}
	}
}
