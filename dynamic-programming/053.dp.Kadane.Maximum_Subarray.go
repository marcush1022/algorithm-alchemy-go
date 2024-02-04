// package dynamtic_programming
// dir_path dynamtic_programming
// name 053.dp.Maximum_Subarray

package dynamic_programming

/**
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.

click to show more practice.

More practice:
If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

/**
Kadane's algorithm:
local_maximum at index i is the maximum of A[i] and the sum of A[i] and local_maximum at index i-1.
local_maximum[i] = max(A[i], A[i] + local_maximum[i - 1])
*/

func maxSubArray(nums []int) int {
	globalMax := nums[0]
	localMax := nums[0]
	for i := 1; i < len(nums); i++ {
		localMax = max(nums[i], localMax+nums[i])
		globalMax = max(globalMax, localMax)
	}
	return globalMax
}
