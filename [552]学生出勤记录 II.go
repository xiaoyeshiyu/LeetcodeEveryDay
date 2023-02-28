//可以用字符串表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
//
//
// 'A'：Absent，缺勤
// 'L'：Late，迟到
// 'P'：Present，到场
//
//
// 如果学生能够 同时 满足下面两个条件，则可以获得出勤奖励：
//
//
// 按 总出勤 计，学生缺勤（'A'）严格 少于两天。
// 学生 不会 存在 连续 3 天或 连续 3 天以上的迟到（'L'）记录。
//
//
// 给你一个整数 n ，表示出勤记录的长度（次数）。请你返回记录长度为 n 时，可能获得出勤奖励的记录情况 数量 。答案可能很大，所以返回对 10⁹ + 7
//取余 的结果。
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出：8
//解释：
//有 8 种长度为 2 的记录将被视为可奖励：
//"PP" , "AP", "PA", "LP", "PL", "AL", "LA", "LL"
//PPA LPA PLA LLA

//只有"AA"不会被视为可奖励，因为缺勤次数为 2 次（需要少于 2 次）。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：3
//
//
// 示例 3：
//
//
//输入：n = 10101
//输出：183236316
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁵
//
//
// Related Topics 动态规划 👍 288 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func checkRecord(n int) int {
	// 定义dp，三位数组，
	// 一维代表第n次
	// 二维代表A的次数
	// 三位对应L的连续次数
	// 值是对应的奖励的次数
	dp := make([][2][3]int, n)

	//	初始化，第一天A、L、P的情况下的次数
	dp[0] = [2][3]int{
		// 0次A
		{
			//	0次L、1次L，2次L
			1, 1, 0,
		},
		// 1次A
		{
			// 0次L、1次L，2次L
			1, 0, 0,
		},
	}

	//	状态转移方程
	for i := 1; i < n; i++ {
		// 第i天，0次A，0次L，只能是昨天0次L，1次L，2次L，然后今天P
		dp[i][0][0] = (dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2]) % 1000000007
		// 第i天，0次A，1次L，只能是昨天0次A，0次L
		dp[i][0][1] = dp[i-1][0][0]
		// 第i天，0次A，2次L，只能是昨天0次A，1次L
		dp[i][0][2] = dp[i-1][0][1]
		// 第i天，1次A，0次L，可以是昨天0次A，0次L，1次L，2次L，然后今天A
		// 也可以是昨天1次A，0次L，1次L，2次L，然后今天P
		dp[i][1][0] = (dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2] + dp[i-1][1][0] + dp[i-1][1][1] + dp[i-1][1][2]) % 1000000007
		// 第i天，1次A，1次L，只能是昨天1次A，0次L
		dp[i][1][1] = dp[i-1][1][0]
		// 第i天，1次A，2次L，只能是昨天1次A，1次L
		dp[i][1][2] = dp[i-1][1][1]
	}
	//fmt.Println(dp)
	return (dp[n-1][0][0] + dp[n-1][0][1] + dp[n-1][0][2] + dp[n-1][1][0] + dp[n-1][1][1] + dp[n-1][1][2]) % 1000000007
}

//leetcode submit region end(Prohibit modification and deletion)
