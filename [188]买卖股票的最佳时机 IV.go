//给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1：
//
//
//输入：k = 2, prices = [2,4,1]
//输出：2
//解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
//
// 示例 2：
//
//
//输入：k = 2, prices = [3,2,6,5,0,3]
//输出：7
//解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
//     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3
//。
//
//
//
// 提示：
//
//
// 0 <= k <= 100
// 0 <= prices.length <= 1000
// 0 <= prices[i] <= 1000
//
//
// Related Topics 数组 动态规划 👍 852 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(k int, prices []int) int {
	length := len(prices)
	if k == 0 || length <= 1 {
		return 0
	}

	// 初始化dp，二维数组，第二个维度代表当天第i次购入的最大收益，当天第i次卖出的最大收益，第二个维度长度是2k
	dp := make([][]int, length)
	dp[0] = make([]int, 2*k)
	// 第1天，可以假设今天一次性购入k次，一次性卖出k次，这个时候的收益
	for i := 0; i < 2*k; i++ {
		if i%2 == 0 {
			// 0         1		  2   3 。。。 k-1			操作次数
			// 购入、卖出、购入、卖出...	    	操作
			// 0    1    2    3	  4 5 6 7		操作名称
			dp[0][i] = -prices[0]
		}
	}

	//	状态转移
	for i := 1; i < length; i++ {
		dp[i] = make([]int, 2*k)
		// 任何一天的第一次购入的最大收益，是比较前一天第一次购入的最大收益和今天的是第一次购入的最大收益的最大值
		dp[i][0] = max(dp[i-1][0], -prices[i])
		// 任何一天的第一次卖出的最大收益，是比较前一天第一次购入的最大收益+今天的价格，和前一天第一次卖出的最大收益
		dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][1])
		for j := 1; j < k; j++ {
			//	第i天的第j次购入的最大收益，是昨天第j-1次卖出的收益，要么昨天买了，要么昨天没买，那么今天价格，或者昨天第j次购入，和今天第j次购入，两者的最大值
			dp[i][2*j] = max(dp[i-1][2*(j-1)+1]-prices[i], dp[i-1][2*j])
			//	第i天的第j次卖出的最大收益，是昨天第j次购入+今天的价格，或者昨天第j次卖出，两者的最大值
			dp[i][2*j+1] = max(dp[i-1][2*j]+prices[i], dp[i-1][2*j+1])
		}
	}
	return dp[length-1][2*(k-1)+1]
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

//leetcode submit region end(Prohibit modification and deletion)
