//给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充
//。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O",
//"X","X"]]
//输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
//解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都
//会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
//
//
// 示例 2：
//
//
//输入：board = [["X"]]
//输出：[["X"]]
//
//
//
//
// 提示：
//
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 200
// board[i][j] 为 'X' 或 'O'
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 914 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
var m, n int

func solve(board [][]byte) {
	m, n = len(board), len(board[0])
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0, board)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1, board)
		}
	}
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			dfs(0, i, board)
		}
		if board[m-1][i] == 'O' {
			dfs(m-1, i, board)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O'+1 {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(i int, j int, board [][]byte) {
	if i < 0 || i >= m ||
		j < 0 || j >= n ||
		board[i][j] != 'O' {
		return
	}
	board[i][j] = 'O' + 1
	dfs(i-1, j, board)
	dfs(i, j+1, board)
	dfs(i+1, j, board)
	dfs(i, j-1, board)
}

//leetcode submit region end(Prohibit modification and deletion)
