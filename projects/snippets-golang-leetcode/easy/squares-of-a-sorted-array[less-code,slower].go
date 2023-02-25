/* https://leetcode.com/problems/squares-of-a-sorted-array/

Runtime: 60 ms, faster than 20.06% of Go online submissions for Squares of a Sorted Array.
Memory Usage: 7.2 MB, less than 44.97% of Go online submissions for Squares of a Sorted Array.

Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Example 2:

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]



Constraints:

    1 <= nums.length <= 104
    -104 <= nums[i] <= 104
    nums is sorted in non-decreasing order.


Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?

*/
package main

func sortedSquares(numbers []int) []int {
	length := len(numbers)
	switch length {
	case 0:
		return numbers
	case 1:
		return []int{numbers[0] * numbers[0]}
	}
	index_left, index_right := 0, length-1
	results := make([]int, length)
	for index := index_right; index_left <= index_right; index-- {
		left_square := numbers[index_left] * numbers[index_left]
		right_square := numbers[index_right] * numbers[index_right]
		if right_square > left_square {
			results[index] = right_square
			index_right--
			continue
		}
		results[index] = left_square
		index_left++
	}
	return results
}
