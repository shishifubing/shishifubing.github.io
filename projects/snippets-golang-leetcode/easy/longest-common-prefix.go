/* https://leetcode.com/problems/longest-common-prefix/

Write a function to find the longest common prefix string amongst an array of
strings.

If there is no common prefix, return an empty string "".
*/
package main

func substrings(input string) map[string]bool {
	results := make(map[string]bool)
	for index := range input {
		results[input[0:index+1]] = true
	}
	return results
}

func string_in_maps(
	source map[string]bool,
	target map[string]bool,
	results map[string]bool) {
	for string := range source {
		is_valid, exists_in_results := results[string]
		_, exists_in_target := target[string]
		switch {
		case !exists_in_target:
			fallthrough
		case exists_in_target && (is_valid || !exists_in_results):
			results[string] = exists_in_target
		}
	}
}

func longestCommonPrefix(strings []string) (result string) {
	results := make(map[string]bool)
	var previous map[string]bool
	for index, string := range strings {
		current := substrings(string)
		if index == 0 {
			previous = current
		}
		string_in_maps(previous, current, results)
		string_in_maps(current, previous, results)
		previous = current
	}
	for string, valid := range results {
		if valid && len(string) > len(result) {
			result = string
		}
	}
	return
}

func main() {
	longestCommonPrefix([]string{"flower", "flow", "flight"})
	longestCommonPrefix([]string{"dog", "racecar", "car"})
	longestCommonPrefix([]string{"a"})
	longestCommonPrefix([]string{"cir", "car"})
}
