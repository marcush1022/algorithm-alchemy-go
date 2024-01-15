// package bitwiseOperation
// dir_path bitwiseOperation
// name 137.Single_Number_II_test

package bitwiseOperation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber2(t *testing.T) {
	assert.Equal(t, "3", fmt.Sprint(singleNumber2([]int{2, 2, 3, 2})))
	assert.Equal(t, "99", fmt.Sprint(singleNumber2([]int{0, 1, 0, 1, 0, 1, 99})))
	assert.Equal(t, "-4", fmt.Sprint(singleNumber2([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2})))
}
