// package backTracking
// dir_path backTracking
// name common_test

package backtracking

import (
	"fmt"
	"testing"
)

func TestRemoveLastElement(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5}
	RemoveLastElement(&originalSlice)
	fmt.Println(">>>>>", originalSlice)
}
