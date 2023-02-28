//给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。
//
// 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
//
// 返回获得利润的最大值。
//
// 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
//
//
//
// 示例 1：
//
//
//输入：prices = [1, 3, 2, 8, 4, 9], fee = 2
//输出：8
//解释：能够达到的最大利润:
//在此处买入 prices[0] = 1
//在此处卖出 prices[3] = 8
//在此处买入 prices[4] = 4
//在此处卖出 prices[5] = 9
//总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8
//
// 示例 2：
//
//
//输入：prices = [1,3,7,5,10,3], fee = 3
//输出：6
//
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 5 * 10⁴
// 1 <= prices[i] < 5 * 10⁴
// 0 <= fee < 5 * 10⁴
//
//
// Related Topics 贪心 数组 动态规划 👍 845 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int, fee int) int {
	length := len(prices)
	if length == 1 {
		return 0
	}

	// dp状态，二维，代表当前是否持有股票的最大收益，这里取卖出的时候缴纳手续费
	// 0 代表持有股票的钱
	// 1 代表不持有股票的钱
	dp := make([][2]int, length)
	//	初始化条件，第1天购入的最大收益，卖出的最大收益
	dp[0] = [2]int{-prices[0], 0}
	for i := 1; i < length; i++ {
		//	状态转移方程
		//	第i天，持有的最大收益，要么今天买，要么今天不买，用昨天的股票，两者取最大值，今天买，就是昨天卖出的收益-今天的价格，今天不买，就是昨天买的收益
		dp[i][0] = max(dp[i-1][1]-prices[i], dp[i-1][0])
		//  第i天，不持有的最大收益，要么是将昨天持有的今天卖掉，要么今天不卖，两者取最大值，今天卖，就是昨天持有的最大收入+今天的价格-今天的手续费，今天不卖，就是昨天不持有股票的最大收益
		dp[i][1] = max(dp[i-1][0]+prices[i]-fee, dp[i-1][1])
	}
	return dp[length-1][1]
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

//leetcode submit region end(Prohibit modification and deletion)
