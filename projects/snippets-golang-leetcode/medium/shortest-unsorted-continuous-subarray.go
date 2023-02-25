/* https://leetcode.com/problems/shortest-unsorted-continuous-subarray/

Runtime: 159 ms, faster than 6.06% of Go online submissions for Shortest Unsorted Continuous Subarray.
Memory Usage: 6.7 MB, less than 37.88% of Go online submissions for Shortest Unsorted Continuous Subarray.

Given an integer array nums, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order.

Return the shortest such subarray and output its length.



Example 1:

Input: nums = [2,6,4,8,10,9,15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.

Example 2:

Input: nums = [1,2,3,4]
Output: 0

Example 3:

Input: nums = [1]
Output: 0

*/
package main

import (
	"fmt"
	"math"
)

func findUnsortedSubarray(numbers []int) int {
	length := len(numbers)
	// checking edge cases
	if length == 1 {
		return 0
	}
	result := make([]int, length)
	// initializing to MaxInt because zeros will impact future sorting
	for index := range result {
		result[index] = math.MaxInt
	}
	index_sort_start, index_sort_end := length, length
	result[0] = numbers[0]
	for index := 1; index < length; index++ {
		fmt.Println(result, index_sort_start, "->", index_sort_end)
		// current is not sorted -> from numbers
		// previous is sorted -> from result
		current, previous := numbers[index], result[index-1]
		// just push the current number in the result array because if it is
		// bigger
		if current >= previous {
			result[index] = current
			continue
		}
		// current number is not sorted -> moving the end index
		index_sort_end = index + 1
		// current < previous -> find the number bigger than the current
		//
		// the current number is smaller than the starting number of the unsorted subarray
		// (or the starting index is not set)
		// -> start index is invalid, finding new start index in numbers
		if index_sort_start == length || current < result[index_sort_start] {
			// insert the number in the correct place and move index_sort_start
			// to the correct place
			if index_sort_start == length {
				// index_sort_start is not set -> setting to the last sorted number
				index_sort_start = index - 1
			}
			index_sort_start = insert_smaller(result, current, index_sort_start, 0)
			continue
		}
		// the current number is bigger than the starting number of the unsorted
		// subarray, so it should be placed inside of it
		insert_smaller(result, current, index, index_sort_start)
		// set start index (if not already set)
		if index < index_sort_start {
			index_sort_start = index - 1
		}
	}
	fmt.Println(result, index_sort_start, "->", index_sort_end)
	// return length of the sorted subarray
	return index_sort_end - index_sort_start
}

func insert_smaller(numbers []int, target int, index_start int, index_end int) int {
	index_result := index_end
	for index := index_start; index >= index_end; index-- {
		if target >= numbers[index] {
			// target is bigger -> it should be placed after this index
			index_result = index + 1
			break
		}
	}
	fmt.Println("index for", target, "-", index_result)
	// target should be placed after index_result, so move everything after
	// it to the right and insert the current number
	copy(numbers[index_result+1:], numbers[index_result:])
	numbers[index_result] = target
	return index_result
}
