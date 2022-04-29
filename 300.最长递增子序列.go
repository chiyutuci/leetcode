/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	ans := 0

	for i := 1; i < n; i++ {
		if nums[i] > dp[ans] {
			ans++
			dp[ans] = nums[i]
		} else {
			l, r := -1, ans+1
			for l+1 < r {
				m := l + (r-l)/2
				if dp[m] < nums[i] {
					l = m
				} else {
					r = m
				}
			}
			dp[r] = nums[i]
		}
	}

	return ans + 1
}

// @lc code=end
