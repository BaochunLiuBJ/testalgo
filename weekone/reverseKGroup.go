package weekone

func reverseKGroup(head *ListNode, k int) *ListNode {
	protectNode := &ListNode{
		Next: head,
	}
	lastNode := protectNode

	for head != nil {
		//Get the group's head and tail
		groupHead := head
		groupTail := getTail(groupHead, k)
		if groupTail == nil {
			break
		}
		nextGroupHead := groupTail.Next
		reverseAGroup(groupHead, k)
		lastNode.Next = groupTail  
		groupHead.Next = nextGroupHead
		lastNode = groupHead 
		head = nextGroupHead
	}

	return protectNode.Next
}

func getTail(h *ListNode, k int) (tail *ListNode) {
	tail = h
	for i := 0; i < k-1; i++ {
		tail = tail.Next
		if tail == nil {
			return nil
		}
	}
	return
}

func reverseAGroup(head *ListNode, k int){ 
	h := head
	l := head.Next
	for i := 0; i < k-1; i++ {
		t := l.Next
		l.Next = h 
		h = l 
		l = t 
	}
}
