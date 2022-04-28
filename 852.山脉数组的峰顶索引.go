/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
package main

func peakIndexInMountainArray(arr []int) int {
	l, r := -1, len(arr)
	for l+1 < r {
		m := l + (r-l)/2
		switch {
		case m == 0:
			l = m
		case m == len(arr):
			r = m
		default:
			switch {
			case arr[m-1] < arr[m] && arr[m] > arr[m+1]:
				return m
			case arr[m-1] < arr[m] && arr[m] <= arr[m+1]:
				l = m
			default:
				r = m
			}
		}
	}

	return -1
}

// @lc code=end
