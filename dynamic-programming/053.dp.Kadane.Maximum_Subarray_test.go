// package dynamic_programming
// dir_path dynamic_programming
// name 053.dp.Maximum_Subarray_test

package dynamic_programming

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	array := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(array))
}
