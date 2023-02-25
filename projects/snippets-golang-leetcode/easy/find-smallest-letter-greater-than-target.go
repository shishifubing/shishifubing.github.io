/* https://leetcode.com/problems/find-smallest-letter-greater-than-target/

Runtime: 0 ms, faster than 100.00% of Go online submissions for Find Smallest Letter Greater Than Target.
Memory Usage: 2.7 MB, less than 63.75% of Go online submissions for Find Smallest Letter Greater Than Target.

Given a characters array letters that is sorted in non-decreasing order and a character target, return the smallest character in the array that is larger than target.

Note that the letters wrap around.

    For example, if target == 'z' and letters == ['a', 'b'], the answer is 'a'.



Example 1:

Input: letters = ["c","f","j"], target = "a"
Output: "c"

Example 2:

Input: letters = ["c","f","j"], target = "c"
Output: "f"

Example 3:

Input: letters = ["c","f","j"], target = "d"
Output: "f"



Constraints:

    2 <= letters.length <= 104
    letters[i] is a lowercase English letter.
    letters is sorted in non-decreasing order.
    letters contains at least two different characters.
    target is a lowercase English letter.

*/
package main

func nextGreatestLetter(letters []byte, target byte) byte {
	length := len(letters)
	left, right := 0, length-1
	for right >= left {
		index := left + (right-left)/2
		character := letters[index]
		// the character is bigger, we found at least one result
		// smaller characters are to the left -> discard right
		if character > target {
			right = index - 1
			continue
		}
		// character is either equal or smaller -> there is no results to the
		// left -> discard left
		left = index + 1
	}
	return letters[left%length]
}
