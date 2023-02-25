/*  https://leetcode.com/problems/merge-two-sorted-lists/

Runtime: 3 ms, faster than 52.58% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.8 MB, less than 11.05% of Go online submissions for Merge Two Sorted Lists.

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.



Example 1:

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:

Input: list1 = [], list2 = []
Output: []

Example 3:

Input: list1 = [], list2 = [0]
Output: [0]



Constraints:

    The number of nodes in both lists is in the range [0, 50].
    -100 <= Node.val <= 100
    Both list1 and list2 are sorted in non-decreasing order.

*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// ensure both lists are valid
	switch {
	case list1 == nil && list2 == nil:
		return nil
	case list1 == nil && list2 != nil:
		return list2
	case list1 != nil && list2 == nil:
		return list1
	}
	var root *ListNode
	if list1.Val < list2.Val {
		root = &ListNode{list1.Val, nil}
		list1 = list1.Next
	} else {
		root = &ListNode{list2.Val, nil}
		list2 = list2.Next
	}
	current := &root
	for {
		switch {
		case list1 == nil && list2 == nil:
			return root
		case list1 != nil && list2 != nil && list1.Val <= list2.Val:
			fallthrough
		case list1 != nil && list2 == nil:
			fmt.Println("1", list1.Val)
			updateResult(&current, &list1)
		case list1 != nil && list2 != nil && list1.Val > list2.Val:
			fallthrough
		case list1 == nil && list2 != nil:
			fmt.Println("2", list2.Val)
			updateResult(&current, &list2)
		}
	}
}
func updateResult(current ***ListNode, node **ListNode) {
	// modify the current result node
	(**current).Next = &ListNode{(*node).Val, nil}
	// move the current pointer
	*current = &(**current).Next
	// there is no next node -> nill it
	// there is next -> move it
	if (*node).Next == nil {
		*node = nil
	} else {
		*node = (*node).Next
	}
}

func main() {
	result := mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}})
	fmt.Println(result)
}
