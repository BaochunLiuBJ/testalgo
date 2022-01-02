package weekone

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var header *ListNode
	m := list1
	n := list2
	
	if m.Val < n.Val {
		header = m
	} else {
		header = n 
	}
	var l *ListNode = nil  	
	for m != nil && n != nil {
		if m.Val < n.Val {
			if l == nil {
				l = m 
			} else {
				l.Next = m 
				l = l.Next
			}
			m = m.Next
		} else {
			if l == nil {
				l = n
			} else {
				l.Next = n 
				l = l.Next	
			} 
			n = n.Next 
		}
	}
	if m == nil {
		l.Next = n 
	} else {
		l.Next = m 
	}

	return header
}
