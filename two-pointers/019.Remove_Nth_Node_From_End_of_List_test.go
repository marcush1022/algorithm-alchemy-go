// package twoPointers
// dir_path twoPointers
// name 019.Remove_Nth_Node_From_End_of_List_test

package two_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	assert.Equal(t, "1235", PrintList(removeNthFromEnd(head, 2)))

	head = nil
	head = &ListNode{Val: 1}
	assert.Equal(t, "", PrintList(removeNthFromEnd(head, 1)))

	head = nil
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	assert.Equal(t, "1", PrintList(removeNthFromEnd(head, 1)))

	head = nil
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	assert.Equal(t, "2", PrintList(removeNthFromEnd(head, 2)))
}
