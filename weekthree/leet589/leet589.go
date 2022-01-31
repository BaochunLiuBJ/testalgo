package leet589

type NNode struct {
	Val      int
	Children []*NNode
}

func preorder(root *NNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	result = append(result, root.Val)
	for _, n := range root.Children {
		result = append(result, preorder(n)...)
	}
	return result
}

func preorder2(root *NNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := make([]*NNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		n := len(stack)
		t := stack[n-1]
		stack = stack[:n-1]
		result = append(result, t.Val)
		for i:= len(t.Children)-1; i >=0; i-- {
			stack = append(stack, t.Children[i])
		}
	}
	return result 
}
