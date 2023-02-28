//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：n = 4, k = 2
//输出：
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
//
// 示例 2：
//
//
//输入：n = 1, k = 1
//输出：[[1]]
//
//
//
// 提示：
//
//
// 1 <= n <= 20
// 1 <= k <= n
//
//
// Related Topics 回溯 👍 1239 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	if k == 1 {
		for i := 1; i < n+1; i++ {
			tmp := make([]int, 1, k)
			tmp[0] = i
			res = append(res, tmp)
		}
		return res
	}
	last := combine(n, k-1)
	for i := 0; i < len(last); i++ {
		for j := last[i][len(last[i])-1] + 1; j <= n; j++ {
			tmp := make([]int, len(last[i]))
			copy(tmp, last[i])
			res = append(res, append(tmp, j))
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
