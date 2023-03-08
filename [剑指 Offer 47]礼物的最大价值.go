// 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直
// 到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
//
// 示例 1:
//
// 输入:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// 输出: 12
// 解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
//
// 提示：
//
// 0 < grid.length <= 200
// 0 < grid[0].length <= 200
//
// Related Topics 数组 动态规划 矩阵 👍 390 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// 定义dp，代表对应格子的最大值
	dp := make([][]int, m)
	dp[0] = make([]int, n)
	// 初始化，第0，0格的价格对应
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] += dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] += dp[i-1][0] + grid[i][0]
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

//leetcode submit region end(Prohibit modification and deletion)
