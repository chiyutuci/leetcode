/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
package main

func searchInsert(nums []int, target int) int {
	l, r := -1, len(nums)
	for l+1 < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m
		} else {
			r = m
		}
	}
	return r
}

// @lc code=end
