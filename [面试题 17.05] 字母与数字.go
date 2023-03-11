// 给定一个放有字母和数字的数组，找到最长的子数组，且包含的字母和数字的个数相同。
//
// 返回该子数组，若存在多个最长子数组，返回左端点下标值最小的子数组。若不存在这样的数组，返回一个空数组。
//
// 示例 1:
//
// 输入: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7","H","I","J","K",
// "L","M"]
//
// 输出: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7"]
//
// 示例 2:
//
// 输入: ["A","A"]
//
// 输出: []
//
// 提示：
//
// array.length <= 100000
//
// Related Topics 数组 哈希表 前缀和 👍 156 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func findLongestSubarray(array []string) []string {
	all := map[int]int{0: -1}

	var num, max int
	var res []string
	for i := 0; i < len(array); i++ {
		if unicode.IsLetter(rune(array[i][0])) {
			num++
		} else {
			num--
		}
		if v, ok := all[num]; ok {
			tmp := i - v
			if tmp > max {
				max = tmp
				res = array[v+1 : i+1]
			}
		} else {
			all[num] =  i
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
