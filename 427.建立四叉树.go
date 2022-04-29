/*
 * @lc app=leetcode.cn id=427 lang=golang
 *
 * [427] 建立四叉树
 */

// @lc code=start
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */
package main

func construct(grid [][]int) *Node {
	n := len(grid)
	val := grid[0][0]
	if n == 1 {
		return &Node{Val: toBool(val), IsLeaf: true}
	}

	sa := true
loop:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != val {
				sa = false
				break loop
			}
		}
	}
	if sa {
		return &Node{Val: toBool(val), IsLeaf: true}
	}

	node := &Node{Val: false, IsLeaf: false}

	topLeft := make([][]int, n/2)
	for i := 0; i < n/2; i++ {
		topLeft[i] = grid[i][:n/2]
	}
	node.TopLeft = construct(topLeft)

	topRight := make([][]int, n/2)
	for i := 0; i < n/2; i++ {
		topRight[i] = grid[i][n/2:]
	}
	node.TopRight = construct(topRight)

	bottomLeft := make([][]int, n/2)
	for i := n / 2; i < n; i++ {
		bottomLeft[i] = grid[i][:n/2]
	}
	node.BottomLeft = construct(bottomLeft)

	bottomRight := make([][]int, n/2)
	for i := n / 2; i < n; i++ {
		bottomRight[i] = grid[i][n/2:]
	}
	node.BottomRight = construct(bottomRight)

	return node
}

func toBool(x int) bool {
	if x == 0 {
		return false
	} else {
		return true
	}
}

// @lc code=end

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func main() {
	grid := [][]int{{0, 1}, {1, 0}}
	construct(grid)
}
