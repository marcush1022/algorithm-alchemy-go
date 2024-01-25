// package golang
// name 015.3_Sum

package twoPointers

import (
	"sort"
)

/**
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k,
and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.



Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.


Constraints:

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	result := make([][]int, 0)
	left, right := 0, 0
	target := 0

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		left = i + 1
		right = len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// skip left duplicate numbers
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// skip right duplicate numbers
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}

		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return result
}
