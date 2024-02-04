// package backTracking
// dir_path backTracking
// name common

package backTracking

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RemoveLastElement(slice *[]int) {
	*slice = (*slice)[:len(*slice)-1] // Reassign the pointer to a new slice that excludes the last element
}

func Popback(slice []int) []int {
	// pop back and deep copy
	slice = slice[:len(slice)-1]
	tmp := make([]int, len(slice), len(slice))
	copy(tmp, slice)
	return tmp
}
