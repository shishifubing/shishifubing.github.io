/* https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/

Runtime: 3 ms, faster than 26.38% of Go online submissions for Average Salary Excluding the Minimum and Maximum Salary.
Memory Usage: 2 MB, less than 5.53% of Go online submissions for Average Salary Excluding the Minimum and Maximum Salary.

You are given an array of unique integers salary where salary[i] is the salary of the ith employee.

Return the average salary of employees excluding the minimum and maximum salary. Answers within 10-5 of the actual answer will be accepted.



Example 1:

Input: salary = [4000,3000,1000,2000]
Output: 2500.00000
Explanation: Minimum salary and maximum salary are 1000 and 4000 respectively.
Average salary excluding minimum and maximum salary is (2000+3000) / 2 = 2500

Example 2:

Input: salary = [1000,2000,3000]
Output: 2000.00000
Explanation: Minimum salary and maximum salary are 1000 and 3000 respectively.
Average salary excluding minimum and maximum salary is (2000) / 1 = 2000



Constraints:

    3 <= salary.length <= 100
    1000 <= salary[i] <= 106
    All the integers of salary are unique.

*/
package main

import "math"

func average(salary []int) float64 {
	length := len(salary)
	highest, lowest := math.MinInt, math.MaxInt
	result := 0
	for _, number := range salary {
		result += number
		if number > highest {
			highest = number
		}
		if number < lowest {
			lowest = number
		}
	}
	return float64(result-highest-lowest) / float64(length-2)
}
