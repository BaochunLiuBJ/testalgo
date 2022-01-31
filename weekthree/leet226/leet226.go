package leet226 

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil 
	}
	l := root.Left
	r := root.Right
	root.Left = r 
	root.Right = l 
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
