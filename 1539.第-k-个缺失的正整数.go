/*
 * @lc app=leetcode.cn id=1539 lang=golang
 *
 * [1539] 第 k 个缺失的正整数
 */

// @lc code=start
package main

func findKthPositive(arr []int, k int) int {
	l, r := -1, len(arr)
	for l+1 < r {
		m := l + (r-l)/2
		num := arr[m] - m - 1
		if num >= k {
			r = m
		} else {
			l = m
		}
	}
	return r + k
}

// @lc code=end
