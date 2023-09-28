// package golang
// name 005.(中心扩展法)Longest_Palindromic_Substring

package uncategorized

import "fmt"

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum
length of s is 1000.

Example:

Input: "babad"

Output: "bab"

Note: "aba" is also a valid answer.
Example:

Input: "cbbd"

Output: "bb"
*/

/**
To find the longest palindromic substring in a given string, you can use the "expand around center" approach
*/

func getLongest(str string, left, right int) string {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}
	return str[left+1 : right]
}

func LongestPalindrome(str string) string {
	longest := str[0:1]
	for i := 0; i < len(str); i++ {
		// aba
		tmp := getLongest(str, i, i)
		if len(tmp) > len(longest) {
			longest = tmp
		}

		// aa
		tmp = getLongest(str, i, i+1)
		if len(tmp) > len(longest) {
			longest = tmp
		}
	}
	fmt.Println(">>>>>> longest", longest)
	return longest
}

// AI generated solution:
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		length := max(len1, len2)

		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	// CAUTION: This selects a half-open range which includes the first element, but excludes the last one
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	/**
	When the loop terminates, the left pointer will be one position to the left of the start of the palindrome, and the
	right pointer will be one position to the right of the end of the palindrome. To calculate the length of the palindrome,
	we subtract left from right and subtract 1 to account for the extra positions moved by the pointers
	*/
	return right - left - 1
}
