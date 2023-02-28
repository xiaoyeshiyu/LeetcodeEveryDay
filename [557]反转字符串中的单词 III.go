//给定一个字符串
// s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
//
//
// 示例 1：
//
//
//输入：s = "Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
//
//
// 示例 2:
//
//
//输入： s = "God Ding"
//输出："doG gniD"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 5 * 10⁴
//
// s 包含可打印的 ASCII 字符。
//
// s 不包含任何开头或结尾空格。
//
// s 里 至少 有一个词。
//
// s 中的所有单词都用一个空格隔开。
//
//
// Related Topics 双指针 字符串 👍 520 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	ss := strings.Split(s," ")
	for i := 0; i < len(ss); i++ {
		ss[i] = reverseWord(ss[i])
	}
	return strings.Join(ss," ")
}

func reverseWord(s string) string {
	ss := []byte(s)
	for i := 0; i < len(ss) >> 1 ; i++ {
		ss[i],ss[len(ss)-1-i] = ss[len(ss)-1-i],ss[i]
	}
	return string(ss)
}

//leetcode submit region end(Prohibit modification and deletion)
