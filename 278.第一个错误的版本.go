/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
package main

func firstBadVersion(n int) int {
	l, r := 0, n+1
	for l+1 < r {
		m := l + (r-l)/2
		if isBadVersion(m) {
			r = m
		} else {
			l = m
		}
	}
	return r
}

// @lc code=end
var isBadVersion func(int) bool
