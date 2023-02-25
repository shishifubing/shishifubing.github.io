/* https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/

Runtime: 3 ms, faster than 49.12% of Go online submissions for Count Odd Numbers in an Interval Range.
Memory Usage: 1.9 MB, less than 21.65% of Go online submissions for Count Odd Numbers in an Interval Range.

Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).



Example 1:

Input: low = 3, high = 7
Output: 3
Explanation: The odd numbers between 3 and 7 are [3,5,7].

Example 2:

Input: low = 8, high = 10
Output: 1
Explanation: The odd numbers between 8 and 10 are [9].



Constraints:

    0 <= low <= high <= 10^9

*/
package main

func countOdds(low int, high int) int {
	low_even, high_even, half := (low&1) == 0, (high&1) == 0, (high-low)/2
	if low_even && high_even {
		return half
	}
	return half + 1
}
