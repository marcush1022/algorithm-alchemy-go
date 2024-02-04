// package backtracking
// dir_path backtracking
// name 129.Sum_Root_to_Leaf_Numbers

package backtracking

import "fmt"

/**
You are given the root of a binary tree containing digits from 0 to 9 only.

Each root-to-leaf path in the tree represents a number.

For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.

A leaf node is a node with no children.

Example 1:

Input: root = [1,2,3]
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.

Example 2:

Input: root = [4,9,0,5,1]
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.


Constraints:

The number of nodes in the tree is in the range [1, 1000].
0 <= Node.val <= 9
The depth of the tree will not exceed 10.
*/

func searchTreeNodes(node *TreeNode, currentSum int, currentPath []int, results *[][]int, resultsOfNumbers *[]int, resultsSum *int) {
	if node == nil {
		return
	}

	currentPath = append(currentPath, node.Val)
	currentSum = currentSum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		path := make([]int, len(currentPath))
		copy(path, currentPath)
		*results = append(*results, path)
		*resultsOfNumbers = append(*resultsOfNumbers, currentSum)
		*resultsSum += currentSum
		return
	}
	searchTreeNodes(node.Left, currentSum, append([]int{}, currentPath...), results, resultsOfNumbers, resultsSum)
	searchTreeNodes(node.Right, currentSum, append([]int{}, currentPath...), results, resultsOfNumbers, resultsSum)
}

func searchTreeNodes2(node *TreeNode, currentSum int, resultsSum *int) {
	if node == nil {
		return
	}

	currentSum = currentSum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		*resultsSum += currentSum
		return
	}
	searchTreeNodes2(node.Left, currentSum, resultsSum)
	searchTreeNodes2(node.Right, currentSum, resultsSum)
}

func sumNumbers(root *TreeNode) int {
	//path := make([]int, 0)
	//results := &[][]int{}
	//resultsOfNumbers := &[]int{}
	resultsSum := 0
	searchTreeNodes2(root, 0, &resultsSum)
	//fmt.Println(">>>>> results", results)
	//fmt.Println(">>>>> resultsOfNumbers", resultsOfNumbers)
	fmt.Println(">>>>> resultsSum", resultsSum)
	return resultsSum
}
