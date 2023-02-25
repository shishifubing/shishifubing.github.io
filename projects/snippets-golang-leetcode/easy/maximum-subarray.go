/* https://leetcode.com/problems/maximum-subarray/

Runtime: 123 ms, faster than 49.08% of Go online submissions for Maximum Subarray.
Memory Usage: 9.5 MB, less than 64.19% of Go online submissions for Maximum Subarray.

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:

Input: nums = [1]
Output: 1

Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23



Constraints:

    1 <= nums.length <= 105
    -104 <= nums[i] <= 104



Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.


*/
package main

import "math"

func maxSubArray(numbers []int) int {
	length := len(numbers)
	switch length {
	case 0:
		return 0
	case 1:
		return numbers[0]
	}
	maxCurrent, maxOverall := 0, math.MinInt
	for _, number := range numbers {
		maxCurrent += number
		if maxCurrent > maxOverall {
			maxOverall = maxCurrent
		}
		if maxCurrent < 0 {
			maxCurrent = 0
		}
	}
	return maxOverall
}
