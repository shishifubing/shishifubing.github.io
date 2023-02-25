/* https://leetcode.com/problems/sqrtx/

Runtime: 2 ms, faster than 65.69% of Go online submissions for Sqrt(x).
Memory Usage: 2 MB, less than 67.28% of Go online submissions for Sqrt(x).

Given a non-negative integer x, compute and return the square root of x.

Since the return type is an integer, the decimal digits are truncated, and only the integer part of the result is returned.

Note: You are not allowed to use any built-in exponent function or operator, such as pow(x, 0.5) or x ** 0.5.



Example 1:

Input: x = 4
Output: 2

Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since the decimal part is truncated, 2 is returned.



Constraints:

    0 <= x <= 231 - 1


*/
package main

func mySqrt(number int) int {
	left, right := 0, number
	for right >= left {
		current := left + (right-left)/2
		square_current := current * current
		square_next := (current + 1) * (current + 1)
		switch {
		case square_current <= number && square_next > number:
			// found the target
			return current
		case square_current > number:
			// target is to the left -> discard right
			right = current - 1
		case square_current < number:
			// target is to the right -> discard left
			left = current + 1
		}
	}
	return -1
}
