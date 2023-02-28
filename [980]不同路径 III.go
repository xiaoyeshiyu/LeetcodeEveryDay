//在二维网格 grid 上，有 4 种类型的方格：
//
//
// 1 表示起始方格。且只有一个起始方格。
// 2 表示结束方格，且只有一个结束方格。
// 0 表示我们可以走过的空方格。
// -1 表示我们无法跨越的障碍。
//
//
// 返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目。
//
// 每一个无障碍方格都要通过一次，但是一条路径中不能重复通过同一个方格。
//
//
//
// 示例 1：
//
// 输入：[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
//输出：2
//解释：我们有以下两条路径：
//1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
//2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
//
// 示例 2：
//
// 输入：[[1,0,0,0],[0,0,0,0],[0,0,0,2]]
//输出：4
//解释：我们有以下四条路径：
//1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
//2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
//3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
//4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
//
// 示例 3：
//
// 输入：[[0,1],[2,0]]
//输出：0
//解释：
//没有一条路能完全穿过每一个空的方格一次。
//请注意，起始和结束方格可以位于网格中的任意位置。
//
//
//
//
// 提示：
//
//
// 1 <= grid.length * grid[0].length <= 20
//
//
// Related Topics 位运算 数组 回溯 矩阵 👍 223 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsIII(grid [][]int) int {
	total := 1
	m, n := len(grid), len(grid[0])
	startm, startn, endm, endn := 0, 0, 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case 1:
				startm, startn = i,j
			case 2:
				endm, endn = i,j
			case 0:
				total++
			}
		}
	}
	return dfs(grid, startm, startn, endm, endn, total)
}

// DFS+回溯
func dfs(grid [][]int, startm, startn, endm, endn, total int) (res int) {
	//	先判断是否超过
	m, n := len(grid), len(grid[0])
	if startm < 0 || startm >= m || startn < 0 || startn >= n {
		return
	}

	// 判断当前是否是障碍
	if grid[startm][startn] == -1 {
		return
	}

	// 判断当前障碍是否是终点
	if grid[startm][startn] == 2 {
		// 如果是终点，判断是否一个空的格子都没有
		if total == 0 {
			return 1
		}
		// 如果还有空格子，就回溯
		return 0
	}

	//都不是，则将四个方向继续DFS
	grid[startm][startn] = -1
	res += dfs(grid, startm-1, startn, endm, endn, total-1)
	res += dfs(grid, startm, startn+1, endm, endn, total-1)
	res += dfs(grid, startm+1, startn, endm, endn, total-1)
	res += dfs(grid, startm, startn-1, endm, endn, total-1)

	// 回溯，将状态改回来
	grid[startm][startn] = 0
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
