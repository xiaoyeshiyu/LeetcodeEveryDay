//给定一个非负整数数组 nums 和一个整数 m ，你需要将这个数组分成 m 个非空的连续子数组。
//
// 设计一个算法使得这 m 个子数组各自和的最大值最小。
//
//
//
// 示例 1：
//
//
//输入：nums = [7,2,5,10,8], m = 2
//输出：18
//解释：
//一共有四种方法将 nums 分割为 2 个子数组。
//其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
//因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
//
// 示例 2：
//
//
//输入：nums = [1,2,3,4,5], m = 2
//输出：9
//
//
// 示例 3：
//
//
//输入：nums = [1,4,4], m = 3
//输出：4
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 10⁶
// 1 <= m <= min(50, nums.length)
//
//
// Related Topics 贪心 数组 二分查找 动态规划 👍 768 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func splitArray(nums []int, k int) int {
	length := len(nums)
	// 定义dp，二维数组
	// 第一维，代表份数
	// 第二维，代表数字长度
	// 值为对应份数和长度下，最大连续数组的和
	dp := make([][]int, k)

	// 初始化，第一个数字，份数为1的时候，此时的最大值，也就是直接相加
	dp[0] = make([]int, length)
	dp[0][0] = nums[0]
	for i := 1; i < length; i++ {
		dp[0][i] = dp[0][i-1] + nums[i]
	}
	// 从1份开始便利
	for i := 1; i < k && i < length; i++ {
		dp[i] = make([]int, length)
		//	长度也从1开始
		for j := i; j < length; j++ {
			allMax := dp[0][length-1]
			// 从0长度开始计算，一直到本身长度-1
			for m := i - 1; m < j; m++ {
				sub:= dp[0][j] - dp[0][m]
				// 先求出最大值
				t := max(dp[i-1][m], sub)
				// 最大值求出之后，再求出其中的最小值
				if t < allMax {
					allMax = t
				}
			}
			//	最终保存的是最小值
			dp[i][j] = allMax
		}
		//	然后在这些最大值中，求出最小值
	}
	//fmt.Println(dp)
	return dp[k-1][length-1]
}

func max(i int, sub int) int {
	if i > sub {
		return i
	}
	return sub
}

//leetcode submit region end(Prohibit modification and deletion)
