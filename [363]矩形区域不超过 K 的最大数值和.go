//给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。
//
// 题目数据保证总会存在一个数值和不超过 k 的矩形区域。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,0,1],[0,-2,3]], k = 2
//输出：2
//解释：蓝色边框圈出来的矩形区域 [[0, 1], [-2, 3]] 的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。
//
//
// 示例 2：
//
//
//输入：matrix = [[2,2,-1]], k = 3
//输出：3
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -100 <= matrix[i][j] <= 100
// -10⁵ <= k <= 10⁵
//
//
//
//
// 进阶：如果行数远大于列数，该如何设计解决方案？
//
// Related Topics 数组 二分查找 矩阵 有序集合 前缀和 👍 436 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func maxSumSubmatrix(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])

	var res int

	// dp为1，1到m，n之间的数字和
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		dp[i] = tmp
		for j := 0; j < n; j++ {
			// 先计算每个dp的值
			//	dp[i][j] = dp[i-1][j]+dp[i][j-1]-dp[i-1][j-1]+matrix[i][j]
			if matrix[i][j] < 0 {
				res += matrix[i][j]
			}
			if i == 0 && j == 0 {
				dp[i][j] = matrix[i][j]
			} else if i > 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + matrix[i][j]
			} else if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + matrix[i][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i][j]
			}
			if dp[i][j] == k {
				return k
			}
		}
	}

	// 暴力便利
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			//	先确定一个点
			//	再确定另外一个点
			for x := 0; x <= i; x++ {
				for y := 0; y <= j; y++ {
					//	计算这两个点之间的大小
					var tmp int
					if x == 0 && y == 0 {
						tmp = dp[i][j]
					} else if x == 0 && y > 0 {
						tmp = dp[i][j] - dp[i][y-1]
					} else if x > 0 && y == 0 {
						tmp = dp[i][j] - dp[x-1][j]
					} else {
						tmp = dp[i][j] - dp[i][y-1] - dp[x-1][j] + dp[x-1][y-1]
					}
					if tmp == k {
						return k
					}
					if tmp < k && tmp > res {
						res = tmp
					}
				}
			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
