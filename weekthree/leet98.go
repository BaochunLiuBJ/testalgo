package weekthree

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := root.Left
	r := root.Right
	v := root.Val

	if l != nil {
		if v <= l.Val {
			return false 
		}
	}
	if r != nil {
		if v >= r.Val {
			return false
		}
	}

	if isValidBST(l) && isValidBST(r) {
		return true
	} else {
		return false
	}
}