/* https://leetcode.com/problems/reverse-words-in-a-string-iii/

Runtime: 5 ms, faster than 81.86% of Go online submissions for Reverse Words in a String III.
Memory Usage: 6.4 MB, less than 87.50% of Go online submissions for Reverse Words in a String III.

Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.



Example 1:

Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"

Example 2:

Input: s = "God Ding"
Output: "doG gniD"



Constraints:

    1 <= s.length <= 5 * 104
    s contains printable ASCII characters.
    s does not contain any leading or trailing spaces.
    There is at least one word in s.
    All the words in s are separated by a single space.

*/
package main

import "fmt"

func reverseWords(_string string) string {
	index_word_start, length, result := 0, len(_string), []rune(_string)
	for index, character := range result {
		// ignore normal characters
		if character != ' ' {
			continue
		}
		// a word has ended -> reverse it
		reverse_word(result, length, index_word_start, index)
		// moving the index to the start of the next word
		index_word_start = index + 1
	}
	reverse_word(result, length, index_word_start, length)
	return string(result)
}

func reverse_word(_string []rune, length int, start int, end int) {
	fmt.Println("reversing", string(_string), start, "->", end)
	length_word := end - start
	for index := start; index < start+length_word/2; index++ {
		index_last := end - (index - start) - 1
		current, last := _string[index], _string[index_last]
		_string[index], _string[index_last] = last, current
	}
	fmt.Println("result", string(_string), start, "->", end)
}
