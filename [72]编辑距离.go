//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
//
//
//
// 示例 1：
//
//
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//
//
// 示例 2：
//
//
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
//
//
//
//
// 提示：
//
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 2735 👎 0
package main

// word1 = "horse", word2 = "ros"
//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	// dp定义，二维数组，第一维是s1的长度+1，第二维度是s2的长度+1，代表的是从s1到s2的变化的次数
	dp := make([][]int, l1+1)
	// 例如s1是""，s2是""，则是0
	// 例如s1是""，s2是"r"，则是1
	// 例如s1是""，s2是"ro"，则是2
	// 例如s1是""，s2是"ros"，则是3
	// 同理，例如s1是"h"，s2是""，则是1，以此类推，就是初始化
	for i := 0; i <= l1; i++ {
		tmp := make([]int, l2+1)
		for j := 0; j <= l2; j++ {
			if i == 0 {
				tmp[j] = j
			}
			if j == 0 {
				tmp[j] = i
			}
		}
		dp[i] = tmp
	}

	// 状态转移方程，从s1到s2有三种情况，要么是加1个字符，要么是删除1个字符，要么是修改1个字符，从短到长，可以理解为s1加1个字符变成s2，s2加1个字符变成s1，s1和s2都加1个字符之后相等
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			// 如果两个字符相同，例如"ho"和"ro"，则是从"h"到"ro"，"ho"到"r"，"h"到"r"-1之间，取最小1个+1，也就是min（2，,2，1-1）+1
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]-1) + 1
			//	如果两个字符串不相同，例如"hor"和"ros"，则是从"ho"到"ros"，"hor"到"ro"，"ho"，"ro"之间，取最小1个+1，也就是min（2，,2，1）+1
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[l1][l2]
}

func min(i int, i2 int, i3 int) int {
	if i > i2 {
		i = i2
	}
	if i > i3 {
		i = i3
	}

	return i
}

//leetcode submit region end(Prohibit modification and deletion)
