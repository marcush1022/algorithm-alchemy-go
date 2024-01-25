// package twoPointers
// dir_path twoPointers
// name 018.4Sum_test

package twoPointers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fourSum(t *testing.T) {
	assert.Equal(t, "[[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]", fmt.Sprint(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)))
	assert.Equal(t, "[[2 2 2 2]]", fmt.Sprint(fourSum([]int{2, 2, 2, 2, 2}, 8)))
	fmt.Println(fourSum([]int{5, 5, 3, 5, 1, -5, 1, -2}, 4))
}
