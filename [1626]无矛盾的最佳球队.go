// 假设你是球队的经理。对于即将到来的锦标赛，你想组合一支总体得分最高的球队。球队的得分是球队中所有球员的分数 总和 。
//
// 然而，球队中的矛盾会限制球员的发挥，所以必须选出一支 没有矛盾 的球队。如果一名年龄较小球员的分数 严格大于 一名年龄较大的球员，则存在矛盾。同龄球员之间
// 不会发生矛盾。
//
// 给你两个列表 scores 和 ages，其中每组 scores[i] 和 ages[i] 表示第 i 名球员的分数和年龄。请你返回 所有可能的无矛盾球队
// 中得分最高那支的分数 。
//
// 示例 1：
//
// 输入：scores = [1,3,5,10,15], ages = [1,2,3,4,5]
// 输出：34
// 解释：你可以选中所有球员。
//
// 示例 2：
//
// 输入：scores = [4,5,6,5], ages = [2,1,2,1]
// 输出：16
// 解释：最佳的选择是后 3 名球员。注意，你可以选中多个同龄球员。
//
// 示例 3：
//
// 输入：scores = [1,2,3,5], ages = [8,9,10,1]
// 输出：6
// 解释：最佳的选择是前 3 名球员。
//
// 提示：
//
// 1 <= scores.length, ages.length <= 1000
// scores.length == ages.length
// 1 <= scores[i] <= 10⁶
// 1 <= ages[i] <= 1000
//
// Related Topics 数组 动态规划 排序 👍 148 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func bestTeamScore(scores, ages []int) (ans int) {
	type member struct {
		score int
		age   int
	}

	members := make([]member, len(scores))
	for i := 0; i < len(scores); i++ {
		members[i] = member{
			score: scores[i],
			age:   ages[i],
		}
	}
	// 先按照分数排序，再按照年龄排序
	sort.Slice(members, func(i, j int) bool {
		return members[i].score < members[j].score || (members[i].score == members[j].score && members[i].age < members[j].age)
	})
	//fmt.Println(members)
	dp := make([]int, len(scores))
	dp[0] = members[0].score
	res := dp[0]
	for i := 1; i < len(scores); i++ {
		for j := i - 1; j >= 0; j-- {
			// 找到所有年龄小于或者等于自己的，最大的dp
			if members[j].age <= members[i].age {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i]+=members[i].score
		res = max(dp[i], res)
	}
	//fmt.Println(dp)
	return res
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

//leetcode submit region end(Prohibit modification and deletion)
