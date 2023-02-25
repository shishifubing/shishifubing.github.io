/* https://leetcode.com/problems/sign-of-the-product-of-an-array/

Runtime: 3 ms, faster than 92.28% of Go online submissions for Sign of the Product of an Array.
Memory Usage: 3.3 MB, less than 75.61% of Go online submissions for Sign of the Product of an Array.

There is a function signFunc(x) that returns:

    1 if x is positive.
    -1 if x is negative.
    0 if x is equal to 0.

You are given an integer array nums. Let product be the product of all values in the array nums.

Return signFunc(product).



Example 1:

Input: nums = [-1,-2,-3,-4,3,2,1]
Output: 1
Explanation: The product of all values in the array is 144, and signFunc(144) = 1

Example 2:

Input: nums = [1,5,0,2,-3]
Output: 0
Explanation: The product of all values in the array is 0, and signFunc(0) = 0

Example 3:

Input: nums = [-1,1,-1,1,-1]
Output: -1
Explanation: The product of all values in the array is -1, and signFunc(-1) = -1



Constraints:

    1 <= nums.length <= 1000
    -100 <= nums[i] <= 100


*/
package main

func arraySign(numbers []int) int {
	negative_count := 0
	for _, number := range numbers {
		switch {
		// the number is 0 -> product of all numbers is definitely zero
		case number == 0:
			return 0
			// the number is negative -> add to count
		case number < 0:
			negative_count++
		}
	}
	// even amount of negative numbers -> result is positive
	if negative_count&1 == 0 {
		return 1
	}
	// uneven amount of negative numbers -> result is negative
	return -1
}
