package leet105 

import (
	"testing"
)
func Test_buildTree(t *testing.T) {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	buildTree(preOrder, inOrder )
}