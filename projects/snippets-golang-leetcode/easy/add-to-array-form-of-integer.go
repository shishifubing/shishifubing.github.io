/* https://leetcode.com/problems/add-to-array-form-of-integer/

Runtime: 64 ms, faster than 11.96% of Go online submissions for Add to Array-Form of Integer.
Memory Usage: 7 MB, less than 53.26% of Go online submissions for Add to Array-Form of Integer.

The array-form of an integer num is an array representing its digits in left to right order.

    For example, for num = 1321, the array form is [1,3,2,1].

Given num, the array-form of an integer, and an integer k, return the array-form of the integer num + k.

Example 1:

Input: num = [1,2,0,0], k = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234

Example 2:

Input: num = [2,7,4], k = 181
Output: [4,5,5]
Explanation: 274 + 181 = 455

Example 3:

Input: num = [2,1,5], k = 806
Output: [1,0,2,1]
Explanation: 215 + 806 = 1021

Constraints:

    1 <= num.length <= 104
    0 <= num[i] <= 9
    num does not contain any leading zeros except for the zero itself.
    1 <= k <= 104

*/

package main

func addToArrayForm(number1 []int, add int) []int {
	if add == 0 {
		return number1
	}
	if len(number1) == 0 {
		return convert(add)
	}
	number2 := convert(add)
	length1, length2, carry := len(number1), len(number2), 0
	hightest := length1
	if length2 > length1 {
		hightest = length2
	}
	result := make([]int, hightest)
	index1, index2, indexResult := length1-1, 0, hightest-1
	for {
		index1Valid, index2Valid := index1 >= 0, index2 < length2
		if !index1Valid && !index2Valid && carry == 0 {
			break
		}
		digit1, digit2 := 0, 0
		if index1Valid {
			digit1 = number1[index1]
			index1--
		}
		if index2Valid {
			digit2 = number2[index2]
			index2++
		}
		digitResult := digit1 + digit2 + carry
		if digitResult > 9 {
			carry = 1
			digitResult -= 10
		} else {
			carry = 0
		}
		if indexResult == -1 {
			result = append(result, 0)
			copy(result[1:], result[0:hightest])
			indexResult = 0
		}
		result[indexResult] = digitResult
		indexResult--
	}
	return result
}

func convert(number int) (result []int) {
	for {
		if number == 0 {
			return
		}
		result = append(result, number%10)
		number /= 10
	}
}
