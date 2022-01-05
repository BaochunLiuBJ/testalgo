package weekone

func reverseList(head *ListNode) *ListNode {
	var t *ListNode = nil

	for head != nil {
		n := head.Next
		head.Next = t
		t = head
		head = n
	}

	return t
}
