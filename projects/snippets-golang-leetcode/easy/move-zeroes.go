/* https://leetcode.com/problems/move-zeroes/

Runtime: 20 ms, faster than 88.31% of Go online submissions for Move Zeroes.
Memory Usage: 7.5 MB, less than 13.78% of Go online submissions for Move Zeroes.

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
	index_zero := -1
	// sliding window algorithm
	for index, number := range numbers {
		// if index_zero is not set, then the first zero is index_zero
		if index_zero == -1 && number == 0 {
			index_zero = index
		}
		// there is no need to move numbers if there were no zeros before
		// there is no need to move zeros
		if index_zero == -1 || number == 0 {
			continue
		}
		// after some zeros we encounter a non-zero number
		// moving that number to the beginning of zeros
		numbers[index_zero] = number
		// current number becomes zero
		numbers[index] = 0
		// moving the index
		index_zero++
	}
}
