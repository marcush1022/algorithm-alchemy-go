// package golang
// dir_path golang
// name 001.hashmap.Two_Sum

package uncategorized

/**
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

/**
To find the indices of two numbers in an array that add up to a specific target, you can use a hashmap to store the
complement of each number as you iterate through the array
*/

func twoSum(nums []int, target int) []int {
	// Create a hashmap to store the complement of each number
	complements := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		if index, ok := complements[complement]; ok {
			// Found the complement, return the indices
			return []int{index, i}
		}
		// Store the current number and its index in the hashmap
		complements[v] = i
	}

	// No solution found
	return nil
}
