/*
 * @lc app=leetcode.cn id=1385 lang=golang
 *
 * [1385] 两个数组间的距离值
 */

// @lc code=start
package main

import "sort"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	n := len(arr2)
	sort.Slice(arr2, func(i, j int) bool { return arr2[i] < arr2[j] })
	ans := 0
	for _, num := range arr1 {
		if num > arr2[n-1] {
			if num > arr2[n-1]+d {
				ans++
				continue
			}
		} else if num < arr2[0] {
			if num < arr2[0]-d {
				ans++
				continue
			}
		} else {
			l, r := -1, n
			for l+1 < r {
				m := l + (r-l)/2
				if arr2[m] == num {
					break
				} else if arr2[m] < num {
					l = m
				} else {
					r = m
				}
			}
			if arr2[r-1]+d < num && num+d < arr2[r] {
				ans++
			}
		}
	}
	return ans
}

// @lc code=end
