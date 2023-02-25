package main

import "sort"

/* https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

Given the root of a Binary Search Tree and a target number k, return true if
there exist two elements in the BST such that their sum is equal to the given
target.

Runtime: 33 ms, faster than 39.76% of Go online submissions for Two Sum IV - Input is a BST.
Memory Usage: 8.1 MB, less than 27.71% of Go online submissions for Two Sum IV - Input is a BST.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func construct(root *TreeNode, results map[int]int) {
	if root == nil {
		return
	}
	if results[root.Val] < 2 {
		results[root.Val] += 1
	}
	construct(root.Left, results)
	construct(root.Right, results)
}

func sorted(input_map map[int]int) (results []int) {
	for key := range input_map {
		results = append(results, key)
	}
	return
}

func findTarget(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	results := map[int]int{}
	construct(root, results)
	if target%2 == 0 && results[target/2] == 2 {
		return true
	}
	results_sorted := sorted(results)
	results_sorted_len := len(results_sorted)
	sort.Ints(results_sorted)
	for index, number := range results_sorted {
		for index_inner := index + 1; index_inner < results_sorted_len; index_inner++ {
			if number+results_sorted[index_inner] == target {
				return true
			}
		}
	}
	return false
}
