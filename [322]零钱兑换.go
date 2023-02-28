//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
//
//
//
// 示例 1：
//
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//
// 示例 2：
//
//
//输入：coins = [2], amount = 3
//输出：-1
//
// 示例 3：
//
//
//输入：coins = [1], amount = 0
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2³¹ - 1
// 0 <= amount <= 10⁴
//
//
// Related Topics 广度优先搜索 数组 动态规划 👍 2246 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)

var all map[int]int

func coinChange(coins []int, amount int) int {
	//if amount == 0 {
	//	return 0
	//}
	//all = make(map[int]int)
	//sort.Ints(coins)
	//for i := 0; i < len(coins); i++ {
	//	all[coins[i]] = 1
	//}
	//return change(coins, amount)
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		min := amount + 1
		dp[i] = min
		for j := 0; j < len(coins); j++ {
			// min (f(i-1) f(i-2) f(i-5))+1
			if i-coins[j] >= 0 && dp[i-coins[j]] < min {
				min = dp[i-coins[j]]
			}
		}
		dp[i] = min + 1
	}
	if dp[amount] == amount+2 {
		return -1
	}
	return dp[amount]
}

//func change(coins []int, amount int) int {
//	if amount < coins[0] {
//		return -1
//	}
//	for i := coins[0] + 1; i <= amount; i++ {
//		min := make([]int, 0, len(coins)+1)
//		if v, ok := all[i]; ok {
//			min = append(min, v)
//		}
//		for j := 0; j < len(coins); j++ {
//			tmp := all[i-coins[j]]
//			if tmp != 0 {
//				min = append(min, tmp+1)
//			}
//		}
//		sort.Ints(min)
//		if len(min) > 0 {
//			all[i] = min[0]
//		}
//	}
//	if v, ok := all[amount]; ok {
//		return v
//	}
//	return -1
//}

//leetcode submit region end(Prohibit modification and deletion)
