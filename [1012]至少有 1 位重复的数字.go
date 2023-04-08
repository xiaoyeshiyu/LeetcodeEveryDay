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
	// 记录没有填过的数字的使用次数，用于提升便利速度
	// 例如 如果计算了 12***，那么就不需要计算21***的数量
	all := make([][1 << 10]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < 1<<10; j++ {
			// 将-1记录为没有使用过
			all[i][j] = -1
		}
	}

	// 分别代表第几位，已经填过的数字，是否有限制，前缀是否是0
	var f func(i int, mask int, limit bool, zero bool) int
	f = func(i int, mask int, limit bool, zero bool) (res int) {
		// 位数结束，代表该数字可以
		if i == m {
			return 1
		}

		// 自身没有限制，并且前缀不是0
		if !limit && !zero {
			p := &all[i][mask]
			if *p > 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		// 先处理前面都是0的情况
		if zero {
			// 此时i+1位不限制，并且前面都是0
			res += f(i+1, 0, false, true)
		}

		// 起始值
		start := 0
		if zero {
			start = 1
		}

		// 结束值
		end := 9
		if limit {
			end = int(s[i] - '0')
		}
		//fmt.Println(start,end)
		for ; start <= end; start++ {
			if mask>>start&1 != 1 {
				res += f(i+1, mask|1<<start, limit && start == end, false)
			}
		}
		return
	}

	// 第0位，一个数字都没有使用，肯定是限制的,前面都是0
	return n + 1 - f(0, 0, true, true)
}

//leetcode submit region end(Prohibit modification and deletion)
