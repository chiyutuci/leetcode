/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] 按奇偶排序数组
 */

// @lc code=startp
package main

func sortArrayByParity(nums []int) []int {
	n := len(nums)
	l, r := 0, 0
	for r < n {
		if nums[r]%2 == 0 {
			if nums[l]%2 == 1 {
				temp := nums[r]
				nums[r] = nums[l]
				nums[l] = temp
			}
			l++
		}
		r++
	}
	return nums
}

// @lc code=end
