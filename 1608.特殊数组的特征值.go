/*
 * @lc app=leetcode.cn id=1608 lang=golang
 *
 * [1608] 特殊数组的特征值
 */

// @lc code=start
package main

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	if n <= nums[0] {
		return n
	}
	for i := 1; i < n; i++ {
		num := n - i
		if num <= nums[i] && num > nums[i-1] {
			return num
		}
	}

	return -1
}

// @lc code=end
