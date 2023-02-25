package main

/* https://leetcode.com/problems/palindrome-number/

Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.

For example, 121 is a palindrome while 123 is not.

Runtime: 27 ms, faster than 51.83% of Go online submissions for Palindrome Number.
Memory Usage: 6.5 MB, less than 7.01% of Go online submissions for Palindrome Number.

*/

func isPalindrome(number int) (result bool) {
	if number < 0 {
		return false
	} else if number/10 == 0 {
		return true
	}
	current := number
	digits := []int{}
	for {
		if current == 0 {
			break
		}
		digit := current % 10
		current /= 10
		digits = append(digits, digit)
	}

	digits_length := len(digits)
	skip := -1
	if digits_length%2 != 0 {
		skip = digits_length / 2
	}
	for index, digit := range digits {
		if index != skip && digit != digits[digits_length-index-1] {
			return false
		}
	}
	return true
}

func main() {}
