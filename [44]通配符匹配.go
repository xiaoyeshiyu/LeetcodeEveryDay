// 给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
//
// '?' 可以匹配任何单个字符。
// '*' 可以匹配任意字符串（包括空字符串）。
//
// 两个字符串完全匹配才算匹配成功。
//
// 说明:
//
// s 可能为空，且只包含从 a-z 的小写字母。
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
//
// 示例 1:
//
// 输入:
// s = "aa"
// p = "a"
// 输出: false
// 解释: "a" 无法匹配 "aa" 整个字符串。
//
//	"" a a
//
// "" 1
// a  0 1
// 示例 2:
//
// 输入:
// s = "aa"
// p = "*"
// 输出: true
// 解释: '*' 可以匹配任意字符串。
//
//	"" *
//
// "" 1  1
// a  0  1
// a  0  1
// 示例 3:
//
// 输入:
// s = "cb"
// p = "?a"
// 输出: false
// 解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
//
//	   "" c  b
// ""  1  0  0
// ?   0  1  0
// a   0  0  0
// 示例 4:
//
// 输入:
// s = "adceb"
// p = "*a*b"
// 输出: true
// 解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
// 都是字母，相等，则i-1,j-1
// 如果是*，0则是i-1，否则是j-1,或者i-1
// 如果是？,则是则i-1,j-1
//
//	    "" a d c e b
//	""  1  0 0 0 0 0
//	*   1  1 1 1 1 1
//	a   0  1 0 0 0 0
//	*   0  1 1 1 1 1
//	b   0  0 0 0 0 1
//
// 示例 5:
//
// 输入:
// s = "acdcb"
// p = "a*c?b"
// 输出: false
//
// Related Topics 贪心 递归 字符串 动态规划 👍 1014 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	m, n := len(p), len(s)

	dp := make([][]bool, m+1)
	dp[0] = make([]bool, n+1)
	dp[0][0] = true

	for i := 1; i <= m; i++ {
		dp[i] = make([]bool,n+1)
		if p[i-1] == '*' {
			dp[i][0] = dp[i-1][0]
		}
		for j := 1; j <= n; j++ {
			if unicode.IsLetter(rune(p[i-1])) {
				if p[i-1] == s[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = false
				}
			} else if p[i-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	//fmt.Println(dp)
	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)
