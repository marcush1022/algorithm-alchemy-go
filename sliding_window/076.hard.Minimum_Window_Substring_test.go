// package sliding_window
// dir_path sliding_window
// name 076.hard.Minimum_Window_Substring_test

package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minWindow(t *testing.T) {
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
	assert.Equal(t, "a", minWindow("a", "a"))
	assert.Equal(t, "", minWindow("a", "aa"))

}
