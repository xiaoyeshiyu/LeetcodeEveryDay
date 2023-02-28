//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返
//回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//
//
//
// 示例 1：
//
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2：
//
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//
//
// 提示：
//
//
// 1 <= intervals.length <= 10⁴
// intervals[i].length == 2
// 0 <= starti <= endi <= 10⁴
//
//
// Related Topics 数组 排序 👍 1789 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	var change bool
	res := make([][]int, 0, len(intervals))
OUT:
	for _, interval := range intervals {
		for i := 0; i < len(res); i++ {
			//	 3 种情况
			if (interval[0] >= res[i][0] && interval[0] <= res[i][1]) ||
				(interval[1] >= res[i][0] && interval[1] <= res[i][1]) ||
				(interval[1] >= res[i][1] && interval[0] <= res[i][0]) {
				res[i] = []int{min(res[i][0], interval[0]), max(res[i][1], interval[1])}
				change = true
				continue OUT
			}
		}
		res = append(res, interval)
	}
	if change {
		return merge(res)
	}
	return res
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

func min(i int, i2 int) int {
	if i > i2 {
		return i2
	}
	return i
}

//leetcode submit region end(Prohibit modification and deletion)
