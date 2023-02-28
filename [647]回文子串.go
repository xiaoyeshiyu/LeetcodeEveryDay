//给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
//
// 回文字符串 是正着读和倒过来读一样的字符串。
//
// 子字符串 是字符串中的由连续字符组成的一个序列。
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
//
//
// 示例 1：
//
//
//输入：s = "abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//
//
// 示例 2：
//
//
//输入：s = "aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 1062 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {
	length := len(s)

	// dp定义，一维数组，代表当以第i个字符结尾时，值为回文个数
	dp := make([]int, length)
	// 初始化，只有1个字符，回文个数为1
	dp[0] = 1

	for i := 1; i < length; i++ {
		// 状态转移
		// dp[i] = dp[i-1] + 以当前字符串结尾的回文
		tmp := 1
		// 从i-1开始往前便利，当字符串是这个的时候，判断是否是回文串
		for start := i - 1; start >= 0; start-- {
			if isSubString(s[start : i+1]) {
				tmp++
			}
		}
		dp[i] = dp[i-1] + tmp
	}
	return dp[length-1]
}

func isSubString(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
