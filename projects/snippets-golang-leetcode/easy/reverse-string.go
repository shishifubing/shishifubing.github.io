/* https://leetcode.com/problems/reverse-string/

Runtime: 38 ms, faster than 54.93% of Go online submissions for Reverse String.
Memory Usage: 6.5 MB, less than 98.99% of Go online submissions for Reverse String.

Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.



Example 1:

Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:

Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]



Constraints:

    1 <= s.length <= 105
    s[i] is a printable ascii character.

*/
package main

func reverseString(characters []byte) {
	length := len(characters)
	for index := 0; index < length/2; index++ {
		index_last := length - index - 1
		current, last := characters[index], characters[index_last]
		characters[index], characters[index_last] = last, current
	}
}
