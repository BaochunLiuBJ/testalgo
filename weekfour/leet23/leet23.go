package leet23 

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, l := range lists {
		if l != nil {
			heap.Push(&pq, l)
		}
	}
	head := &ListNode{}
	tail := head
	for pq.Len() > 0 {
		n := heap.Pop(&pq).(*ListNode)
		tail.Next = n
		tail = tail.Next
		next := n.Next
		if next != nil {
			heap.Push(&pq, next)
		}
	}
	return head.Next
}
