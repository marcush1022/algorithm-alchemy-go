// package twoPointers
// dir_path twoPointers
// name 015.3Sum_test

package twoPointers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threeSum(t *testing.T) {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	assert.Equal(t, fmt.Sprint([][]int{{-1, -1, 2}, {-1, 0, 1}}), fmt.Sprint(res))
}
