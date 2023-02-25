/* https://leetcode.com/problems/sort-array-by-parity/

Runtime: 9 ms, faster than 60.25% of Go online submissions for Sort Array By Parity.
Memory Usage: 5.1 MB, less than 17.15% of Go online submissions for Sort Array By Parity.

Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.

Return any array that satisfies this condition.



Example 1:

Input: nums = [3,1,2,4]
Output: [2,4,3,1]
Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

Example 2:

Input: nums = [0]
Output: [0]



Constraints:

    1 <= nums.length <= 5000
    0 <= nums[i] <= 5000

*/
package main

func sortArrayByParity(numbers []int) []int {
	length := len(numbers)
	if length == 1 {
		return numbers
	}
	results, index_even, index_odd := make([]int, length), 0, length-1
	for _, number := range numbers {
		if (number & 1) == 0 {
			// number is even -> place it from the beginning
			results[index_even] = number
			index_even++
		} else {
			// number is odd -> place it from the end
			results[index_odd] = number
			index_odd--
		}
	}
	return results
}
