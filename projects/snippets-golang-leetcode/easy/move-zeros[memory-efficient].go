/* https://leetcode.com/problems/move-zeroes/

Runtime: 26 ms, faster than 62.90% of Go online submissions for Move Zeroes.
Memory Usage: 6.6 MB, less than 72.70% of Go online submissions for Move Zeroes.

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:

Input: nums = [0]
Output: [0]



Constraints:

    1 <= nums.length <= 104
    -231 <= nums[i] <= 231 - 1

*/
package main

func moveZeroes(numbers []int) {
	// ensure there are at least two numbers
	length := len(numbers)
	if length == 1 {
		return
	}
	result, index_result := make([]int, length), 0
	for _, number := range numbers {
		if number == 0 {
			continue
		}
		result[index_result] = number
		index_result++
	}
	copy(numbers, result)
}
