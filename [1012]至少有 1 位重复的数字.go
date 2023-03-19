//给定正整数 n，返回在 [1, n] 范围内具有 至少 1 位 重复数字的正整数的个数。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 20
//输出：1
//解释：具有至少 1 位重复数字的正数（<= 20）只有 11 。
// 
//
// 示例 2： 
//
// 
//输入：n = 100
//输出：10
//解释：具有至少 1 位重复数字的正数（<= 100）有 11，22，33，44，55，66，77，88，99 和 100 。
// 
//
// 示例 3： 
//
// 
//输入：n = 1000
//输出：262
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 10⁹ 
// 
//
// Related Topics 数学 动态规划 👍 133 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func numDupDigitsAtMostN(n int) (ans int) {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][1 << 10]int, m)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1 // -1 表示没有计算过
		}
	}
	var f func(int, int, bool, bool) int
	f = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1 // 得到了一个合法数字
			}
			return
		}
		if !isLimit && isNum {
			dv := &dp[i][mask]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		if !isNum { // 可以跳过当前数位
			res += f(i+1, mask, false, false)
		}
		d := 0
		if !isNum {
			d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0') // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
		}
		for ; d <= up; d++ { // 枚举要填入的数字 d
			if mask>>d&1 == 0 { // d 不在 mask 中
				res += f(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return n - f(0, 0, true, false)
}

//leetcode submit region end(Prohibit modification and deletion)
