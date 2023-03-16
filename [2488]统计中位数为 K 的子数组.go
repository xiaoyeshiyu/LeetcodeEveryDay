//给你一个长度为 n 的数组 nums ，该数组由从 1 到 n 的 不同 整数组成。另给你一个正整数 k 。 
//
// 统计并返回 nums 中的 中位数 等于 k 的非空子数组的数目。 
//
// 注意： 
//
// 
// 数组的中位数是按 递增 顺序排列后位于 中间 的那个元素，如果数组长度为偶数，则中位数是位于中间靠 左 的那个元素。 
// 
//
// 
// 例如，[2,3,1,4] 的中位数是 2 ，[8,4,3,5,1] 的中位数是 4 。 
// 
// 
// 子数组是数组中的一个连续部分。 
//
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [3,2,1,4,5], k = 4
//输出：3
//解释：中位数等于 4 的子数组有：[4]、[4,5] 和 [1,4,5] 。
// 
//
// 示例 2： 
//
// 
//输入：nums = [2,3,1], k = 3
//输出：1
//解释：[3] 是唯一一个中位数等于 3 的子数组。
// 
//
// 
//
// 提示： 
//
// 
// n == nums.length 
// 1 <= n <= 10⁵ 
// 1 <= nums[i], k <= n 
// nums 中的整数互不相同 
// 
//
// Related Topics 数组 哈希表 前缀和 👍 61 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func countSubarrays(nums []int, k int) int {
	var flag int
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			flag = i
			break
		}
	}

	count := make(map[int]int) // 计数器
	// 保留一个等于自身的值
	count[0] = 1
	var res,c int
	for i := 0; i < len(nums); i++ {
		c += getCount(nums[i],k)
		if i < flag {
			count[c]++
		} else {
			// 结果为0，代表是相加是奇数的情况
			p0 := count[c]		// 等于对应数可以
			p1 := count[c-1]	// 等于对应数-1也可以
			res += p0+p1
		}
	}
	return res
}

func getCount(i int,k int) int  {
	if i > k {
		return 1
	} else if i < k {
		return -1
	}
	return 0
}
//leetcode submit region end(Prohibit modification and deletion)
