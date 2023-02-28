//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
//
// 示例 2：
//
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
//
// 示例 3：
//
//
//输入：s = ""
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'
//
//
// Related Topics 栈 字符串 动态规划 👍 2131 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	length := len(s)
	if length < 2 {
		return 0
	}

	// dp为到这个字符串结尾时，括号的长度，只有结尾是）才有可能是一个合法的字符串
	dp := make([]int, length)
	var res int
	// 初始化条件，（（、））、）（的情况，最长都是0，只有（）时，dp[1]是2，此时最长的就是2
	if s[0] == '(' && s[1] == ')' {
		dp[1] = 2
		res = 2
	}
	for i := 2; i < length; i++ {
		//	当最后1位是）,才是一个合法的括号
		if s[i] == ')' {
			if s[i-1] == '(' {
				//	倒数第二位是（，例如（（），则是2+再往前数2位的长度
				dp[i] = dp[i-2] + 2
			} else {
				// 倒数第二位是），例如 （（）），则要往前数1位的dp值，在字符串中查询这个dp值前1位的字符是否是（
				tmp := i - dp[i-1] - 1
				// 是（，才是一个合法的括号，此时的长度就是前1位的dp+2
				if tmp >= 0 && s[tmp] == '(' {
					dp[i] = dp[i-1] + 2
					if tmp > 1 {
						dp[i] += dp[tmp-1]
					}
				}
			}
			res = max(dp[i], res)
		}
	}
	return res
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

//leetcode submit region end(Prohibit modification and deletion)
