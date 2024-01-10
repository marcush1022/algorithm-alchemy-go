// package dynamic_programming
// dir_path dynamic_programming
// name 300.dp.Longest_Increasing_Subsequence

package dynamicProgramming

/**
Given an unsorted array of integers, find the length of longest increasing subsequence.

For example,
Given [10, 9, 2, 5, 3, 7, 101, 18],
The longest increasing subsequence is [2, 3, 7, 101], therefore the length is 4.
Note that there may be more than one LIS combination, it is only necessary for you to return the length.

Your algorithm should run in O(n2) complexity.

Follow up: Could you improve it to O(n log n) time complexity?
*/

// AI generated solution:
/**
The function uses dynamic programming to calculate the lengths of increasing subsequences for each element in the array.
It initializes a dp array with the first element set to 1. Then, it iterates through the remaining elements and checks if
there are any previous elements that are smaller than the current element. If there are, it updates the dp array with the
maximum length of the increasing subsequence ending at the current element.
*/

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	// Create a dp array to store the lengths of increasing subsequences
	dp := make([]int, length)
	dp[0] = 1
	globalMax := 1

	// Calculate the length of the longest increasing subsequence
	for i := 1; i < length; i++ {
		localMax := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				localMax = max(dp[j], localMax)
			}
		}
		dp[i] = localMax + 1
		globalMax = max(globalMax, dp[i])
	}
	return globalMax
}
