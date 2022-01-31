package leet429 

import (
	"container/list"
)

type Node struct {
	Val      int
	Children []*Node
}

type nodeWithDepth struct {
	node  *Node
	depth int
}

func levelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	queue := list.New()

	if root == nil {
		return result
	}

	queue.PushBack(&nodeWithDepth{
		node:  root,
		depth: 0,
	})

	for queue.Len() > 0 {
		e := queue.Front()
		nd, _ := e.Value.(*nodeWithDepth)

		if nd.depth >= len(result) {
			result = append(result, make([]int, 0))
		}

		l := result[nd.depth]
		l = append(l, nd.node.Val)
		result[nd.depth] = l

		for _, n := range nd.node.Children {
			queue.PushBack(&nodeWithDepth{
				node:  n,
				depth: nd.depth + 1,
			})
		}

		queue.Remove(e)
	}

	return result
}
