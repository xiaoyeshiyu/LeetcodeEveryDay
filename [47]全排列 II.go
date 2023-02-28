//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
//
// Related Topics 数组 回溯 👍 1258 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 1 {
		res = append(res, nums)
		return res
	}
	last := permuteUnique(nums[1:])
	for i := 0; i < len(last); i++ {
		tmp := make([]int, len(last[i]))
		copy(tmp, last[i])
	OUT:
		for j := 0; j < len(tmp)+1; j++ {
			newNum := make([]int, len(tmp)+1)
			if j == 0 {
				newNum[0] = nums[0]
				copy(newNum[1:], tmp[:])
				// 去重
				for k := 0; k < len(res); k++ {
					flag := false
					for l := 0; l < len(newNum); l++ {
						if res[k][l] != newNum[l] {
							flag = true
							break
						}
					}
					if !flag {
						continue OUT
					}
				}
				res = append(res, newNum)
				continue
			}
			copy(newNum, tmp[:j])
			newNum[j] = nums[0]
			if j != len(tmp) {
				copy(newNum[j+1:], tmp[j:])
			}
			// 去重
			for k := 0; k < len(res); k++ {
				flag := false
				for l := 0; l < len(newNum); l++ {
					if res[k][l] != newNum[l] {
						flag = true
						break
					}
				}
				if !flag {
					continue OUT
				}
			}
			res = append(res, newNum)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
