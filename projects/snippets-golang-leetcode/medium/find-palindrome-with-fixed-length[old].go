package main

import (
	"fmt"
	"math"
)

/* https://leetcode.com/problems/find-palindrome-with-fixed-length/

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
	queries_map := make(map[int]bool, len(queries))
	for _, query := range queries {
		queries_map[query] = true
	}
	palindromes := constructPalindromes(intLength, queries_map)
	for _, query := range queries {
		palindrom, exists := palindromes[query]
		if !exists {
			palindrom = -1
		}
		answer = append(answer, palindrom)
	}
	return
}

func updateQueries(count *int, length int, palendrom []int, queries map[int]bool, palendromes map[int]int64) {
	*count++
	fmt.Println(*count, palendrom)
	if _, exists := queries[*count]; !exists {
		return
	}
	delete(queries, *count)
	new_palendrom := make([]int, length)
	copy(new_palendrom, palendrom)
	palendromes[*count] = fromList(new_palendrom)

}

func constructPalindromes(length int, queries map[int]bool) map[int]int64 {
	palindromes := make(map[int]int64, len(queries))
	if length == 0 {
		return palindromes
	} else if length == 1 {
		for index := 0; index < 10; index++ {
			palindromes[index] = int64(index)
		}
		return palindromes
	}
	root := math.Pow10(length) / 10
	if root != 1 {
		root += 1
	}
	count_max := int(math.Pow10(1 + length/2))
	middle, is_uneven := length/2, length%2 != 0
	last, count := toList(int(root)), new(int)
	updateQueries(count, length, last, queries, palindromes)
	for {
		switch {
		case last[middle] == 9 && last[0] == 9:
			fallthrough
		case len(queries) == 0 || *count > count_max:
			return palindromes
		case is_uneven && last[middle] < 9:
			last[middle] += 1
			updateQueries(count, length, last, queries, palindromes)
			continue
		case is_uneven && last[middle] == 9:
			last[middle] = 0
		}
		for index := middle - 1; index >= 0; index-- {
			fmt.Println("index", index)
			if last[index] == 9 {
				last[index] = 0
				last[length-index-1] = 0
				continue
			}
			if last[index] < 9 {
				last[index]++
				last[length-index-1]++
				updateQueries(count, length, last, queries, palindromes)
				if len(queries) == 0 || *count > count_max {
					return palindromes
				}
			}
			break
		}
	}
}

func toList(number int) (result []int) {
	if number/10 == 0 {
		return []int{number}
	}
	current := number
	for {
		if current == 0 {
			return
		}
		result = append(result, int(current%10))
		current /= 10
	}
}

func fromList(number []int) (result int64) {
	for index, digit := range number {
		result += int64(digit) * int64(math.Pow10(index))
	}
	return
}

func main() {
	kthPalindrome([]int{200}, 5)
	//kthPalindrome([]int{696771750, 62, 47, 14, 17, 192356691, 209793716, 23, 220935614, 447911113, 5, 4, 72}, 4)
}
