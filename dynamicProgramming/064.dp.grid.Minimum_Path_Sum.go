// package dynamic_programming
// dir_path dynamic_programming
// name 064.dp.Minimum_Path_Sum

package dynamicProgramming

import "fmt"

/**
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom
right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.
*/

/**
To find a path from the top left to the bottom right of a grid that minimizes the sum of all numbers along its path,
you can use dynamic programming
*/

func minPathSum(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])

	// Create a 2D array to store the minimum path sum
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
	}

	dp[0][0] = grid[0][0]
	// Initialize the first row and first column of dp
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for j := 1; j < columns; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// Calculate the minimum path sum for each cell in dp
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	fmt.Println(dp)
	return dp[rows-1][columns-1]
}
