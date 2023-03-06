// 给你一个字符串 s ，它仅包含字符 'a' 和 'b' 。
//
// 你可以删除 s 中任意数目的字符，使得 s 平衡 。当不存在下标对 (i,j) 满足 i < j ，且 s[i] = 'b' 的同时 s[j]= 'a'
// ，此时认为 s 是 平衡 的。
//
// 请你返回使 s 平衡 的 最少 删除次数。
//
// 示例 1：
//
// 输入：s = "aababbab"
// 输出：2
// 解释：你可以选择以下任意一种方案：
// 下标从 0 开始，删除第 2 和第 6 个字符（"aababbab" -> "aaabbb"），
// 下标从 0 开始，删除第 3 和第 6 个字符（"aababbab" -> "aabbbb"）。
//
// 示例 2：
//
// 输入：s = "bbaaaaabb"
// 输出：2
// 解释：唯一的最优解是删除最前面两个字符。
//
// 提示：
//
// 1 <= s.length <= 10⁵
// s[i] 要么是 'a' 要么是 'b' 。
//
// Related Topics 栈 字符串 动态规划 👍 66 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func minimumDeletions(s string) int {
	length := len(s)

	// dp，记录对应长度字符串的最小变化次数
	dp := make([]int, length)
	// 并且记录b的次数
	var b int
	// 初始化条件，只有一个字符的时候，肯定是0次
	// 如果初始字符是b，则b为1
	if s[0] == 'b' {
		b++
	}
	// 变化条件，如果是a，则要么将这个a删掉，要么把前面所有的b删掉，取最小值
	// 如果是b，则一定是合法的，加上b
	for i := 1; i < length; i++ {
		if s[i] == 'a' {
			dp[i] = min(dp[i-1] + 1,b)
		} else {
			dp[i] = dp[i-1]
			b++
		}
	}

	return dp[length-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
