package leet105

import (
	"fmt"
)

/*
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder,  0, len(preorder)-1, 0, len(inorder)-1)
}

func build(preorder []int, inorder []int, l1 int, r1 int, l2 int, r2 int) *TreeNode {
	if l1 > r1 {
		return nil
	}
	v := preorder[l1]
	root := &TreeNode{
		Val: v,
	}
	mid := l2
	for inorder[mid] != root.Val {
		mid++
	}

	root.Left = build(preorder, inorder, l1+1, l1+(mid-l2), l2, mid-1)
	root.Right = build(preorder, inorder, l1+(mid-l2)+1, r1, mid+1, r2)
	return root
}
