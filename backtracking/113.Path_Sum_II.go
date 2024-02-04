// package golang
// name 113.(dfs)Path_Sum_II

package backtracking

/*
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

For example:
Given the below binary tree and sum = 22,
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
return
[
   [5,4,11,2],
   [5,8,4,5]
]

与前一题同样的思路
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	path := make([]int, 0)
	results := &[][]int{}
	searchDFS2(root, targetSum, 0, path, results)
	return *results
}

func searchDFS(node *TreeNode, targetSum int, currentPath *[]int, results *[][]int) {
	if node == nil {
		return
	}

	*currentPath = append(*currentPath, node.Val)
	targetSum -= node.Val

	if targetSum == 0 && node.Left == nil && node.Right == nil {
		path := make([]int, len(*currentPath))
		copy(path, *currentPath)
		*results = append(*results, path)
	} else {
		searchDFS(node.Left, targetSum, currentPath, results)
		searchDFS(node.Right, targetSum, currentPath, results)
	}
	RemoveLastElement(currentPath)
}

/**
The problem with the provided solution is related to the use of a shared reference to the current path across recursive calls in the searchDFS2 function.
When a slice is passed as a pointer and modified within a recursive function, it can lead to unexpected behavior due to the shared reference.

The issue arises from the fact that the same slice reference is being modified and shared across multiple recursive calls.
As a result, when backtracking and removing the last element from the slice using the RemoveLastElement function, it affects the shared slice reference,
potentially leading to incorrect or unexpected results.

To address this issue, each recursive call should work with its own independent copy of the current path to avoid interference from other recursive calls.
This can be achieved by creating a new slice for each recursive call to ensure that modifications to the current path are isolated to the specific call stack.
*/

func searchDFSWithProblem(node *TreeNode, targetSum, currentSum int, currentPath *[]int, results *[][]int) {
	if node == nil {
		return
	}

	*currentPath = append(*currentPath, node.Val)

	if currentSum+node.Val == targetSum && node.Left == nil && node.Right == nil {
		path := make([]int, len(*currentPath))
		copy(path, *currentPath)
		*results = append(*results, path)
		return
	}
	searchDFSWithProblem(node.Left, targetSum, currentSum+node.Val, currentPath, results)
	searchDFSWithProblem(node.Right, targetSum, currentSum+node.Val, currentPath, results)
	RemoveLastElement(currentPath)
}

func searchDFS2(node *TreeNode, targetSum, currentSum int, currentPath []int, results *[][]int) {
	if node == nil {
		return
	}

	currentPath = append(currentPath, node.Val)

	if currentSum+node.Val == targetSum && node.Left == nil && node.Right == nil {
		path := make([]int, len(currentPath))
		copy(path, currentPath)
		*results = append(*results, path)
		return
	}
	searchDFS2(node.Left, targetSum, currentSum+node.Val, append([]int{}, currentPath...), results)
	searchDFS2(node.Right, targetSum, currentSum+node.Val, append([]int{}, currentPath...), results)
}

/**
What's the difference between passing a slice itself to a function and passing a pointer to the slice to a function?

When passing a slice itself to a function, a copy of the slice header is made, but the underlying array is not copied. This means that changes made to the elements of the slice within the function will be visible outside the function, but changes to the slice header itself (such as appending elements) will not be visible outside the function.

When passing a pointer to a slice to a function, the function receives a reference to the original slice, and any modifications made to the slice within the function will be visible outside the function. This includes changes to the slice header as well as changes to the elements of the slice.

Here's a simple example to illustrate the difference:

func modifySlice(slice []int) {
    slice[0] = 100 // Modifying the original slice
    slice = append(slice, 4) // This change is not visible outside the function
}

func modifySliceByReference(slice *[]int) {
    (*slice)[0] = 200 // Modifying the original slice through the pointer
    *slice = append(*slice, 5) // This change is visible outside the function
}

func main() {
    originalSlice := []int{1, 2, 3}

    fmt.Println("Before modification:", originalSlice)

    modifySlice(originalSlice)
    fmt.Println("After modification (passing slice):", originalSlice) // Changes are not visible

    modifySliceByReference(&originalSlice)
    fmt.Println("After modification (passing pointer to slice):", originalSlice) // Changes are visible
}
*/
