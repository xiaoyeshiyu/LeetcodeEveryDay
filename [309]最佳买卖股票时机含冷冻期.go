//给定一个整数数组
// prices，其中第 prices[i] 表示第 i 天的股票价格 。
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1:
//
//
//输入: prices = [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
//
// 示例 2:
//
//
//输入: prices = [1]
//输出: 0
//
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 5000
// 0 <= prices[i] <= 1000
//
//
// Related Topics 数组 动态规划 👍 1388 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	length := len(prices)
	if length == 1 {
		return 0
	}

	// 定义dp，0表示今天有股票的最大收益，1表示今天没有股票的最大收益，2表示冷冻期的最大收益
	dp := make([][3]int, length)
	// 初始化条件，也就是今天买入股票，或者今天不购入股票
	dp[0] = [3]int{-prices[0], 0, 0}
	for i := 1; i < length; i++ {
		//	状态转移方程
		//	第i天有股票的最大收益，是昨天冷冻期的收益-今天股票价格，或者昨天有股票的收益，取最大值
		dp[i][0] = max(dp[i-1][2]-prices[i], dp[i-1][0])
		//	第i天没有股票的最大收益，是把持有的股票卖出，和昨天冷冻期，今天不购入取最大值
		dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][2])
		//	第i天冷冻期的最大收益，是昨天没有股票，和昨天冷冻期的最大收益
		dp[i][2] = max(dp[i-1][1],dp[i-1][2])
	}
	return max(dp[length-1][1], dp[length-1][2])
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

//leetcode submit region end(Prohibit modification and deletion)
