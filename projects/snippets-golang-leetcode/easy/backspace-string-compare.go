/* https://leetcode.com/problems/backspace-string-compare/

Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.



Example 1:

Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

Example 2:

Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

Example 3:

Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".



Constraints:

    1 <= s.length, t.length <= 200
    s and t only contain lowercase letters and '#' characters.

*/
package main

import (
	"container/list"
)

func backspaceCompare(string_1 string, string_2 string) bool {
	length_1, length_2 := len(string_1), len(string_2)
	list_1, list_2, length_biggest := list.List{}, list.List{}, length_1
	if length_2 > length_1 {
		length_biggest = length_2
	}
	for index := 0; index < length_biggest; index++ {
		if index < length_1 {
			backspace_action(&list_1, string_1[index])
		}
		if index < length_2 {
			backspace_action(&list_2, string_2[index])
		}
	}
	// lists are not equal, there is no need to check
	if list_1.Len() != list_2.Len() {
		return false
	}
	// checking elements after all deletions
	element_1, element_2 := list_1.Back(), list_2.Back()
	for element_1 != nil {
		if element_1.Value != element_2.Value {
			// elements are not equal -> strings are not equal
			return false
		}
		element_1, element_2 = element_1.Next(), element_2.Next()
	}
	// checked all elements, strings are equal
	return true
}

func backspace_action(list *list.List, character byte) {
	switch {
	case character == '#' && list.Len() != 0:
		// delete last character
		list.Remove(list.Front())
		fallthrough
	case character == '#' && list.Len() == 0:
		// just return if the list is empty
		return
	}
	list.PushFront(character)
}
