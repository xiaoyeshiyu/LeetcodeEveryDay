//有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
//
// 现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i -
// 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
//
// 求所能获得硬币的最大数量。
//
//
//示例 1：
//
//
//输入：nums = [3,1,5,8]
//输出：167
//解释：
//nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
//coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
//
// 示例 2：
//
//
//输入：nums = [1,5]
//输出：10
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 300
// 0 <= nums[i] <= 100
//
//
// Related Topics 数组 动态规划 👍 1159 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func maxCoins(nums []int) int {
	length := len(nums)
	// 组装新数组
	newLength := length + 2
	newNums := make([]int, newLength)
	copy(newNums[1:], nums)
	newNums[0], newNums[newLength-1] = 1, 1

	// 定义dp，二维数组，代表区间i，j之间，插入一个气球，得到的最大的硬币数量
	dp := make([][]int, newLength)

	// 初始化，计算出有一个数区间的数字，[0,2],[1,3],[2,4],[3,5]，计算这个区间之内，插入一个气球，得到的硬币数量
	for i := 0; i < newLength; i++ {
		dp[i] = make([]int, newLength)
		if i < newLength-2 {
			dp[i][i+2] = newNums[i] * newNums[i+1] * newNums[i+2]
		}
	}

	// 在[i,j]区间内，要想硬币最多，从后往前推导，也就是可以按照往i,j之间，往里面插入气球，计算硬币。
	// 这样的好处是可以将大问题划分为小问题，小问题之间不相互影响
	// 在[i,j]之间，插入一个气球，硬币最多，则是k从i+1到j-1的max(dp[i][k]+newNums[i]*newNums[k]*newNums[j]+dp[k][j])
	// 这个里面的子系统就是dp[i][k]，子系统的计算，不会影响dp[i][j]。相比计算戳破气球，会涉及到左边界或者右边界的改动，往里面插入的思想更好。
	// 从间隔3到间隔5
	for i := 3; i < newLength; i++ {
		// 从角标0开始，[0,3]
		for j := 0; j < newLength-i; j++ {
			// dp[j][j+i] =  max(xxx)
			tmpMax := 0
			// 开始计算中间的max，也就是j到j+i之间
			for m := j + 1; m < j+i; m++ {
				tmp := dp[j][m] + newNums[j]*newNums[m]*newNums[j+i] + dp[m][j+i]
				if tmp > tmpMax {
					tmpMax = tmp
				}
			}
			dp[j][j+i] = tmpMax
		}
	}

	return dp[0][newLength-1]
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

//leetcode submit region end(Prohibit modification and deletion)
