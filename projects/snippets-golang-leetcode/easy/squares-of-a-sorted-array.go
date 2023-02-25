/* https://leetcode.com/problems/squares-of-a-sorted-array/

Runtime: 26 ms, faster than 86.19% of Go online submissions for Squares of a Sorted Array.
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

import "fmt"

func sortedSquares(numbers []int) []int {
	length := len(numbers)
	if length == 0 || length == 1 {
		return square(numbers, false)
	}
	negativesIndex := -1
	for index, number := range numbers {
		if number >= 0 {
			negativesIndex = index
			break
		}
	}
	if negativesIndex == 0 || negativesIndex == -1 {
		return square(numbers, negativesIndex == -1)
	}
	fmt.Println("negativesIndex", negativesIndex)
	result, resultIndex, negativesIndex := make([]int, length), 0, negativesIndex-1
	for positivesIndex := negativesIndex + 1; resultIndex < length; positivesIndex++ {
		positiveOverflow := positivesIndex >= length
		for {
			if negativesIndex < 0 || resultIndex >= length {
				break
			}
			if !positiveOverflow && numbers[negativesIndex]*-1 > numbers[positivesIndex] {
				break
			}
			result[resultIndex] = numbers[negativesIndex]
			resultIndex++
			fmt.Println("negative", negativesIndex, resultIndex, numbers[negativesIndex], result, numbers)
			negativesIndex--
		}
		if resultIndex < length && !positiveOverflow {
			fmt.Println("positivesIndex", positivesIndex, resultIndex, numbers[positivesIndex], result, numbers)
			result[resultIndex] = numbers[positivesIndex]
			resultIndex++
		}
	}
	return square(result, false)
}

func square(array []int, reverse bool) []int {
	if reverse {
		length := len(array)
		reversed := make([]int, length)
		for index := length - 1; index >= 0; index-- {
			reversed[length-index-1] = array[index] * array[index]
		}
		return reversed
	} else {
		for index, number := range array {
			array[index] = number * number
		}
		return array
	}
}
