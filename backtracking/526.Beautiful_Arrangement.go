// package backtracking
// dir_path backtracking
// name 526.Beautiful_Arrangement

package backtracking

/**
Suppose you have n integers labeled 1 through n. A permutation of those n integers perm (1-indexed) is considered a beautiful arrangement if for every i (1 <= i <= n),
either of the following is true:

1. perm[i] is divisible by i.
2. i is divisible by perm[i].

Given an integer n, return the number of the beautiful arrangements that you can construct.



Example 1:

Input: n = 2
Output: 2
Explanation:
The first beautiful arrangement is [1,2]:
    - perm[1] = 1 is divisible by i = 1
    - perm[2] = 2 is divisible by i = 2
The second beautiful arrangement is [2,1]:
    - perm[1] = 2 is divisible by i = 1
    - i = 2 is divisible by perm[2] = 1

Example 2:

Input: n = 1
Output: 1


Constraints:

1 <= n <= 15
*/

func countArrangement(n int) int {
	visited := make([]bool, n+1)
	count := 0
	searchArrangement(1, n, visited, &count)
	return count
}

func searchArrangement(pos int, length int, visited []bool, count *int) {
	// First check if this value has gone beyond our length, if so we increment the count and return, so then we can backtrack
	if pos > length {
		*count++
		return
	}
	for i := 1; i <= length; i++ {
		// If it is not yet filled
		if !visited[i] && (i%pos == 0 || pos%i == 0) {
			// Set pos as filled
			visited[i] = true
			// Move to the next position in the arrangement (pos+1)
			searchArrangement(pos+1, length, visited, count)
			// Backtrack
			visited[i] = false
		}
	}
}
