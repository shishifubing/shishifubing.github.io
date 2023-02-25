/* https://leetcode.com/problems/first-bad-version/

Runtime: 0 ms, faster than 100.00% of Go online submissions for First Bad Version.
Memory Usage: 1.9 MB, less than 83.88% of Go online submissions for First Bad Version.

You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.



Example 1:

Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.

Example 2:

Input: n = 1, bad = 1
Output: 1



Constraints:

    1 <= bad <= n <= 231 - 1

*/
package main

import "math"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	// checking edge cases
	if isBadVersion(1) {
		return 1
	}
	left, right, bad_version := 1, n, math.MaxInt
	for right >= left {
		// overflow protection
		version := left + (right-left)/2
		switch isBadVersion(version) {
		case false:
			// it is good -> versions to the left are good -> discard left
			left = version + 1
		case true:
			// it is bad -> the first bad version is either this one or to the left
			// discard right
			bad_version = version
			right = version - 1
		}
	}
	// the search space is empty
	return bad_version
}
