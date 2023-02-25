/* https://leetcode.com/problems/valid-perfect-square/

Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Perfect Square.
Memory Usage: 1.9 MB, less than 12.93% of Go online submissions for Valid Perfect Square.

Given a positive integer num, write a function which returns True if num is a perfect square else False.

Follow up: Do not use any built-in library function such as sqrt.



Example 1:

Input: num = 16
Output: true

Example 2:

Input: num = 14
Output: false



Constraints:

    1 <= num <= 2^31 - 1


*/
package main

func isPerfectSquare(number int) bool {
	left, right := 1, number
	for right >= left {
		current := left + (right-left)/2
		square := current * current
		switch {
		case square == number:
			// found the target
			return true
		case square > number:
			// square is bigger -> root is to the left -> discard right
			right = current - 1
		case square < number:
			// square is smaller -> root is to the right -> discard left
			left = current + 1
		}
	}
	return false
}
