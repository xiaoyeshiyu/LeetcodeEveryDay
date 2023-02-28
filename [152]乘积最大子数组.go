//给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
// 测试用例的答案是一个 32-位 整数。
//
// 子数组 是数组的连续子序列。
//
//
//
// 示例 1:
//
//
//输入: nums = [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//
//
// 示例 2:
//
//
//输入: nums = [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 2 * 10⁴
// -10 <= nums[i] <= 10
// nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数
//
//
// Related Topics 数组 动态规划 👍 1885 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	dp := make([][2]int, length)
	if nums[0] > 0 {
		dp[0][0] = nums[0]
	} else {
		dp[0][1] = nums[0]
	}

	for i := 1; i < length; i++ {
		dp[i][0]= maxInt(dp[i-1][0]*nums[i], dp[i-1][1]*nums[i], nums[i])
		dp[i][1] =minInt(dp[i-1][0]*nums[i], dp[i-1][1]*nums[i], nums[i])
	}
	fmt.Println(dp)
	res := 0
	for i := 0; i < length; i++ {
		if res < dp[i][0] {
			res = dp[i][0]
		}
	}

	return res
}

func maxInt(a, b, c int) int {
	if b > a {
		a = b
	}
	if c > a {
		a = c
	}
	return a
}

func minInt(a, b, c int) int {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
