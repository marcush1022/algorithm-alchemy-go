// package sliding_window
// dir_path sliding_window
// name 076.hard.Minimum_Window_Substring

package slidingWindow

import (
	"math"
)

/**
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character
in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.



Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
*/

/**
To find the minimum window substring of string s that contains all the characters from string t,
you can use the sliding window technique along with two pointers.
*/

func minWindow(s string, t string) string {
	// Create a map to store the frequency of characters in t
	targetFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		targetFreq[t[i]]++
	}

	// Initialize variables for the sliding window
	left := 0
	matchCount := 0
	minLen := math.MaxInt32
	minLeft := 0
	lenTarget := len(t)

	// Expand the window by moving the right pointer
	for right := 0; right < len(s); right++ {
		// If right exists in target
		if _, ok := targetFreq[s[right]]; ok {
			targetFreq[s[right]]--
			if targetFreq[s[right]] >= 0 {
				matchCount++
			}

			// Shrink the window by moving the left pointer
			for matchCount == lenTarget {
				if right-left+1 < minLen {
					minLen = right - left + 1
					minLeft = left
				}

				// If left exists in target
				if _, ok := targetFreq[s[left]]; ok {
					targetFreq[s[left]]++
					if targetFreq[s[left]] > 0 {
						matchCount--
					}
				}
				left++
			}
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}
