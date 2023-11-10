// package sliding_window
// dir_path sliding_window
// name 003.medium.Longest_Substring_Without_Repeating_Characters_test

package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 1, lengthOfLongestSubstring("b"))
}
