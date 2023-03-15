// 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//
// 示例 1：
//
// 输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
// ,["1","0","0","1","0"]]
// 输出：6
// 解释：最大矩形如上图所示。
//
// 示例 2：
//
// 输入：matrix = []
// 输出：0
//
// 示例 3：
//
// 输入：matrix = [["0"]]
// 输出：0
//
// 示例 4：
//
// 输入：matrix = [["1"]]
// 输出：1
//
// 示例 5：
//
// 输入：matrix = [["0","0"]]
// 输出：0
//
// 提示：
//
// rows == matrix.length
// cols == matrix[0].length
// 1 <= row, cols <= 200
// matrix[i][j] 为 '0' 或 '1'
//
// Related Topics 栈 数组 动态规划 矩阵 单调栈 👍 1463 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	var res int
	tmp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				tmp[j]++
			} else {
				tmp[j] = 0
			}
		}
		tmpR := getMax(tmp)
		if tmpR > res {
			res = tmpR
		}
	}

	return res
}

func getMax(array []int) int {
	stack := make([]int, 0, len(array)+1)
	// 预先将-1存入
	stack = append(stack, -1)
	var res int
	for i := 0; i < len(array); i++ {
		// 出栈的时候，自己是高，左边界是新栈顶，i是右边界
		for len(stack) > 1 && array[i] < array[stack[len(stack)-1]] {
			out := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tmp := (i - stack[len(stack)-1]-1) * array[out]
			if tmp > res {
				res = tmp
			}
		}
		//	入栈
		stack = append(stack, i)
	}
	// 最终将栈中的数据全部都取出来，直到存储到-1
	// 左边界是新栈顶，右边界是最大长度，高度是自己
	for len(stack) > 1 {
		out := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		tmp := (len(array) - stack[len(stack)-1]-1) * array[out]
		if tmp > res {
			res = tmp
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
