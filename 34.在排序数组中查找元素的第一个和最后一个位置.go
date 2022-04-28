/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
package main

func searchRange(nums []int, target int) []int {
	l, r := -1, len(nums)
	find := -1
	for l+1 < r {
		m := l + (r-l)/2
		if nums[m] == target {
			find = m
			break
		} else if nums[m] < target {
			l = m
		} else {
			r = m
		}
	}
	if find == -1 {
		return []int{-1, -1}
	}

	lr, rl := find, find
	for l+1 < lr {
		m := l + (lr-l)/2
		if nums[m] == target {
			lr = m
		} else {
			l = m
		}
	}
	for rl+1 < r {
		m := rl + (r-rl)/2
		if nums[m] == target {
			rl = m
		} else {
			r = m
		}
	}
	return []int{lr, rl}
}

// @lc code=end
