/*
 * @lc app=leetcode.cn id=908 lang=golang
 *
 * [908] 最小差值 I
 */

// @lc code=start
package main

func smallestRangeI(nums []int, k int) int {
	min, max := 10000, 0
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	if max-min > 2*k {
		return max - min - 2*k
	} else {
		return 0
	}
}

// @lc code=end
