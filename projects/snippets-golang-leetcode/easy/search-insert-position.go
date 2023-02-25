/* https://leetcode.com/problems/search-insert-position/

Runtime: 3 ms, faster than 70.38% of Go online submissions for Search Insert Position.
Memory Usage: 3 MB, less than 5.87% of Go online submissions for Search Insert Position.

Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4



Constraints:

    1 <= nums.length <= 104
    -104 <= nums[i] <= 104
    nums contains distinct values sorted in ascending order.
    -104 <= target <= 104

*/
package main

import "fmt"

func searchInsert(numbers []int, target int) int {
	// checking edge cases
	length := len(numbers)
	if numbers[0] == target {
		return 0
	} else if numbers[length-1] == target {
		return length - 1
	}
	left, right, index, was_bigger := 0, length-1, 0, false
	for right >= left {
		// overflow protection
		index = left + (right-left)/2
		number := numbers[index]
		fmt.Println("index", index, "left", left, "right", right)
		switch {
		case number == target:
			// found the target
			return index
		case number > target:
			// the number is bigger -> the target is to the left -> discard right
			right = index - 1
			was_bigger = true
		case number < target:
			// the number is smaller -> the target is to the right -> discard left
			left = index + 1
			was_bigger = false
		}
	}
	// the target is not in the array

	// the last number was bigger -> target should be to the left
	if was_bigger {
		return index
	}
	// the last number was smaller -> target should be to the right
	return index + 1
}
