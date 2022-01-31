package leet1245

import "container/list"

type tree struct {
	edges          [][]int
	nodeToEdgesMap map[int][]*edge
	startNode      int
}

type farthestNode struct {
	startNode int
	to        int
	depth     int
}

type edge struct {
	from     int
	to       int
	accessed bool
}

func treeDiameter(edges [][]int) int {
	t := newTree(edges)
	if t == nil {
		return 0
	}
	fnode := t.findDiameter(t.startNode)
	fnode = t.findDiameter(fnode.to)
	return fnode.depth
}

func newTree(edges [][]int) *tree {
	if len(edges) == 0 {
		return nil
	}

	t := &tree{
		edges: edges,
	}
	t.initTree()
	return t
}
func (t *tree) initTree() {
	t.startNode = t.edges[0][0]
	t.nodeToEdgesMap = make(map[int][]*edge, 0)
	for _, e := range t.edges {
		ed := edge{from: e[0], to: e[1]}
		from := e[0]
		if t.nodeToEdgesMap[from] == nil {
			t.nodeToEdgesMap[from] = make([]*edge, 0)
		}
		t.nodeToEdgesMap[from] = append(t.nodeToEdgesMap[from], &ed)
		to := e[1]
		if t.nodeToEdgesMap[to] == nil {
			t.nodeToEdgesMap[to] = make([]*edge, 0)
		}
		t.nodeToEdgesMap[to] = append(t.nodeToEdgesMap[to], &ed)
	}
}

func (t *tree) findDiameter(startNode int) farthestNode {

	// initialize the edges' access
	for _, edges := range t.nodeToEdgesMap {
		for _, edge := range edges {
			edge.accessed = false
		}
	}

	// nodeDepthMap's key is node, value is the depth from the key to startNode
	nodeDepthMap := make(map[int]int, 0)
	nodeDepthMap[startNode] = 0

	// go through the tree by width
	queue := list.New()
	queue.PushBack(startNode)
	for element := queue.Front(); element != nil; element = queue.Front() {
		n := element.Value.(int)
		queue.Remove(element)
		edges := t.nodeToEdgesMap[n]
		for _, edge := range edges {
			if edge.accessed {
				continue
			}
			from := edge.from
			to := edge.to
			if from == n {
				nodeDepthMap[to] = nodeDepthMap[from] + 1
				queue.PushBack(to)
			} else {
				nodeDepthMap[from] = nodeDepthMap[to] + 1
				queue.PushBack(from)
			}
			edge.accessed = true
		}
	}

	maxWidth := -1
	maxNode := -100
	for n, w := range nodeDepthMap {
		if maxWidth < w {
			maxWidth = w
			maxNode = n
		}
	}

	return farthestNode{
		startNode: startNode,
		to:        maxNode,
		depth:     maxWidth,
	}
}
