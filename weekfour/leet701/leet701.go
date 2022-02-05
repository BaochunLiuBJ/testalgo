package leet701

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
			Left: nil,
			Right: nil ,
		}
	} else {
		if val < root.Val {
			root.Left = insertIntoBST(root.Left, val )
		} else {
			root.Right = insertIntoBST(root.Right, val)
		}
		return root 
	}
}
