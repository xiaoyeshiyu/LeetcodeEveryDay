//一只青蛙想要过河。 假定河流被等分为若干个单元格，并且在每一个单元格内都有可能放有一块石子（也有可能没有）。 青蛙可以跳上石子，但是不可以跳入水中。
//
// 给你石子的位置列表 stones（用单元格序号 升序 表示）， 请判定青蛙能否成功过河（即能否在最后一步跳至最后一块石子上）。开始时， 青蛙默认已站在第一
//块石子上，并可以假定它第一步只能跳跃 1 个单位（即只能从单元格 1 跳至单元格 2 ）。
//
// 如果青蛙上一步跳跃了 k 个单位，那么它接下来的跳跃距离只能选择为 k - 1、k 或 k + 1 个单位。 另请注意，青蛙只能向前方（终点的方向）跳跃。
//
//
//
//
// 示例 1：
//
//
//输入：stones = [0,1,3,5,6,8,12,17]
//输出：true
//解释：青蛙可以成功过河，按照如下方案跳跃：跳 1 个单位到第 2 块石子, 然后跳 2 个单位到第 3 块石子, 接着 跳 2 个单位到第 4 块石子, 然
//后跳 3 个单位到第 6 块石子, 跳 4 个单位到第 7 块石子, 最后，跳 5 个单位到第 8 个石子（即最后一块石子）。
//
// 示例 2：
//
//
//输入：stones = [0,1,2,3,4,8,9,11]
//输出：false
//解释：这是因为第 5 和第 6 个石子之间的间距太大，没有可选的方案供青蛙跳跃过去。
//
//
//
// 提示：
//
//
// 2 <= stones.length <= 2000
// 0 <= stones[i] <= 2³¹ - 1
// stones[0] == 0
// stones 按严格升序排列
//
//
// Related Topics 数组 动态规划 👍 475 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func canCross(stones []int) bool {
	// 定义dp表达式，代表第i个石子能够被跳到时，跳的步数
	dp := make(map[int]map[int]bool)

	//	先便利一遍所有的石子，记录石子的位置
	for i := 1; i < len(stones); i++ {
		dp[stones[i]] = make(map[int]bool)
	}
	// 初始化，第1颗石子能够达到，并且是0步，这样可以确保第1颗石子可以1步达到
	dp[0] = map[int]bool{0: true}
	for i := 0; i < len(stones); i++ {
		point := stones[i]
		// 状态转移方程
		// 便利每一颗可以达到的石子的步数
		for k := range dp[point] {
			// 分别记录步数k-1、k、k+1，能够达到的石子，记录这个石子到达的时候的步数
			// 如果达到的石子是最后一个石子，则返回true，否则返回false
			if _, ok := dp[point+k-1]; ok && k-1 > 0 {
				if point+k-1 == stones[len(stones)-1] {
					return true
				}
				dp[point+k-1][k-1] = true
			}
			if _, ok := dp[point+k]; ok && k > 0 {
				if point+k == stones[len(stones)-1] {
					return true
				}
				dp[point+k][k] = true
			}
			if _, ok := dp[point+k+1]; ok && k+1 > 0 {
				if point+k+1 == stones[len(stones)-1] {
					return true
				}
				dp[point+k+1][k+1] = true
			}
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
