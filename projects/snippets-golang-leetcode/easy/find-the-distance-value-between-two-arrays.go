/* https://leetcode.com/problems/find-the-distance-value-between-two-arrays/

Runtime: 5 ms, faster than 82.20% of Go online submissions for Find the Distance Value Between Two Arrays.
Memory Usage: 3.9 MB, less than 11.02% of Go online submissions for Find the Distance Value Between Two Arrays.

Given two integer arrays arr1 and arr2, and the integer d, return the distance value between the two arrays.

The distance value is defined as the number of elements arr1[i] such that there is not any element arr2[j] where |arr1[i]-arr2[j]| <= d.



Example 1:

Input: arr1 = [4,5,8], arr2 = [10,9,1,8], d = 2
Output: 2
Explanation:
For arr1[0]=4 we have:
|4-10|=6 > d=2
|4-9|=5 > d=2
|4-1|=3 > d=2
|4-8|=4 > d=2
For arr1[1]=5 we have:
|5-10|=5 > d=2
|5-9|=4 > d=2
|5-1|=4 > d=2
|5-8|=3 > d=2
For arr1[2]=8 we have:
|8-10|=2 <= d=2
|8-9|=1 <= d=2
|8-1|=7 > d=2
|8-8|=0 <= d=2

Example 2:

Input: arr1 = [1,4,2,3], arr2 = [-4,-3,6,10,20,30], d = 3
Output: 2

Example 3:

Input: arr1 = [2,1,100,3], arr2 = [-5,-2,10,-3,7], d = 6
Output: 1



Constraints:

    1 <= arr1.length, arr2.length <= 500
    -1000 <= arr1[i], arr2[j] <= 1000
    0 <= d <= 100


*/
package main

import (
	"sort"
)

func findTheDistanceValue(array_1 []int, array_2 []int, target int) int {
	// sorting it to use binary search
	sort.Ints(array_2)
	length_2, count := len(array_2), 0
	for _, current_1 := range array_1 {
		left, right, add_to_count := 0, length_2-1, true
		for right >= left {
			index := left + (right-left)/2
			current_2 := array_2[index]
			if abs(current_1, current_2) <= target {
				// current_1 is inside |arr1[i]-arr2[j]| <= d
				// -> ignore it
				add_to_count = false
				break
			}
			switch {
			case current_2 > current_1:
				// current_2 is bigger than current_1
				// -> all numbers to the right are bigger
				// -> discard right, add to count
				right = index - 1
			case current_2 < current_1:
				// current_2 is smaller than current_1
				// -> all numbers to the left are smaller
				// -> discard left, add to count
				left = index + 1
			}
		}
		if add_to_count {
			count++
		}
	}
	return count
}

func abs(number_1 int, number_2 int) int {
	difference := number_1 - number_2
	if difference < 0 {
		return difference * -1
	}
	return difference
}
