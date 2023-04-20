//给你两个整数数组 arr1 和 arr2，返回使 arr1 严格递增所需要的最小「操作」数（可能为 0）。 
//
// 每一步「操作」中，你可以分别从 arr1 和 arr2 中各选出一个索引，分别为 i 和 j，0 <= i < arr1.length 和 0 <= j 
//< arr2.length，然后进行赋值运算 arr1[i] = arr2[j]。 
//
// 如果无法让 arr1 严格递增，请返回 -1。 
//
// 
//
// 示例 1： 
//
// 输入：arr1 = [1,5,3,6,7], arr2 = [1,3,2,4]
//输出：1
//解释：用 2 来替换 5，之后 arr1 = [1, 2, 3, 6, 7]。
// 
//
// 示例 2： 
//
// 输入：arr1 = [1,5,3,6,7], arr2 = [4,3,1]
//输出：2
//解释：用 3 来替换 5，然后用 4 来替换 3，得到 arr1 = [1, 3, 4, 6, 7]。
// 
//
// 示例 3： 
//
// 输入：arr1 = [1,5,3,6,7], arr2 = [1,6,3,3]
//输出：-1
//解释：无法使 arr1 严格递增。 
//
// 
//
// 提示： 
//
// 
// 1 <= arr1.length, arr2.length <= 2000 
// 0 <= arr1[i], arr2[i] <= 10^9 
// 
//
// 
//
// Related Topics 数组 二分查找 动态规划 排序 👍 120 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)
	n := len(arr1)
	m := len(arr2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, min(m, n)+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = -1
	for i := 1; i <= len(arr1); i++ {
		for j := 0; j <= min(i, m); j++ {
			if arr1[i-1] > dp[i-1][j] {
				dp[i][j] = arr1[i-1]
			}
			if j > 0 && dp[i-1][j-1] != math.MaxInt {
				k := j - 1 + sort.SearchInts(arr2[j-1:], dp[i-1][j-1]+1)
				if k < m {
					dp[i][j] = min(dp[i][j], arr2[k])
				}
			}
			if i == n && dp[i][j] != math.MaxInt {
				return j
			}
		}
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
