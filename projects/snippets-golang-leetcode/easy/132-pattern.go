/* https://leetcode.com/problems/132-pattern/

Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k] such that i < j < k and nums[i] < nums[k] < nums[j].

Return true if there is a 132 pattern in nums, otherwise, return false.



Example 1:

Input: nums = [1,2,3,4]
Output: false
Explanation: There is no 132 pattern in the sequence.

Example 2:

Input: nums = [3,1,4,2]
Output: true
Explanation: There is a 132 pattern in the sequence: [1, 4, 2].

Example 3:

Input: nums = [-1,3,2,0]
Output: true
Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].



Constraints:

    n == nums.length
    1 <= n <= 2 * 105
    -109 <= nums[i] <= 109

*/
package main

import (
	"container/list"
	"math"
)

func find132pattern(numbers []int) bool {
	length := len(numbers)
	if length < 3 {
		// if the array doesn't have at least three numbers, it cannot have
		// '123' pattern
		return false
	}
	stack, third_element := list.List{}, math.MinInt
	// [3,1,4,2]
	for index := length - 1; index >= 0; index-- {
		current := numbers[index]
		if current < third_element {
			// the stack is taking care of '32' part of the pattern
			// if 'current' is smaller, '132' is formed
			return true
		}
		for stack.Len() != 0 && stack.Front().Value.(int) < current {
			third_element = stack.Remove(stack.Front()).(int)
		}
		stack.PushFront(current)
	}
	return false
}
