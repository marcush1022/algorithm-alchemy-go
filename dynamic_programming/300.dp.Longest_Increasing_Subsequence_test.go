// package dynamic_programming
// dir_path dynamic_programming
// name 300.dp.Longest_Increasing_Subsequence_test

package dynamic_programming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	assert.Equal(t, 4, lengthOfLIS(nums))
}
