package main

/* https://leetcode.com/problems/two-sum/

Runtime: 62 ms, faster than 8.66% of Go online submissions for Two Sum.
Memory Usage: 3.6 MB, less than 90.43% of Go online submissions for Two Sum.

Given an array of integers nums and an integer target, return indices of the
two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may
not use the same element twice.

You can return the answer in any order.

*/
func twoSum(numbers []int, target int) []int {
	numbers_len := len(numbers)
	for index, number := range numbers {
		for index_inner := index + 1; index_inner < numbers_len; index_inner++ {
			if numbers[index_inner]+number == target {
				return []int{index, index_inner}
			}
		}
	}
	return []int{}
}

func main() {
	twoSum([]int{3, 2, 4}, 6)
}
