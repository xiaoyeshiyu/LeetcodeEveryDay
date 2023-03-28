// 给出两个字符串 str1 和 str2，返回同时以 str1 和 str2 作为子序列的最短字符串。如果答案不止一个，则可以返回满足条件的任意一个答案。
//
// （如果从字符串 T 中删除一些字符（也可能不删除，并且选出的这些字符可以位于 T 中的 任意位置），可以得到字符串 S，那么 S 就是 T 的子序列）
//
// 示例：
//
// 输入：str1 = "abac", str2 = "cab"
// 输出："cabac"
// 解释：
// str1 = "abac" 是 "cabac" 的一个子串，因为我们可以删去 "cabac" 的第一个 "c"得到 "abac"。
// str2 = "cab" 是 "cabac" 的一个子串，因为我们可以删去 "cabac" 末尾的 "ac" 得到 "cab"。
// 最终我们给出的答案是满足上述属性的最短字符串。
//
// 提示：
//
// 1 <= str1.length, str2.length <= 1000
// str1 和 str2 都由小写英文字母组成。
//
// Related Topics 字符串 动态规划 👍 139 👎 0
package main

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	//fmt.Println(dp)

	var ans strings.Builder
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if str1[i] == str2[j] {
			ans.WriteByte(str1[i])
			i--
			j--
		} else {
			if dp[i][j+1] > dp[i+1][j] {
				ans.WriteByte(str1[i])
				i--
			} else {
				ans.WriteByte(str2[j])
				j--
			}
		}
	}



	for ;i >= 0;i-- {
		ans.WriteByte(str1[i])
	}
	for ;j >= 0;j-- {
		ans.WriteByte(str2[j])
	}
	str := ans.String()
	return reverseString(str)
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

func reverseString(s string) string {
	var ans strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		ans.WriteByte(s[i])
	}
	return ans.String()
}

//leetcode submit region end(Prohibit modification and deletion)
