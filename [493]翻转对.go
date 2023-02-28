//给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。
//
// 你需要返回给定数组中的重要翻转对的数量。
//
// 示例 1:
//
//
//输入: [1,3,2,3,1]
//输出: 2
//
//
// 示例 2:
//
//
//输入: [2,4,3,5,1]
//输出: 3
//
//
// 注意:
//
//
// 给定数组的长度不会超过50000。
// 输入数组中的所有数字都在32位整数的表示范围内。
//
//
// Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 397 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func reversePairs(nums []int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > 2*nums[j] {
				res++
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
