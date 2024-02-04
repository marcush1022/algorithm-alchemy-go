// package twoPointers
// dir_path twoPointers
// name 003.Longest_Substring_Without_Repeating_Characters

package sliding_window

/**
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3.
Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

/**
To find the length of the longest substring without repeating characters, you can use the sliding window technique.
*/

func lengthOfLongestSubstring(s string) int {
	// Create a map to store the last index of each character
	lastIndex := make(map[byte]int)

	// Initialize the start of the window and the result
	startIndex := 0
	maxLength := 0

	// Iterate over the string
	for i := 0; i < len(s); i++ {
		// Check if the current character is already in the map
		if index, ok := lastIndex[s[i]]; ok {
			// If it is, update the start of the window
			startIndex = index + 1
		}

		// Update the last index of the current character
		lastIndex[s[i]] = i

		// Update the result if necessary
		if i-(startIndex-1) > maxLength {
			maxLength = i - (startIndex - 1)
		}
	}
	return maxLength
}
