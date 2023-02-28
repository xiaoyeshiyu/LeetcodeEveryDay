//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
// Related Topics 字符串 动态规划 回溯 👍 3009 👎 0

package main

import (
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	last := generateParenthesis(n - 1)
	lastSli := make([][]string, 0, len(last))
	for i := 0; i < len(last); i++ {
		lastSli = append(lastSli, strings.Split(last[i], ""))
	}

	res := make([]string, 0)
	all := make(map[string]bool)

	for m := 0; m < len(lastSli); m++ {
		for i := 0; i < 2*n-1; i++ {
			for j := i + 1; j < 2*n; j++ {
				tmp := make([]string, 2*n)
				for k, l := 0, 0; k < 2*n; k++ {
					if k == i {
						tmp[k] = "("
					} else if k == j {
						tmp[k] = ")"
					} else {
						tmp[k] = lastSli[m][l]
						l++
					}
				}
				tmpstr := strings.Join(tmp, "")
				if all[tmpstr] {
					continue
				} else {
					all[tmpstr] = true
					res = append(res, tmpstr)
				}
			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
