// package bitwiseOperation
// dir_path bitwiseOperation
// name 137.Single_Number_II

package bitwise_operation

/**
Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.

You must implement a solution with a linear runtime complexity and use only constant extra space.



Example 1:

Input: nums = [2,2,3,2]
Output: 3
Example 2:

Input: nums = [0,1,0,1,0,1,99]
Output: 99


Constraints:

1 <= nums.length <= 3 * 104
-231 <= nums[i] <= 231 - 1
Each element in nums appears exactly three times except for one element which appears once.
*/

/**
解法：位操作。不管非孤异元素重复多少次，这是通用做法。

对于右数第i位，如果孤异元素该位为0，则该位为1的元素总数为3的整数倍。

如果孤异元素该位为1，则该位为1的元素总数不为3的整数倍（也就是余1）。

换句话说，如果第i位为1的元素总数不为3的整数倍，则孤异数的第i位为1，否则为0.

（如果非孤异元素重复n次，则判断是否为n的整数倍）
*/

// [2, 2, 3, 2]
func singleNumber2(nums []int) int {
	mask := 1
	result := 0
	for mask != 0 {
		countOf1 := 0
		for i := 0; i < len(nums); i++ {
			if nums[i]&mask != 0 {
				countOf1++
			}
		}

		if countOf1%3 == 1 {
			result |= mask
		}
		mask <<= 1
	}
	return result
}
