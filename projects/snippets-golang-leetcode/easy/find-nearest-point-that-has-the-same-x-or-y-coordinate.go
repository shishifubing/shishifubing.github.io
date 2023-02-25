/* https://leetcode.com/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/

Runtime: 301 ms, faster than 5.34% of Go online submissions for Find Nearest Point That Has the Same X or Y Coordinate.
Memory Usage: 7.6 MB, less than 50.49% of Go online submissions for Find Nearest Point That Has the Same X or Y Coordinate.

You are given two integers, x and y, which represent your current location on a Cartesian grid: (x, y). You are also given an array points where each points[i] = [ai, bi] represents that a point exists at (ai, bi). A point is valid if it shares the same x-coordinate or the same y-coordinate as your location.

Return the index (0-indexed) of the valid point with the smallest Manhattan distance from your current location. If there are multiple, return the valid point with the smallest index. If there are no valid points, return -1.

The Manhattan distance between two points (x1, y1) and (x2, y2) is abs(x1 - x2) + abs(y1 - y2).



Example 1:

Input: x = 3, y = 4, points = [[1,2],[3,1],[2,4],[2,3],[4,4]]
Output: 2
Explanation: Of all the points, only [3,1], [2,4] and [4,4] are valid. Of the valid points, [2,4] and [4,4] have the smallest Manhattan distance from your current location, with a distance of 1. [2,4] has the smallest index, so return 2.

Example 2:

Input: x = 3, y = 4, points = [[3,4]]
Output: 0
Explanation: The answer is allowed to be on the same location as your current location.

Example 3:

Input: x = 3, y = 4, points = [[2,3]]
Output: -1
Explanation: There are no valid points.



Constraints:

    1 <= points.length <= 104
    points[i].length == 2
    1 <= x, y, ai, bi <= 104


*/
package main

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	point_target, smallest_distance, smallest_index := []int{x, y}, math.MaxInt, -1
	for index, point := range points {
		// ignore the point if it is not valid
		if !valid(point_target, point) {
			continue
		}
		distance := distance(point_target, point)
		// ignore the point if its Manhattan distance is bigger
		if distance >= smallest_distance {
			continue
		}
		// the distance is smaller, update the index and the distance
		smallest_index, smallest_distance = index, distance
	}
	// there are no valid points
	if smallest_index == -1 {
		return -1
	}
	return smallest_index
}

func valid(point_target []int, point []int) bool {
	switch {
	case point_target[0] == point[0]:
		fallthrough
	case point_target[1] == point[1]:
		return true
	default:
		return false
	}
}

func distance(point_target []int, point []int) int {
	return abs(point_target[0]-point[0]) + abs(point_target[1]-point[1])
}

func abs(number int) int {
	if number < 0 {
		return number * -1
	}
	return number
}
