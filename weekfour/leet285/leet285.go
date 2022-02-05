package leet285 

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	return getNext(root, p.Val)
}

func getNext(root *TreeNode, val int ) *TreeNode {
	var  ans *TreeNode 
	cur := root 
	
	for  cur != nil {
		if cur.Val == val {
			if cur.Right != nil {
				ans = cur.Right
				for ans.Left != nil {
					ans = ans.Left
				}
			}
			break
		}
		if val < cur.Val  {
			if ans == nil || ans.Val > cur.Val {
				ans = cur 
			}
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return ans 
}
