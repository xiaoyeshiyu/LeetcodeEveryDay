//给你一个字符串 s ，找出它的所有子串并按字典序排列，返回排在最后的那个子串。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "abab"
//输出："bab"
//解释：我们可以找出 7 个子串 ["a", "ab", "aba", "abab", "b", "ba", "bab"]。按字典序排在最后的子串是 
//"bab"。
// 
//
// 示例 2： 
//
// 
//输入：s = "leetcode"
//输出："tcode"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 4 * 10⁵ 
// s 仅含有小写英文字符。 
// 
//
// Related Topics 双指针 字符串 👍 100 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lastSubstring(s string) string {
	i, j, n := 0, 1, len(s)
	for j < n {
		k := 0
		for j + k < n && s[i + k] == s[j + k] {
			k++
		}
		if j + k < n && s[i + k] < s[j + k] {
			i, j = j, max(j + 1, i + k + 1)
		} else {
			j = j + k + 1
		}
	}
	return s[i:]
}

//leetcode submit region end(Prohibit modification and deletion)
