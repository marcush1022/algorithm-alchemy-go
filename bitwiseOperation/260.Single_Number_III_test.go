// package bitwiseOperation
// dir_path bitwiseOperation
// name 260.Single_Number_III_test

package bitwiseOperation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber3(t *testing.T) {
	assert.Equal(t, "[3 5]", fmt.Sprint(singleNumber3([]int{1, 2, 1, 3, 2, 5})))
	assert.Equal(t, "[1 0]", fmt.Sprint(singleNumber3([]int{0, 1})))
	assert.Equal(t, "[-1 0]", fmt.Sprint(singleNumber3([]int{-1, 0})))
}
