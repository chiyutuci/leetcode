/*
 * @lc app=leetcode.cn id=1351 lang=golang
 *
 * [1351] 统计有序矩阵中的负数
 */

// @lc code=start
package main

func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	row, col := 0, n-1
	ans := 0
	for ; row < m; row++ {
		for ; col >= 0; col-- {
			if grid[row][col] >= 0 {
				break
			}
		}
		ans += n - 1 - col
	}
	return ans
}

// @lc code=end
