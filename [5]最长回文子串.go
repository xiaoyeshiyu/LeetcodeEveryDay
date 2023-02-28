//给你一个字符串 s，找到 s 中最长的回文子串。
//
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 6155 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	// 暴力法
	max := 1
	res := s[0:1]
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if isPalindrome(s[i:j+1]) && j-i+1 > max {
				max = j - i
				res = s[i : j+1]
			}
		}
	}

	return res
}

func isPalindrome(s string) bool {
	length := len(s)
	start, end := 0, length-1
	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
