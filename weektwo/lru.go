package weektwo

import "fmt"

type LRUCache struct {
	head *node
	tail *node
	cap  int
	lMap map[int]*node
}

type node struct {
	key   int
	value int
	pre   *node
	next  *node
}

func Constructor(capacity int) LRUCache {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.pre = head
	lMap := make(map[int]*node)
	cache := LRUCache{
		head: head,
		tail: tail,
		lMap: lMap,
		cap:  capacity,
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.lMap[key]
	if !ok {
		return -1
	} else {
		this.removeNode(n)
		this.insertFront(n)
		return n.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.lMap[key]
	if !ok {
		n = &node{
			key:   key,
			value: value,
		}
		this.lMap[key] = n
		if len(this.lMap) > this.cap {
			tn := this.tail.pre
			this.removeNode(tn)
			delete(this.lMap,tn.key)
			tn = nil
			this.insertFront(n)
		} else {
			this.insertFront(n)
		}
	} else {
		n.value = value
		this.removeNode(n)
		this.insertFront(n)
	}

	for h:= this.head ; h.next != nil ; h= h.next {
		fmt.Printf("%d, %d \n", h.key, h.value)	
	}
}

func (this *LRUCache) insertFront(n *node) {
	n.next = this.head.next
	n.pre = this.head
	this.head.next.pre = n
	this.head.next = n
}

func (this *LRUCache) removeNode(n *node) {
	n.pre.next = n.next
	n.next.pre = n.pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
