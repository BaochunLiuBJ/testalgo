package leet684

type graph struct {
	toMap    map[int][]int
	visited  map[int]bool
	hasCycle bool
}

func findRedundantConnection(edges [][]int) []int {
	return findEdge(edges) 
}

func findEdge(edges [][]int) []int {
	g := &graph{}
	g.toMap = make(map[int][]int)
	g.visited = make(map[int]bool)

	for _, ed := range edges {
		from := ed[0]
		to := ed[1]
		s, ok := g.toMap[from]
		if !ok {
			s = make([]int, 0)
		}
		g.toMap[from] = append(s, to)

		s, ok = g.toMap[to]
		if !ok {
			s = make([]int, 0)
		}
		g.toMap[to] = append(s, from)

		g.hasCycle = false
		for k, _ := range g.visited {
			g.visited[k] = false 
		}
		g.dfs(from, 0)
		if g.hasCycle {
			return ed 
		}
	}
	return []int {} 
}

func (g *graph) dfs(x int, fa int) {
	g.visited[x] = true
	for _, y := range g.toMap[x] {
		if y == fa {
			continue
		}
		if !g.visited[y] {
			g.dfs(y, x)
		} else {
			g.hasCycle = true 
		}
	}
}
