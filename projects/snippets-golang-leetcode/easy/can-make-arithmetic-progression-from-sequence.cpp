/* https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/

A sequence of numbers is called an arithmetic progression if the difference
between any two consecutive elements is the same.

Given an array of numbers arr, return true if the array can be rearranged to
form an arithmetic progression. Otherwise, return false.



Example 1:

Input: arr = [3,5,1]
Output: true
Explanation: We can reorder the elements as [1,3,5] or [5,3,1] with differences
2 and -2 respectively, between each consecutive elements.

Example 2:

Input: arr = [1,2,4]
Output: false
Explanation: There is no way to reorder the elements to obtain an arithmetic
progression.



Constraints:

    2 <= arr.length <= 1000
    -106 <= arr[i] <= 106


*/
#include <bits/stdc++.h>

#include <vector>
class Solution {
 public:
  bool canMakeArithmeticProgression(std::vector<int> &array) {
    std::sort(array.begin(), array.end());
    int k = 0;
    for (int index = 0; index < array.size() - 2; index++) {
      if (array[index] - array[index + 1] ==
          array[index + 1] - array[index + 2])
        ++k;
    }
    if (k == array.size() - 2) return true;
    return false;
  }
};