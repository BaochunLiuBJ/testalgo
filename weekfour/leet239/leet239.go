package leet239

import "fmt"

type node struct {
	value int
	index int
}
type binaryHeap struct {
	heap []node
}

func newBinaryHeap() *binaryHeap {
	heap := make([]node, 1)
	return &binaryHeap{
		heap: heap,
	}
}

func (bh *binaryHeap) empty() bool {
	return len(bh.heap) == 1
}

func (bh *binaryHeap) insert(v int, index int) {
	bh.heap = append(bh.heap, node{v, index})
	bh.heapUp(len(bh.heap) - 1)
}

func (bh *binaryHeap) getFront() node {
	if len(bh.heap) == 1 {
		return node{}
	}
	return bh.heap[1]
}

func (bh *binaryHeap) removeFront() {
	if len(bh.heap) == 1 {
		return
	} else {
		bh.heap[1] = bh.heap[len(bh.heap)-1]
		bh.heap = bh.heap[0 : len(bh.heap)-1]
		bh.heapDown(1)
	}
}

func (bh *binaryHeap) heapUp(p int) {
	for p > 1 {
		if bh.heap[p].value > bh.heap[p/2].value {
			bh.heap[p], bh.heap[p/2] = bh.heap[p/2], bh.heap[p]
		} else {
			break
		}
		p = p/2 
	}
}

func (bh *binaryHeap) heapDown(p int) {
	if p < 0 {
		return
	}
	i := p * 2
	for i < len(bh.heap) {
		j := p*2 + 1
		if j < len(bh.heap) && bh.heap[j].value > bh.heap[i].value {
			i = j
		}
		if bh.heap[p].value < bh.heap[i].value {
			bh.heap[p], bh.heap[i] = bh.heap[i], bh.heap[p]
			p = i
			i = 2 * p
		} else {
			break
		}
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	resultSlice := make([]int, 0)
	bh := newBinaryHeap()
	for i := 0; i < len(nums); i++ {
		bh.insert(nums[i], i)
		if i >= k-1 {
			// i-k+1, i max
			for bh.getFront().index <= i-k {
				bh.removeFront()
			}
			resultSlice = append(resultSlice, bh.getFront().value)
		}
	}
	return resultSlice
}
