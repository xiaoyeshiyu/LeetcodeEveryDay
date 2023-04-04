//有 N 堆石头排成一排，第 i 堆中有 stones[i] 块石头。
//
// 每次移动（move）需要将连续的 K 堆石头合并为一堆，而这个移动的成本为这 K 堆石头的总数。
//
// 找出把所有石头合并成一堆的最低成本。如果不可能，返回 -1 。
//
//
//
// 示例 1：
//
// 输入：stones = [3,2,4,1], K = 2
//输出：20
//解释：
//从 [3, 2, 4, 1] 开始。
//合并 [3, 2]，成本为 5，剩下 [5, 4, 1]。
//合并 [4, 1]，成本为 5，剩下 [5, 5]。
//合并 [5, 5]，成本为 10，剩下 [10]。
//总成本 20，这是可能的最小值。
//
//
// 示例 2：
//
// 输入：stones = [3,2,4,1], K = 3
//输出：-1
//解释：任何合并操作后，都会剩下 2 堆，我们无法再进行合并。所以这项任务是不可能完成的。.
//
//
// 示例 3：
//
// 输入：stones = [3,5,1,2,6], K = 3
//输出：25
//解释：
//从 [3, 5, 1, 2, 6] 开始。
//合并 [5, 1, 2]，成本为 8，剩下 [3, 8, 6]。
//合并 [3, 8, 6]，成本为 17，剩下 [17]。
//总成本 25，这是可能的最小值。
//
//
//
//
// 提示：
//
//
// 1 <= stones.length <= 30
// 2 <= K <= 30
// 1 <= stones[i] <= 100
//
//
// Related Topics 数组 动态规划 👍 239 👎 0

package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func mergeStones(stones []int, k int) int {
	m := len(stones)

	if (m-1)%(k-1) != 0 {
		return -1
	}

	sum := make([][]int, m)
	for i := 0; i < m; i++ {
		sum[i] = make([]int, m)
		sum[i][i] = stones[i]
		for j := i + 1; j < m; j++ {
			sum[i][j] = sum[i][j-1] + stones[j]
		}
	}

	// 定义三维数组，一维是左边边界，二维是右边边界，三维是堆数，值是最低成本
	//
	// 则有状态转移
	// 1. 将k堆，集合成1堆
	// dp[i][j][1] = dp[i][j][k] + sum[i][j]
	// 2. 集合成p堆
	// dp[i][j][p] = min(dp[i][l][1] + dp[l+1][j][p-1]);其中l=[i,j-1],p=[2,k]
	// 初始化 dp[i][i][1] = 0 代表不需要移动

	dp := make([][][]int, m)
	for i := m - 1; i >= 0; i-- {
		dp[i] = make([][]int, m)
		dp[i][i] = make([]int, k+1)
		dp[i][i][1] = 0
		for j := i + 1; j < m; j++ {
			dp[i][j] = make([]int, k+1)
			for p := 2; p < k+1; p++ {
				dp[i][j][p] = math.MaxInt
				for l := i; l < j; l += k - 1 {
					dp[i][j][p] = min(dp[i][j][p], dp[i][l][1]+dp[l+1][j][p-1])
				}
			}
			dp[i][j][1] = dp[i][j][k] + sum[i][j]
		}
	}

	return dp[0][m-1][1]
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}

//leetcode submit region end(Prohibit modification and deletion)
