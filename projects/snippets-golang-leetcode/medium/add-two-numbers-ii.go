/* https://leetcode.com/problems/add-two-numbers-ii/
You are given two non-empty linked lists representing two non-negative
integers. The most significant digit comes first and each of their nodes
contains a single digit. Add the two numbers and return the sum as a linked
list.

You may assume the two numbers do not contain any leading zero, except the
number 0 itself.
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(list *ListNode) *ListNode {
	if list == nil || list.Next == nil {
		return list
	}
	result := &ListNode{list.Val, nil}
	current := list.Next
	for {
		if current == nil {
			return result
		}
		result = &ListNode{current.Val, result}
		current = current.Next
	}
}

func addTwoNumbers(list_1 *ListNode, list_2 *ListNode) *ListNode {
	current_1, current_2, result := reverse(list_1), reverse(list_2), &ListNode{}
	var carry int
	result_current := result
	for {
		var value_1, value_2 int
		if current_1 != nil {
			value_1 = current_1.Val
			current_1 = current_1.Next
		}
		if current_2 != nil {
			value_2 = current_2.Val
			current_2 = current_2.Next
		}
		sum := value_1 + value_2 + carry
		if sum > 9 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		result_current.Val = sum
		if current_1 == nil && current_2 == nil && carry == 0 {
			result_current.Next = nil
			return reverse(result)
		} else if current_1 == nil && current_2 == nil && carry != 0 {
			current_1 = &ListNode{}
		}
		result_current.Next = &ListNode{}
		result_current = result_current.Next
	}
}

func main() {
	addTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}})
}
