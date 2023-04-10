// 作为项目经理，你规划了一份需求的技能清单 req_skills，并打算从备选人员名单 people 中选出些人组成一个「必要团队」（ 编号为 i 的备选人员
// people[i] 含有一份该备选人员掌握的技能列表）。
//
// 所谓「必要团队」，就是在这个团队中，对于所需求的技能列表 req_skills 中列出的每项技能，团队中至少有一名成员已经掌握。可以用每个人的编号来表示团
// 队中的成员：
//
// 例如，团队 team = [0, 1, 3] 表示掌握技能分别为 people[0]，people[1]，和 people[3] 的备选人员。
//
// 请你返回 任一 规模最小的必要团队，团队成员用人员编号表示。你可以按 任意顺序 返回答案，题目数据保证答案存在。
//
// 示例 1：
//
// 输入：req_skills = ["java","nodejs","reactjs"], people = [["java"],["nodejs"],[
// "nodejs","reactjs"]]
// 输出：[0,2]
//
// 示例 2：
//
// 输入：req_skills = ["algorithms","math","java","reactjs","csharp","aws"], people
// = [["algorithms","math","java"],["algorithms","math","reactjs"],["java",
// "csharp","aws"],["reactjs","csharp"],["csharp","math"],["aws","java"]]
// 输出：[1,2]
//
// 提示：
//
// 1 <= req_skills.length <= 16
// 1 <= req_skills[i].length <= 16
// req_skills[i] 由小写英文字母组成
// req_skills 中的所有字符串 互不相同
// 1 <= people.length <= 60
// 0 <= people[i].length <= 16
// 1 <= people[i][j].length <= 16
// people[i][j] 由小写英文字母组成
// people[i] 中的所有字符串 互不相同
// people[i] 中的每个技能是 req_skills 中的技能
// 题目数据保证「必要团队」一定存在
//
// Related Topics 位运算 数组 动态规划 状态压缩 👍 156 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func smallestSufficientTeam(reqSkills []string, people [][]string) (ans []int) {
	// 将人员的能力压缩成二进制
	peopleSkill := make([]int, len(people))
	for i, p := range people {
		var tmp int
		for _, skill := range p {
			for k, reqSkill := range reqSkills {
				if reqSkill == skill {
					tmp += 1 << k
					break
				}
			}
		}
		peopleSkill[i] = tmp
	}
	//fmt.Println(peopleSkill)

	dp := make([][]int, 1<<len(reqSkills))
	// 便利所有人的技能
	for i := 0; i < len(peopleSkill); i++ {
		// 便利已经存在的技能
		for j := len(dp) - 1; j >=0; j-- {
			if len(dp[j]) == 0 {
				continue
			}
			// 这个人的技能和已存在的技能合并
			tmp := j | peopleSkill[i]
			// 新的技能集合不存在
			// 或者存在，但是长度太大
			if dp[tmp] == nil || len(dp[j])+1 < len(dp[tmp]) {
				// golang注意，append的时候，可能会影响原数组，因此要copy
				dp[tmp] = make([]int,len(dp[j]))
				copy(dp[tmp],dp[j])
				dp[tmp] = append(dp[tmp], i)
			}
		}
		if len(dp[peopleSkill[i]]) != 1 {
			dp[peopleSkill[i]] = []int{i}
		}
	}

	return dp[(1<<len(reqSkills))-1]
}

//leetcode submit region end(Prohibit modification and deletion)
