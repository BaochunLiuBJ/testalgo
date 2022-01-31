package leet236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type solution struct {
	answer *TreeNode
	p *TreeNode
	q *TreeNode
}
type searchResult struct {
	first bool
	second bool 
} 
func (s *solution) dfs(root *TreeNode)*searchResult {
	if root == nil {
		return &searchResult{
			first: false,
			second: false,
		}
	} 

	leftResult := s.dfs(root.Left)
	rightResult := s.dfs(root.Right)	
	result := &searchResult{}
	result.first = leftResult.first || rightResult.first || root.Val == s.p.Val
	result.second = leftResult.second || rightResult.second || root.Val == s.q.Val
	if result.first && result.second && s.answer == nil {
		s.answer = root 
	}
	return result 
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	s := &solution{
		p : p,
		q: q, 
	}
	s.dfs(root)
	return s.answer
}
