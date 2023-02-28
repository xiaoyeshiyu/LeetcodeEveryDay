//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。
//
//
//
// 示例 1：
//
//
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
//
//
// 示例 2：
//
//
//输入：nums = [1]
//输出：1
//
//
// 示例 3：
//
//
//输入：nums = [5,4,-1,7,8]
//输出：23
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
//
// Related Topics 数组 分治 动态规划 👍 5591 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	dp := make([]int, length)
	dp[0] = nums[0]
	max := nums[0]
	for i := 1; i < length; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1]+nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	fmt.Println(dp)
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
