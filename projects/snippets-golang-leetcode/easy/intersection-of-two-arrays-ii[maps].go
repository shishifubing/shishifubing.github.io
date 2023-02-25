/* https://leetcode.com/problems/intersection-of-two-arrays-ii/

Runtime: 4 ms, faster than 60.83% of Go online submissions for Intersection of Two Arrays II.
Memory Usage: 3.7 MB, less than 7.89% of Go online submissions for Intersection of Two Arrays II.

Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.



Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]

Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.



Constraints:

    1 <= nums1.length, nums2.length <= 1000
    0 <= nums1[i], nums2[i] <= 1000



Follow up:

    What if the given array is already sorted? How would you optimize your algorithm?
    What if nums1's size is small compared to nums2's size? Which algorithm is better?
    What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?


*/
package main

func intersect(numbers_1 []int, numbers_2 []int) []int {
	length_1, length_2 := len(numbers_1), len(numbers_2)
	length_biggest, result := length_1, make([]int, length_1+length_2)
	if length_2 > length_1 {
		length_biggest = length_2
	}
	occurences := make(map[int][]int, length_biggest)

	for index := 0; index < length_biggest; index++ {
		// if numbers with this index exists, they will be added to the map
		if index < length_1 {
			add_to_occurences(numbers_1[index], 0, occurences)
		}
		if index < length_2 {
			add_to_occurences(numbers_2[index], 1, occurences)
		}
	}
	index := 0
	for number, occurence := range occurences {
		// finding how many times the number should be repeated
		repeat := occurence[0]
		if occurence[1] < repeat {
			repeat = occurence[1]
		}
		// repeatedly inserting the number into the result array
		for ; repeat > 0; repeat-- {
			result[index] = number
			index++
		}
	}
	// returning only intersection
	return result[0:index]
}
func add_to_occurences(number int, index int, occurences map[int][]int) {
	if _, occured := occurences[number]; !occured {
		occurences[number] = make([]int, 2)
	}
	occurences[number][index]++
}
