// 给定正整数 n，返回在 [1, n] 范围内具有 至少 1 位 重复数字的正整数的个数。
//
// 示例 1：
//
// 输入：n = 20
// 输出：1
// 解释：具有至少 1 位重复数字的正数（<= 20）只有 11 。
//
// 示例 2：
//
// 输入：n = 100
// 输出：10
// 解释：具有至少 1 位重复数字的正数（<= 100）有 11，22，33，44，55，66，77，88，99 和 100 。
//
// 示例 3：
//
// 输入：n = 1000
// 输出：262
//
// 提示：
//
// 1 <= n <= 10⁹
//
// Related Topics 数学 动态规划 👍 133 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func numDupDigitsAtMostN(n int) (ans int) {
	s := strconv.Itoa(n)
	m := len(s)
	all := make([][1 << 10]int, m)
	var f func(i int) int
	f = func(i int) int {
		if i == m {
			return 1 
		}
	}

	return f(1)
}

//leetcode submit region end(Prohibit modification and deletion)
