// package backtracking
// dir_path backtracking
// name 129.Sum_Root_to_Leaf_Numbers_test

package backtracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumNumbers(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	assert.Equal(t, 25, sumNumbers(root))

	root = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   9,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{Val: 0},
	}
	assert.Equal(t, 1026, sumNumbers(root))
}
