//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//
//
//
// 示例 1：
//
//
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
//,["1","0","0","1","0"]]
//输出：4
//
//
// 示例 2：
//
//
//输入：matrix = [["0","1"],["1","0"]]
//输出：1
//
//
// 示例 3：
//
//
//输入：matrix = [["0"]]
//输出：0
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 300
// matrix[i][j] 为 '0' 或 '1'
//
//
// Related Topics 数组 动态规划 矩阵 👍 1355 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]byte) int {
	length := len(matrix)
	height := len(matrix[0])

	// 定义dp，二维数组，代表这个点的最大正方形边长
	dp := make([][]int, length)
	// 最长边长
	var res int
	for i := 0; i < length; i++ {
		tmp := make([]int, height)
		dp[i] = tmp
		for j := 0; j < height; j++ {
			if (i == 0 || j == 0) && matrix[i][j] == '1' {
				dp[i][j] = 1
				res = 1
			}
		}
	}

	for i := 1; i < length; i++ {
		for j := 1; j < height; j++ {
			//	状态转移方程
			// 只有当本身是1，则可以是一个正方形
			if matrix[i][j] == '1' {
				// 最大正方形的周长，是（左边，上边，左上边）正方形最大周长中的最小值+1
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}

	return res * res
}

func min(i int, i2 int, i3 int) int {
	if i > i2 {
		i = i2
	}
	if i > i3 {
		return i3
	}
	return i
}

//leetcode submit region end(Prohibit modification and deletion)
