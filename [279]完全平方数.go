//给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//
//
//
//
// 示例 1：
//
//
//输入：n = 12
//输出：3
//解释：12 = 4 + 4 + 4
//
// 示例 2：
//
//
//输入：n = 13
//输出：2
//解释：13 = 4 + 9
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁴
//
//
// Related Topics 广度优先搜索 数学 动态规划 👍 1577 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func numSquares(n int) int {
	dp := make([]int, n+1)
	//	res := 1 +
	for i := 0; i < n+1; i++ {
		tmp := int(math.Sqrt(float64(i)))
		if tmp*tmp == i {
			dp[i] = 1
			continue
		}
		min := i
		for j := 1; j <= tmp; j++ {
			if min > dp[i-j*j] {
				min = dp[i-j*j]
			}
		}
		dp[i] = min + 1
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
