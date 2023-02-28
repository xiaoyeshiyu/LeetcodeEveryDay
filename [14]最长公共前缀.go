//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
//
// Related Topics 字典树 字符串 👍 2636 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	res := len(strs[0])

	for i := 1; i < len(strs); i++ {
		var tmp int
		for j := 0; j < len(strs[i]) && j < res; j++ {
			if strs[0][j] == strs[i][j] {
				tmp ++
			} else {
				break
			}
		}
		res = tmp
	}
	return strs[0][:res]
}

//leetcode submit region end(Prohibit modification and deletion)
