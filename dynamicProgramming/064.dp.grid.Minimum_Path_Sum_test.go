// package dynamic_programming
// dir_path dynamic_programming
// name 064.dp.Minimum_Path_Sum_test

package dynamicProgramming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minPathSum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	assert.Equal(t, 7, minPathSum(grid))
}
