//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
//
// Related Topics 数组 回溯 👍 2353 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 1 {
		res = append(res, nums)
		return res
	}
	last := permute(nums[1:])
	for i := 0; i < len(last); i++ {
		tmp := make([]int, len(last[i]))
		copy(tmp, last[i])
		for j := 0; j < len(tmp)+1; j++ {
			newNum := make([]int, len(tmp)+1)
			if j == 0 {
				newNum[0] = nums[0]
				copy(newNum[1:], tmp[:])
				res = append(res, newNum)
				continue
			}
			copy(newNum, tmp[:j])
			newNum[j] = nums[0]
			if j != len(tmp) {
				copy(newNum[j+1:], tmp[j:])
			}
			res = append(res, newNum)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
