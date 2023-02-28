//一个机器人位于一个
// m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
//
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//
//
// 示例 1：
//
//
//输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//输出：2
//解释：3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
//
//
// 示例 2：
//
//
//输入：obstacleGrid = [[0,1],[0,0]]
//输出：1
//
//
//
//
// 提示：
//
//
// m == obstacleGrid.length
// n == obstacleGrid[i].length
// 1 <= m, n <= 100
// obstacleGrid[i][j] 为 0 或 1
//
//
// Related Topics 数组 动态规划 矩阵 👍 946 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	// dp定义，二维数组，代表当前节点能够走到的方式
	length := len(obstacleGrid)
	dp := make([][]int, length)
	tmpl := len(obstacleGrid[0])

	// 初始化条件，也就是初始化左边界和上边界，需要注意石头的位置，出现石头则不能再往右走或者往下走，更右边的或者更下面的，都是0
	tmp := make([]int, tmpl)
	flag := false
	for i := 0; i < tmpl; i++ {
		if obstacleGrid[0][i] == 1 {
			flag = true
		}
		if flag {
			tmp[i] = 0
		} else {
			tmp[i] = 1
		}
	}
	dp[0] = tmp

	flag = false
	for i := 1; i < length; i++ {
		tmp = make([]int, tmpl)
		dp[i] = tmp
		if obstacleGrid[i][0] == 1 {
			flag = true
		}
		if flag {
			tmp[0] = 0
		} else {
			tmp[0] = 1
		}
	}

	// 条件转移方程，有石头的地方路线为0，其他的地方，依然是左边的+上面的
	for i := 1; i < length; i++ {
		for j := 1; j < tmpl; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[length-1][tmpl-1]
}

//leetcode submit region end(Prohibit modification and deletion)
