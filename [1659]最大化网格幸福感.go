//给你四个整数 m、n、introvertsCount 和 extrovertsCount 。有一个 m x n 网格，和两种类型的人：内向的人和外向的人。总
//共有 introvertsCount 个内向的人和 extrovertsCount 个外向的人。 
//
// 请你决定网格中应当居住多少人，并为每个人分配一个网格单元。 注意，不必 让所有人都生活在网格中。 
//
// 每个人的 幸福感 计算如下： 
//
// 
// 内向的人 开始 时有 120 个幸福感，但每存在一个邻居（内向的或外向的）他都会 失去 30 个幸福感。 
// 外向的人 开始 时有 40 个幸福感，每存在一个邻居（内向的或外向的）他都会 得到 20 个幸福感。 
// 
//
// 邻居是指居住在一个人所在单元的上、下、左、右四个直接相邻的单元中的其他人。 
//
// 网格幸福感 是每个人幸福感的 总和 。 返回 最大可能的网格幸福感 。 
//
// 
//
// 示例 1： 
// 
// 
//输入：m = 2, n = 3, introvertsCount = 1, extrovertsCount = 2
//输出：240
//解释：假设网格坐标 (row, column) 从 1 开始编号。
//将内向的人放置在单元 (1,1) ，将外向的人放置在单元 (1,3) 和 (2,3) 。
//- 位于 (1,1) 的内向的人的幸福感：120（初始幸福感）- (0 * 30)（0 位邻居）= 120
//- 位于 (1,3) 的外向的人的幸福感：40（初始幸福感）+ (1 * 20)（1 位邻居）= 60
//- 位于 (2,3) 的外向的人的幸福感：40（初始幸福感）+ (1 * 20)（1 位邻居）= 60
//网格幸福感为：120 + 60 + 60 = 240
//上图展示该示例对应网格中每个人的幸福感。内向的人在浅绿色单元中，而外向的人在浅紫色单元中。
// 
//
// 示例 2： 
//
// 
//输入：m = 3, n = 1, introvertsCount = 2, extrovertsCount = 1
//输出：260
//解释：将内向的人放置在单元 (1,1) 和 (3,1) ，将外向的人放置在单元 (2,1) 。
//- 位于 (1,1) 的内向的人的幸福感：120（初始幸福感）- (1 * 30)（1 位邻居）= 90
//- 位于 (2,1) 的外向的人的幸福感：40（初始幸福感）+ (2 * 20)（2 位邻居）= 80
//- 位于 (3,1) 的内向的人的幸福感：120（初始幸福感）- (1 * 30)（1 位邻居）= 90
//网格幸福感为 90 + 80 + 90 = 260
// 
//
// 示例 3： 
//
// 
//输入：m = 2, n = 2, introvertsCount = 4, extrovertsCount = 0
//输出：240
// 
//
// 
//
// 提示： 
//
// 
// 1 <= m, n <= 5 
// 0 <= introvertsCount, extrovertsCount <= min(m * n, 6) 
// 
//
// Related Topics 位运算 记忆化搜索 动态规划 状态压缩 👍 106 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int {
	mx := int(math.Pow(3, float64(n)))
	f := make([]int, mx)
	g := make([][]int, mx)
	h := [3][3]int{{0, 0, 0}, {0, -60, -10}, {0, -10, 40}}
	bits := make([][]int, mx)
	ix := make([]int, mx)
	ex := make([]int, mx)
	memo := make([][][][]int, m)
	for i := range g {
		g[i] = make([]int, mx)
		bits[i] = make([]int, n)
	}
	for i := range memo {
		memo[i] = make([][][]int, mx)
		for j := range memo[i] {
			memo[i][j] = make([][]int, introvertsCount+1)
			for k := range memo[i][j] {
				memo[i][j][k] = make([]int, extrovertsCount+1)
				for l := range memo[i][j][k] {
					memo[i][j][k][l] = -1
				}
			}
		}
	}
	for i := 0; i < mx; i++ {
		mask := i
		for j := 0; j < n; j++ {
			x := mask % 3
			mask /= 3
			bits[i][j] = x
			if x == 1 {
				ix[i]++
				f[i] += 120
			} else if x == 2 {
				ex[i]++
				f[i] += 40
			}
			if j > 0 {
				f[i] += h[x][bits[i][j-1]]
			}
		}
	}
	for i := 0; i < mx; i++ {
		for j := 0; j < mx; j++ {
			for k := 0; k < n; k++ {
				g[i][j] += h[bits[i][k]][bits[j][k]]
			}
		}
	}
	var dfs func(int, int, int, int) int
	dfs = func(i, pre, ic, ec int) int {
		if i == m || (ic == 0 && ec == 0) {
			return 0
		}
		if memo[i][pre][ic][ec] != -1 {
			return memo[i][pre][ic][ec]
		}
		ans := 0
		for cur := 0; cur < mx; cur++ {
			if ix[cur] <= ic && ex[cur] <= ec {
				ans = max(ans, f[cur]+g[pre][cur]+dfs(i+1, cur, ic-ix[cur], ec-ex[cur]))
			}
		}
		memo[i][pre][ic][ec] = ans
		return ans
	}
	return dfs(0, 0, introvertsCount, extrovertsCount)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
