/* https://leetcode.com/problems/max-number-of-k-sum-pairs/

Runtime: 206 ms, faster than 36.36% of Go online submissions for Max Number of K-Sum Pairs.
Memory Usage: 13 MB, less than 9.09% of Go online submissions for Max Number of K-Sum Pairs.

You are given an integer array nums and an integer k.

In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

Return the maximum number of operations you can perform on the array.



Example 1:

Input: nums = [1,2,3,4], k = 5
Output: 2
Explanation: Starting with nums = [1,2,3,4]:
- Remove numbers 1 and 4, then nums = [2,3]
- Remove numbers 2 and 3, then nums = []
There are no more pairs that sum up to 5, hence a total of 2 operations.

Example 2:

Input: nums = [3,1,3,4,3], k = 6
Output: 1
Explanation: Starting with nums = [3,1,3,4,3]:
- Remove the first two 3's, then nums = [1,4,3]
There are no more pairs that sum up to 6, hence a total of 1 operation.



Constraints:

    1 <= nums.length <= 105
    1 <= nums[i] <= 109
    1 <= k <= 109

*/
package main

func maxOperations(numbers []int, sum int) int {
	count, previous := 0, make(map[int]int, len(numbers))
	for _, current := range numbers {
		target := sum - current
		target_unmatched, target_occured := previous[target]
		// if the current number has not occured before, then it is not in the map
		// -> it needs to be initialized
		if _, current_occured := previous[current]; !current_occured {
			previous[current] = 0
		}
		// number of duplicates of the current number has increased
		previous[current]++
		switch {
		case target_occured && target_unmatched == 0:
			// the target has occured before but there are no unmatched duplicates
			fallthrough
		case !target_occured:
			// in order to get the sum we need the target, but it has not appeared yet
			continue
		case target_occured && target_unmatched > 0:
			// the target has appeared before and there are some unmached duplicates left
			// -> removing the current number and the target from possible matches
			previous[current]--
			previous[target]--
			count++
		}
	}
	return count
}
