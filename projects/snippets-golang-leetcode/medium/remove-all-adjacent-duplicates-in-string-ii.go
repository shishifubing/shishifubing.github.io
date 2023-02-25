/* https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/

cannot make it in golang -> if will run out of memory if you use list.List

You are given a string s and an integer k, a k duplicate removal consists of choosing k adjacent and equal letters from s and removing them, causing the left and the right side of the deleted substring to concatenate together.

We repeatedly make k duplicate removals on s until we no longer can.

Return the final string after all such duplicate removals have been made. It is guaranteed that the answer is unique.



Example 1:

Input: s = "abcd", k = 2
Output: "abcd"
Explanation: There's nothing to delete.

Example 2:

Input: s = "deeedbbcccbdaa", k = 3
Output: "aa"
Explanation:
First delete "eee" and "ccc", get "ddbbbdaa"
Then delete "bbb", get "dddaa"
Finally delete "ddd", get "aa"

Example 3:

Input: s = "pbbcggttciiippooaais", k = 2
Output: "ps"



Constraints:

    1 <= s.length <= 105
    2 <= k <= 104
    s only contains lower case English letters.

*/
package main

import (
	"strings"
)

type Character struct {
	character byte
	count     int
}

func removeDuplicates(_string string, duplicate_removal int) string {
	length := len(_string)
	// ensure there are at least two characters
	if length < duplicate_removal || length == 1 {
		return _string
	}
	result, index_result := []Character{{_string[0], 1}}, 0
	result_string_builder := strings.Builder{}
	for index := 1; index < length; index++ {
		current, current_result := _string[index], result[index_result]
		is_equal := current == current_result.character
		if is_equal {
			// it is a duplicate -> increase the count
			current_result.count++
		}
		if current_result.count == duplicate_removal {
			// it reached the duplicate removal threshold -> 'remove' it
			current_result.count = 0
			// moving backwards
			index_result--
		}
		if is_equal {
			continue
		}
		// not a duplicate -> append it
		index_result++
		new_character := Character{current, 1}
		if index_result < len(result) {
			result[index_result] = new_character
			continue
		}
		result = append(result, new_character)
	}
	// build the result string
	for _, character := range result {
		for ; character.count > 0; character.count-- {
			result_string_builder.WriteByte(character.character)
		}
	}
	return result_string_builder.String()
}
