package leet297

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	results []string
	nodes   []string
	current int
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.results = make([]string, 0)
	this.preOrder(root)
	return strings.Join(this.results, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.current = 0
	this.nodes = strings.Split(data, ",")
	return this.restore()
}

func (this *Codec) preOrder(root *TreeNode) {
	if root == nil {
		this.results = append(this.results, "nil")
	}
	this.results = append(this.results, strconv.Itoa(root.Val))
	this.preOrder(root.Left)
	this.preOrder(root.Right)
}

func (this *Codec) restore() *TreeNode {
	if this.nodes[this.current] == "nil" {
		this.current++
		return nil
	}
	if len(this.nodes) == this.current {
		return nil
	}
	n := &TreeNode{}
	v, _ := strconv.Atoi(this.nodes[this.current])
	n.Val = v
	this.current++
	n.Left = this.restore()
	n.Right = this.restore()
	return n
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
