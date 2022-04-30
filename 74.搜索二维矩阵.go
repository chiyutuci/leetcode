/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
package main

func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	l, r := -1, rows*cols
	for l+1 < r {
		m := l + (r-l)/2
		row := m / cols
		col := m % cols
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			l = m
		} else {
			r = m
		}
	}
	return false
}

// @lc code=end
