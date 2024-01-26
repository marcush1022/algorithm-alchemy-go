// package twoPointers
// dir_path twoPointers
// name types

package twoPointers

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) string {
	var res string
	for head != nil {
		res += strconv.Itoa(head.Val)
		head = head.Next
	}
	return res
}
