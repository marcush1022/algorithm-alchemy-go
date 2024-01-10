// package bitOperation
// dir_path bitOperation
// name 136.Single_Number_test

package bitwiseOperation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	assert.Equal(t, 1, singleNumber([]int{2, 2, 1}))
	assert.Equal(t, 4, singleNumber([]int{4, 1, 2, 1, 2}))
	assert.Equal(t, 1, singleNumber([]int{1}))
}
