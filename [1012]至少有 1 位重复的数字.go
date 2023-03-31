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

import "strconv"

// leetcode submit region begin(Prohibit modification and deletion)
func numDupDigitsAtMostN(n int) (ans int) {
	s := strconv.Itoa(n)
	m := len(s)
	all := make([][10]bool, m)
	var f func(i int, limit bool, zero bool) int
	f = func(i int, limit bool, zero bool) (res int) {
		// 位数结束，代表该数字可以
		if i == m {
			return 1
		}
		// 先处理前缀是0的情况，此时前面的0不做限制
		if zero {
			// 第一位要限制，后面的位都不限制
			res += f(i+1, false, true)
		}

		// 不限制
		max := 10
		if limit {
			max = int(s[i])
		}
	OUT:
		for j := 0; j < max; j++ {
			// 前缀都是0，当前就不能为0，要从1开始
			// 前缀非0，当前可以从0开始
			if zero && j == 0 {
				continue
			}
			// 判断每一位是否有重复，不限制，可以从0到9
			for k := 0; k < len(all); k++ {
				// 有重复，则下一个数字
				if all[k][j] {
					continue OUT
				}
			}
			all[i][j] = true
			// 没有重复，下一个数字
			if j == max {
				limit = true
			} else {
				limit = false
			}
			res += f(i+1, limit, false)
		}
		return
	}

	// 第0位，肯定是限制的，
	return f(0, true, false)
}

//leetcode submit region end(Prohibit modification and deletion)
