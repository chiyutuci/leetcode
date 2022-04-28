/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 */

// @lc code=start
package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	var dfs func(i, j int, ocean [][]bool)
	dfs = func(i, j int, ocean [][]bool) {
		if ocean[i][j] {
			return
		}
		ocean[i][j] = true
		for _, dir := range dirs {
			nx, ny := i+dir.x, j+dir.y
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				if heights[nx][ny] >= heights[i][j] {
					dfs(nx, ny, ocean)
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		dfs(0, i, pacific)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
	}
	for i := 0; i < n; i++ {
		dfs(m-1, i, atlantic)
	}
	for i := 0; i < m; i++ {
		dfs(i, n-1, atlantic)
	}

	var result [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

// @lc code=end
