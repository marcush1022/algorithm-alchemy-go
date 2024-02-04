// package twoPointers
// dir_path twoPointers
// name 19.Remove_Nth_Node_From_End_of_List

package two_pointers

/**
Given the head of a linked list, remove the nth node from the end of the list and return its head.


Example 1:

Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]


Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return nil
	}

	left := head
	right := head
	for i := 0; i < n; i++ {
		right = right.Next
	}

	// if n = length
	if right == nil {
		return head.Next
	}

	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	// Now left is the node in front of the nth node from the end of the list
	left.Next = left.Next.Next
	return head
}
