//给你一个整数数组 arr，请你将该数组分隔为长度 最多 为 k 的一些（连续）子数组。分隔完成后，每个子数组的中的所有值都会变为该子数组中的最大值。 
//
// 返回将数组分隔变换后能够得到的元素最大和。本题所用到的测试用例会确保答案是一个 32 位整数。 
//
// 
//
// 示例 1： 
//
// 
//输入：arr = [1,15,7,9,2,5,10], k = 3
//输出：84\]wq
//解释：数组变为 [15,15,15,9,10,10,10] 
//
// 示例 2： 
//
// 
//输入：arr = [1,4,1,5,7,3,6,1,9,9,3], k = 4
//输出：83
// 
//
// 示例 3： 
//
// 
//输入：arr = [1], k = 1
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= arr.length <= 500 
// 0 <= arr[i] <= 10⁹ 
// 1 <= k <= arr.length 
// 
//
// Related Topics 数组 动态规划 👍 229 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func maxSumAfterPartitioning(arr []int, k int) int {
	l := len(arr)
	// 使用DP，代表当数组长度为i时，元素最大和
	dp := make([]int,l+1)
	//  便利数组长度
	for i := 1; i <= l; i++ {
		tmpMax := arr[i-1]
		dp[i] =  dp[i-1] + arr[i-1]
		// 便利子数组长度，从0到k
		for j := 1; j < k && i-j > 0 ; j++ {
			tmpMax = max(tmpMax,arr[i-j-1])
			dp[i] = max(dp[i],dp[i-j-1] + tmpMax*(j+1))
		}
	}
	return dp[l]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
