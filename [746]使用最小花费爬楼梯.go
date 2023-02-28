//给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
//
// 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
//
// 请你计算并返回达到楼梯顶部的最低花费。
//
//
//
// 示例 1：
//
//
//输入：cost = [10,15,20]
//输出：15
//解释：你将从下标为 1 的台阶开始。
//- 支付 15 ，向上爬两个台阶，到达楼梯顶部。
//总花费为 15 。
//
//
// 示例 2：
//
//
//输入：cost = [1,100,1,1,1,100,1,1,100,1]
//输出：6
//解释：你将从下标为 0 的台阶开始。
//- 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
//- 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
//- 支付 1 ，向上爬一个台阶，到达楼梯顶部。
//总花费为 6 。
//
//
//
//
// 提示：
//
//
// 2 <= cost.length <= 1000
// 0 <= cost[i] <= 999
//
//
// Related Topics 数组 动态规划 👍 1116 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	// 定义DP，为爬到这一层的最少花费
	dp := make([]int, length+1)
	// 遍历，从第二层开始遍历
	for i := 2; i < length+1; i++ {
		//	状态转移方程，第i层，可以由i-1层爬过来，也可以从i-2层爬过来，两者取最小值
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[length]
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}

//leetcode submit region end(Prohibit modification and deletion)
