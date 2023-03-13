// 给你 n 个城市，编号为从 1 到 n 。同时给你一个大小为 n-1 的数组 edges ，其中 edges[i] = [ui, vi] 表示城市 ui 和
// vi 之间有一条双向边。题目保证任意城市之间只有唯一的一条路径。换句话说，所有城市形成了一棵 树 。
//
// 一棵 子树 是城市的一个子集，且子集中任意城市之间可以通过子集中的其他城市和边到达。两个子树被认为不一样的条件是至少有一个城市在其中一棵子树中存在，但在另
// 一棵子树中不存在。
//
// 对于 d 从 1 到 n-1 ，请你找到城市间 最大距离 恰好为 d 的所有子树数目。
//
// 请你返回一个大小为 n-1 的数组，其中第 d 个元素（下标从 1 开始）是城市间 最大距离 恰好等于 d 的子树数目。
//
// 请注意，两个城市间距离定义为它们之间需要经过的边的数目。
//
// 示例 1：
//
// 输入：n = 4, edges = [[1,2],[2,3],[2,4]]
// 输出：[3,4,0]
// 解释：
// 子树 {1,2}, {2,3} 和 {2,4} 最大距离都是 1 。
// 子树 {1,2,3}, {1,2,4}, {2,3,4} 和 {1,2,3,4} 最大距离都为 2 。
// 不存在城市间最大距离为 3 的子树。
//
// 示例 2：
//
// 输入：n = 2, edges = [[1,2]]
// 输出：[1]
//
// 示例 3：
//
// 输入：n = 3, edges = [[1,2],[2,3]]
// 输出：[2,1]
//
// 提示：
//
// 2 <= n <= 15
// edges.length == n-1
// edges[i].length == 2
// 1 <= ui, vi <= n
// 题目保证 (ui, vi) 所表示的边互不相同。
//
// Related Topics 位运算 树 动态规划 状态压缩 枚举 👍 112 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	// 建树，m代表城市编号，n是这个城市连接的编号，用于后面求解最长路径时快速便利
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0]-1, e[1]-1 // 编号改为从 0 开始
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// inSet保存要遍历的城市的个数
	// vis保存便利的时候已经遍历过的城市，不会出现反过头成环
	var inSet, vis [15]bool

	// dfs 求树的直径，使用递归dfs，返回对应节点x的最深深度，调用时，需要重置diameter
	var dfs func(int) int
	dfs = func(x int) (maxLen int) {
		vis[x] = true
		for _, y := range g[x] {
			if !vis[y] && inSet[y] {
				ml := dfs(y) + 1
				if ml > maxLen {
					maxLen = ml
				}
			}
		}
		return
	}

	ans := make([]int, n-1)
	// 方法1：递归选择城市
	//var f func(int)
	//f = func(i int) {
	//	if i == n {
	//		for v, b := range inSet {
	//			if b {
	//				vis, diameter = [15]bool{}, 0
	//				dfs(v)
	//				break
	//			}
	//		}
	//		if diameter > 0 && vis == inSet {
	//			ans[diameter-1]++
	//		}
	//		return
	//	}
	//
	//	// 不选城市 i
	//	f(i + 1)
	//
	//	// 选城市 i
	//	inSet[i] = true
	//	f(i + 1)
	//	inSet[i] = false // 恢复现场
	//}
	//f(0)

	// 方法2：使用二进制，选择城市
	for i := 0; i < 1<<n; i++ {
		var count int
		inSet = [15]bool{}
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				inSet[j] = true
				count++
			}
		}

		// 最起码有两个城市，才计算距离
		if count > 1 {
			// 通过for循环每一个城市，确定直径

			// 初始化最大距离
			var diameter int
			for j := 0; j < n; j++ {
				if inSet[j] {
					// 初始化便利过的节点
					vis = [15]bool{}
					tmp := dfs(j)
					if tmp > diameter && vis == inSet {
						diameter = tmp
					}
				}
			}

			// 要确保所有节点都便利到，因此vis需要等于inSet
			if diameter > 0 {
				ans[diameter-1]++
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
