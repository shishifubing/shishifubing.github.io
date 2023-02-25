/* https://leetcode.com/problems/largest-perimeter-triangle/

Runtime: 57 ms, faster than 30.65% of Go online submissions for Largest Perimeter Triangle.
Memory Usage: 6.9 MB, less than 55.28% of Go online submissions for Largest Perimeter Triangle.

Given an integer array nums, return the largest perimeter of a triangle with a non-zero area, formed from three of these lengths. If it is impossible to form any triangle of a non-zero area, return 0.



Example 1:

Input: nums = [2,1,2]
Output: 5

Example 2:

Input: nums = [1,2,1]
Output: 0



Constraints:

    3 <= nums.length <= 104
    1 <= nums[i] <= 106

*/
package main

import (
	"sort"
)

func largestPerimeter(numbers []int) int {
	sort.Ints(numbers)
	for index := len(numbers) - 1; index > 1; index-- {
		current, sum_previous := numbers[index], numbers[index-1]+numbers[index-2]
		if current >= sum_previous {
			continue
		}
		return current + sum_previous
	}
	return 0
}
