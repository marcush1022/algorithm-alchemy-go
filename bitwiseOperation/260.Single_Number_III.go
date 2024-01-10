// package bitOperation
// dir_path bitOperation
// name 260.Single_Number_III

package bitwiseOperation

import "fmt"

/**
Given an integer array nums, in which exactly two elements appear only once and all the other elements appear exactly twice.
Find the two elements that appear only once. You can return the answer in any order.

You must write an algorithm that runs in linear runtime complexity and uses only constant extra space.



Example 1:

Input: nums = [1,2,1,3,2,5]
Output: [3,5]
Explanation:  [5, 3] is also a valid answer.
Example 2:

Input: nums = [-1,0]
Output: [-1,0]
Example 3:

Input: nums = [0,1]
Output: [1,0]


Constraints:

2 <= nums.length <= 3 * 104
-231 <= nums[i] <= 231 - 1
*/

/**
How represent negative numbers in binary?
https://cs.calvin.edu/activities/books/rit/chapter5/negative.htm#:~:text=Negative%20Numbers&text=The%20simplest%20is%20to%20simply,would%20be%20written%20as%2011100.

Preliminary Concepts: Negative Binary Numbers
Negative Numbers

We can represent negative numbers in several ways. The simplest is to simply use the leftmost digit of the number as a special value to represent the sign of the number: 0 = positive, 1 = negative.
For example, a value of positive 12 (decimal) would be written as 01100 in binary, but negative 12 (decimal) would be written as 11100.
Notice that in this system, it is important to show the leading 0 (to indicate a positive value).

For technical reasons, a different scheme, called "two's complement" is more often used for representing negative numbers.
In this system, a positive 12 is still 01100, but -12 would be written as 10100. Notice that there is nothing instrinsically correct about one system over another.
Either 11100 or 10100 can be used to represent -12, it just depends on what system of interpretation is used. That is, a human programmer chooses the meaning of the bits.

Answer of GPT:
Negative numbers are typically represented in binary using a system called two's complement. In two's complement representation:
- The leftmost bit (most significant bit) is used to represent the sign of the number. 0 indicates a positive number, and 1 indicates a negative number.
- To obtain the two's complement representation of a negative number, the positive binary representation of the number is inverted (each 0 becomes 1 and each 1 becomes 0) and then 1 is added to the result.

For example, to represent the decimal number -5 in 8-bit binary using two's complement:
1. The binary representation of 5 is 00000101.
2. Invert the bits to get 11111010.
3. Add 1 to the result to obtain 11111011, which is the two's complement representation of -5 in 8-bit binary.

This representation allows for simple addition and subtraction operations to be performed using the same logic for both positive and negative numbers.

Where is 8-bit binary used?
8-bit binary representation is commonly used in various computing systems and architectures. Some common applications of 8-bit binary representation include:
1. Character Encoding: In early computer systems, 8-bit binary representation was used to encode characters in character sets such as ASCII (American Standard Code for Information Interchange) and extended ASCII,
where each character is represented by an 8-bit binary code.
2. Processor Architectures: Many early microprocessors and microcontrollers used 8-bit binary representation as their native word size. Examples include Intel 8080, Zilog Z80, and the original Intel x86 architecture.
3. Graphics and Color Depth: In graphics and image processing, 8-bit binary representation is often used to represent color depth, allowing for 256 different colors to be represented in an image.
4. Embedded Systems: In embedded systems and low-power devices, 8-bit microcontrollers are still widely used due to their simplicity and cost-effectiveness.

While 8-bit binary representation has been largely superseded by larger word sizes in modern computing, it still finds use in specific applications and legacy systems.
*/

// Example: {1, 2, 1, 3, 2, 5}
func singleNumber3(nums []int) []int {
	bitmask := 0
	for _, v := range nums {
		bitmask ^= v
	}
	// bitmask = 00000110
	fmt.Printf(">>>>> bitmask %08b\n", bitmask)
	// The negation of bitmask is calculated using the two's complement representation, which involves flipping all the bits of bitmask and adding 1 to the result.
	// The result of this bitwise AND operation isolates the rightmost set bit of bitmask.
	// This technique is commonly used to identify the rightmost set bit in a binary number, which is then used in further bitwise operations to manipulate and extract specific bits from the original number.
	diff := bitmask & (-bitmask)
	// -bitmask = -0000110 = 11111001 + 1 = 11111010
	// diff = 00000110 & 11111010 = 00000010
	fmt.Printf(">>>>> -bitmask %08b, %d\n", -bitmask, -bitmask)
	fmt.Printf(">>>>> bitmask & (-bitmask) %08b, %d\n", diff, diff)
	fmt.Println()

	x := 0
	for _, v := range nums {
		fmt.Printf(">>>>> v %08b, %d\n", v, v)
		// Find the number that has the same rightmost set bit with 'diff'
		if (v & diff) != 0 {
			fmt.Printf(">>>>> v & diff %08b\n", v&diff)
			x ^= v
			fmt.Printf(">>>>> x ^= v %08b\n", x)
		}
	}

	fmt.Println()
	// The result is 3
	fmt.Printf(">>>>> final x %08b, %d\n", x, x)
	// 3^5 = 6, 6^5 = 3, 6^3 = 5
	fmt.Printf(">>>>> bitmask ^ x %08b, %d\n", bitmask^x, bitmask^x)
	return []int{x, bitmask ^ x}
}
