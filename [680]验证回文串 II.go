//给你一个字符串 s，最多 可以从中删除一个字符。
//
// 请你判断 s 是否能成为回文字符串：如果能，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：s = "aba"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "abca"
//输出：true
//解释：你可以删除字符 'c' 。
//
//
// 示例 3：
//
//
//输入：s = "abc"
//输出：false
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁵
// s 由小写英文字母组成
//
//
// Related Topics 贪心 双指针 字符串 👍 566 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start <= end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			ss := []byte(s)
			newS1 := ss[start+1 : end+1]
			newS2 := ss[start:end]
			return isPalindoroms(newS1) || isPalindoroms(newS2)
		}
	}
	return true
}

func isPalindoroms(s []byte) bool {
	start, end := 0, len(s)-1
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
