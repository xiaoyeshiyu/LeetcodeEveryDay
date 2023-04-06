//给你一个整数 n ，以二进制字符串的形式返回该整数的 负二进制（base -2）表示。
//
// 注意，除非字符串就是 "0"，否则返回的字符串中不能含有前导零。
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出："110"
//解释：(-2)² + (-2)¹ = 2
//
//
// 示例 2：
//
//
//输入：n = 3
//输出："111"
//解释：(-2)² + (-2)¹ + (-2)⁰ = 3
//
//
// 示例 3：
//
//
//输入：n = 4
//输出："100"
//解释：(-2)² = 4
//
//
//
//
// 提示：
//
//
// 0 <= n <= 10⁹
//
//
// Related Topics 数学 👍 80 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	var ans []string
	k := -1
	for n != 0 {
		if n%2 == 0 {
			ans = append(ans, "0")
		} else {
			ans = append(ans, "1")
			n += k
		}
		k *= -1
		n >>= 1
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}

	return strings.Join(ans,"")
}

//leetcode submit region end(Prohibit modification and deletion)
