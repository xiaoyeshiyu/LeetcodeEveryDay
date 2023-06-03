// 如果字符串中的所有字符都相同，那么这个字符串是单字符重复的字符串。
//
// 给你一个字符串 text，你只能交换其中两个字符一次或者什么都不做，然后得到一些单字符重复的子串。返回其中最长的子串的长度。
//
// 示例 1：
//
// 输入：text = "ababa"
// 输出：3
//
// 示例 2：
//
// 输入：text = "aaabaaa"
// 输出：6
//
// 示例 3：
//
// 输入：text = "aaabbaaa"
// 输出：4
//
// 示例 4：
//
// 输入：text = "aaaaa"
// 输出：5
//
// 示例 5：
//
// 输入：text = "abcdef"
// 输出：1
//
// 提示：
//
// 1 <= text.length <= 20000
// text 仅由小写英文字母组成。
//
// Related Topics 哈希表 字符串 滑动窗口 👍 176 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxRepOpt1(text string) int {
	count := make(map[rune]int)
	for _, c := range text {
		count[c]++
	}

	res := 0
	i := 0
	for i < len(text) {
		// step1: 找出当前连续的一段 [i, j)
		j := i
		for j < len(text) && text[j] == text[i] {
			j++
		}
		curCnt := j - i

		// step2: 如果这一段长度小于该字符出现的总数，并且前面或后面有空位，则使用 cur_cnt + 1 更新答案
		if curCnt < count[rune(text[i])] && (j < len(text) || i > 0) {
			res = max(res, curCnt+1)
		}

		// step3: 找到这一段后面与之相隔一个不同字符的另一段 [j + 1, k)，如果不存在则 k = j + 1
		k := j + 1
		for k < len(text) && text[k] == text[i] {
			k++
		}
		res = max(res, min(k-i, count[rune(text[i])]))
		i = j
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
