//你的赛车可以从位置 0 开始，并且速度为 +1 ，在一条无限长的数轴上行驶。赛车也可以向负方向行驶。赛车可以按照由加速指令 'A' 和倒车指令 'R' 组成
//的指令序列自动行驶。
//
//
// 当收到指令 'A' 时，赛车这样行驶：
//
// position += speed
// speed *= 2
//
// 当收到指令 'R' 时，赛车这样行驶：
//
// 如果速度为正数，那么speed = -1
// 否则 speed = 1
// 当前所处位置不变。
//
//
// 例如，在执行指令 "AAR" 后，赛车位置变化为 0 --> 1 --> 3 --> 3 ，速度变化为 1 --> 2 --> 4 --> -1 。
//
// 给你一个目标位置 target ，返回能到达目标位置的最短指令序列的长度。
//
//
//
// 示例 1：
//
//
//输入：target = 3
//输出：2
//解释：
//最短指令序列是 "AA" 。
//位置变化 0 --> 1 --> 3 。
//
//
// 示例 2：
//
//
//输入：target = 6
//输出：5
//解释：
//最短指令序列是 "AAARA" 。
//位置变化 0 --> 1 --> 3 --> 7 --> 7 --> 6 。
//
//
//
//
// 提示：
//
//
// 1 <= target <= 10⁴
//
//
// Related Topics 动态规划 👍 143 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func racecar(target int) int {
	// dp，记录到target的结果
	dp := make([]int, target+1)
	// 记录一直加速的情况下，可以达到的距离，加速1次，2次，3次。。。到达的距离
	all := make([]int, target)
	position, speed, number := 0, 1, 0
	// 首先记录可以一直加速到的地方，这些点的次数最小就是一直加速
	for ; position < target; position, speed, number = position+speed, speed*2, number+1 {
		dp[position] = number
		all[number] = position
	}
	// 继续加速的下一个地方，就是position
	// 如果是终点，则直接返回
	if position == target {
		return number
	}

	// 开始便利dp，从2开始
	for i := 2; i < target+1; i++ {
		if dp[i] == 0 {
			// 从第2个位置开始，走到第i个位置时
			// 要么是加速到k+1次之后，往回走，也是重复子问题
			var tmpNumber int
			for j := i; j < target; j++ {
				// 这种情况是走到终点之前可以找到往回走的点
				if dp[j] > 0 {
					dp[i] = dp[j] + 1 + dp[j-i]
					tmpNumber = dp[j]
					break
				}
			}
			// 要么走到终点找不到往回走的点，此时就是在position位置返回
			if dp[i] == 0 {
				dp[i] = number + 1 + dp[position-i]
				tmpNumber = number
			}
			//要么是加速k-1次之后，往回走，往回走0次、1次、2次，然后重复子问题。至于说往回走多少步，就需要便利
			tmpNumber--
			// k-1次，其实就是tmpNumber-1
			// 往回走0步，一直到往回走最大步数，也就是j-1
			for m := 0; m < tmpNumber; m++ {
				dp[i] = min(dp[i], dp[all[tmpNumber]]+1+m+1+dp[i-all[tmpNumber]+all[m]])
			}
		}
	}

	return dp[target]
}

func min(number1 int, number2 int) int {
	if number1 < number2 {
		return number1
	}
	return number2
}

//leetcode submit region end(Prohibit modification and deletion)
