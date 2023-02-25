package main

/* https://leetcode.com/problems/valid-parentheses/

Runtime: 3 ms, faster than 30.25% of Go online submissions for Valid Parentheses.
Memory Usage: 2.1 MB, less than 43.83% of Go online submissions for Valid Parentheses.

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.

*/

func isValid(inputString string) bool {
	if len(inputString) == 0 || (len(inputString)&1) != 0 {
		return false
	}
	occurences := []rune{}
	ends := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, character := range inputString {
		length := len(occurences)
		last_valid, is_end := ends[character]
		if (is_end && length == 0) || (is_end && occurences[length-1] != last_valid) {
			return false
		}
		if is_end {
			occurences = occurences[0 : length-1]
		} else {
			occurences = append(occurences, character)
		}
	}
	if len(occurences) > 0 {
		return false
	}
	return true
}

func main() {
	isValid("()")
}
