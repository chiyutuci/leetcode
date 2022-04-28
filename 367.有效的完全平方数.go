/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
package main

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l, r := 0, num
	for l+1 < r {
		m := l + (r-l)/2
		n := m * m
		if n == num {
			return true
		} else if n < num {
			l = m
		} else {
			r = m
		}
	}

	return false
}

// @lc code=end
