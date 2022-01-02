package weekone

type element struct {
	value int
	pre   *element
	back  *element
}
type MyCircularDeque struct {
	front element
	end   element
	cap   int
	len   int
}

func Constructor(k int) MyCircularDeque {
	m := MyCircularDeque{cap: k}
	m.front.back = &m.end
	m.front.pre = nil
	m.end.pre = &m.front
	m.end.back = nil
	return m
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.cap == this.len {
		return false
	}
	e := &element{
		value: value,
	}
	e.back = this.front.back
	e.pre = &this.front
	this.front.back = e
	this.len++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.cap == this.len {
		return false
	}
	e := &element{
		value: value,
	}
	e.pre = this.end.pre
	e.back = &this.end
	this.end.pre = e
	this.len++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.len == 0 {
		return false
	}
	this.front.back = this.front.back.back
	this.len--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.len == 0 {
		return false
	}
	this.end.pre = this.end.pre.pre
	this.len--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.len == 0 {
		return -1
	}
	return this.front.back.value
}

func (this *MyCircularDeque) GetRear() int {
	if this.len == 0 {
		return -1 
	}
	return this.end.pre.value
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.len == 0 
}

func (this *MyCircularDeque) IsFull() bool {
	return this.len == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
