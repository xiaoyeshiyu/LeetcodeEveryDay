//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -10⁴ <= matrix[i][j], target <= 10⁴
//
//
// Related Topics 数组 二分查找 矩阵 👍 750 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	lengthM := len(matrix)
	startM, endM := 0, lengthM-1
	for startM <= endM {
		mid := (startM + endM) >> 1
		if matrix[mid][0] > target {
			endM = mid -1
		} else if matrix[mid][0] < target {
			startM = mid +1
		} else {
			return true
		}
	}
	if endM < 0 {
		return false
	}
	m := matrix[endM]
	fmt.Println(m)
	lengthN := len(m)
	startN, endN := 0, lengthN-1
	for startN <= endN {
		mid := (startN + endN) >> 1
		if m[mid] > target {
			endN = mid - 1
		} else if m[mid] < target {
			startN = mid + 1
		} else {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
