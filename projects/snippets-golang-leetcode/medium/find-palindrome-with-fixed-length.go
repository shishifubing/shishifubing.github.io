package main

import (
	"fmt"
	"math"
)

/* https://leetcode.com/problems/find-palindrome-with-fixed-length/

Runtime: 151 ms, faster than 33.82% of Go online submissions for Find Palindrome With Fixed Length.
Memory Usage: 9.5 MB, less than 45.59% of Go online submissions for Find Palindrome With Fixed Length.

Given an integer array queries and a positive integer intLength, return an
array answer where answer[i] is either the queries[i]th smallest positive
palindrome of length intLength or -1 if no such palindrome exists.

A palindrome is a number that reads the same backwards and forwards.
Palindromes cannot have leading zeros.

Example 1:

Input: queries = [1,2,3,4,5,90], intLength = 3
Output: [101,111,121,131,141,999]
Explanation:
The first few palindromes of length 3 are:
101, 111, 121, 131, 141, 151, 161, 171, 181, 191, 202, ...
The 90th palindrome of length 3 is 999.

Example 2:

Input: queries = [2,4,6], intLength = 4
Output: [1111,1331,1551]
Explanation:
The first six palindromes of length 4 are:
1001, 1111, 1221, 1331, 1441, and 1551.



Constraints:

    1 <= queries.length <= 5 * 104
    1 <= queries[i] <= 109
    1 <= intLength <= 15

*/

func kthPalindrome(queries []int, intLength int) (answer []int64) {
	for _, query := range queries {
		answer = append(answer, getPalindrom(query, intLength))
	}
	return
}

func getPalindrom(query int, length int) (result int64) {
	is_even := (length & 1) == 0
	power := length / 2
	if is_even {
		power -= 1
	}
	palindrome := int64(math.Pow10(power)) + int64(query) - 1
	result = palindrome
	if !is_even {
		palindrome /= 10
	}
	for palindrome > 0 {
		result = result*10 + palindrome%10
		palindrome /= 10
	}
	if len(fmt.Sprint(result)) != length {
		return -1
	}
	return
}

func main() {
	kthPalindrome([]int{200}, 5)
	//kthPalindrome([]int{696771750, 62, 47, 14, 17, 192356691, 209793716, 23, 220935614, 447911113, 5, 4, 72}, 4)
}
