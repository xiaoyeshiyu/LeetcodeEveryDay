// 你需要制定一份 d 天的工作计划表。工作之间存在依赖，要想执行第 i 项工作，你必须完成全部 j 项工作（ 0 <= j < i）。
//
// 你每天 至少 需要完成一项任务。工作计划的总难度是这 d 天每一天的难度之和，而一天的工作难度是当天应该完成工作的最大难度。
//
// 给你一个整数数组 jobDifficulty 和一个整数 d，分别代表工作难度和需要计划的天数。第 i 项工作的难度是 jobDifficulty[i]。
//
// 返回整个工作计划的 最小难度 。如果无法制定工作计划，则返回 -1 。
//
// 示例 1：
//
// 输入：jobDifficulty = [6,5,4,3,2,1], d = 2
// 输出：7
// 解释：第一天，您可以完成前 5 项工作，总难度 = 6.
// 第二天，您可以完成最后一项工作，总难度 = 1.
// 计划表的难度 = 6 + 1 = 7
//
// 示例 2：
//
// 输入：jobDifficulty = [9,9,9], d = 4
// 输出：-1
// 解释：就算你每天完成一项工作，仍然有一天是空闲的，你无法制定一份能够满足既定工作时间的计划表。
//
// 示例 3：
//
// 输入：jobDifficulty = [1,1,1], d = 3
// 输出：3
// 解释：工作计划为每天一项工作，总难度为 3 。
//
// 示例 4：
//
// 输入：jobDifficulty = [7,1,7,1,7,1], d = 3
// 输出：15
//
// 示例 5：
//
// 输入：jobDifficulty = [11,111,22,222,33,333,44,444], d = 6
// 输出：843
//
// 提示：
//
// 1 <= jobDifficulty.length <= 300
// 0 <= jobDifficulty[i] <= 1000
// 1 <= d <= 10
//
// Related Topics 数组 动态规划 👍 107 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func minDifficulty(jobDifficulty []int, d int) int {
	l := len(jobDifficulty)
	if d > l {
		return -1
	}
	// 二维数组，记录只有i天的时候，解决j个任务时，所需要的最小难度
	dp := make([][]int, d)
	dp[0] = make([]int, l)
	dp[0][0] = jobDifficulty[0]
	// 第1天的最小难度
	for i := 1; i < l; i++ {
		dp[0][i] = max(dp[0][i-1], jobDifficulty[i])
	}
	// 从第2天开始的最小难度，为1天处理新增任务，直到1天处理所有能处理的任务的最小难度  dp[i][j] = min(dp[i-1][j-k]+)  1<= k <= j-(i-1)
	for i := 1; i < d; i++ {
		dp[i] = make([]int, l)
		for j := i; j < l; j++ {
			tmpMax := jobDifficulty[j]
			dp[i][j] = dp[i-1][j-1] + tmpMax
			for k := 1; k < j-i+1; k++ {
				tmpMax = max(jobDifficulty[j-k], tmpMax)
				dp[i][j] = min(dp[i][j], dp[i-1][j-1-k]+tmpMax)
			}
		}
	}
	return dp[d-1][l-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
