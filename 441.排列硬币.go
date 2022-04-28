/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
package main

func arrangeCoins(n int) int {
	l, r := 0, n+1
	for l+1 < r {
		m := l + (r-l)/2
		if m*(m+1) <= 2*n {
			l = m
		} else {
			r = m
		}
	}
	return l
}

// @lc code=end
