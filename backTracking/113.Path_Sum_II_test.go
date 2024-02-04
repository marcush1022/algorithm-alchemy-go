// package uncategorized
// dir_path uncategorized
// name 113.Path_Sum_II_test

package backTracking

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// [5,4,8,11,null,13,4,7,2,null,null,5,1]

func Test_pathSum(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2}},
			Right: nil},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1}}}}

	assert.Equal(t, "[[5 4 11 2] [5 8 4 5]]", fmt.Sprint(pathSum(root, 22)))
}
