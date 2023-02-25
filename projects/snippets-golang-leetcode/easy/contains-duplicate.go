/* https://leetcode.com/problems/contains-duplicate/

Runtime: 75 ms, faster than 78.19% of Go online submissions for Contains Duplicate.
Memory Usage: 9 MB, less than 52.00% of Go online submissions for Contains Duplicate.

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.



Example 1:

Input: nums = [1,2,3,1]
Output: true

Example 2:

Input: nums = [1,2,3,4]
Output: false

Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true



Constraints:

    1 <= nums.length <= 105
    -109 <= nums[i] <= 109

*/

package main

func containsDuplicate(numbers []int) bool {
	if len(numbers) == 0 || len(numbers) == 1 {
		return false
	}
	occured := make(map[int]bool)
	for _, number := range numbers {
		_, isDuplicate := occured[number]
		if isDuplicate {
			return true
		}
		occured[number] = true
	}
	return false
}
