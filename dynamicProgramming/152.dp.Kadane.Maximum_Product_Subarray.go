// package dynamic_programming
// dir_path dynamic_programming
// name 152.dp.Maximum_Product_Subarray

package dynamicProgramming

/**
Find the contiguous subarray within an array (containing at least one number) which has the largest product.

For example, given the array [2,3,-2,4],
the contiguous subarray [2,3] has the largest product = 6.

max[i]=max[i-1]*num[i]
	OR min[i-1]*num[i]
	OR num[i];
*/

/**
The function uses a modified version of Kadane's algorithm to keep track of both the maximum and minimum product at each position.
This is necessary because a negative number can turn the smallest product into the largest product when multiplied
*/

func maxProductArray(nums []int) int {
	globalMax := nums[0]
	localMax := nums[0]
	localMin := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			localMax, localMin = localMin, localMax
		}

		localMax = max(nums[i], localMax*nums[i])
		localMin = min(nums[i], localMin*nums[i])

		globalMax = max(globalMax, localMax)
	}
	return globalMax
}
