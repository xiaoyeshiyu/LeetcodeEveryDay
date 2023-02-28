//爱丽丝和鲍勃继续他们的石子游戏。许多堆石子 排成一行，每堆都有正整数颗石子 piles[i]。游戏以谁手中的石子最多来决出胜负。
//
// 爱丽丝和鲍勃轮流进行，爱丽丝先开始。最初，M = 1。
//
// 在每个玩家的回合中，该玩家可以拿走剩下的 前 X 堆的所有石子，其中 1 <= X <= 2M。然后，令 M = max(M, X)。
//
// 游戏一直持续到所有石子都被拿走。
//
// 假设爱丽丝和鲍勃都发挥出最佳水平，返回爱丽丝可以得到的最大数量的石头。
//
//
//
// 示例 1：
//
//
//输入：piles = [2,7,9,4,4]
//输出：10
//解释：如果一开始Alice取了一堆，Bob取了两堆，然后Alice再取两堆。爱丽丝可以得到2 + 4 + 4 = 10堆。如果Alice一开始拿走了两堆，那
//么Bob可以拿走剩下的三堆。在这种情况下，Alice得到2 + 7 = 9堆。返回10，因为它更大。
//
//
// 示例 2:
//
//
//输入：piles = [1,2,3,4,5,100]
//输出：104
//
//
//
//
// 提示：
//
//
// 1 <= piles.length <= 100
//
// 1 <= piles[i] <= 10⁴
//
//
// Related Topics 数组 数学 动态规划 博弈 👍 226 👎 0
package main

//    A     B
// M: 1     2
// X: 1-2   1-4

//leetcode submit region begin(Prohibit modification and deletion)
func stoneGameII(piles []int) int {
	n := len(piles)
	// 二维数组，记录当A拿第一堆时，第i堆被A拿的时候对应的M的最大石头数量
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	// 从最后一堆开始往前循环
	for i, s := n-1, 0; i >= 0; i-- {
		s += piles[i]
		// M 从1开始循环，直到 i/2+1，也就是假设可以一次性拿一半，M最大是已经拿了的一倍
		for m := 1; m <= i/2+1; m++ {
			// 状态转移方程，可以全部都拿，就全部都拿
			if i+m*2 >= n {
				f[i][m] = s
			} else {
				// 不能全部都拿
				mn := math.MaxInt
				// 便利M，获取B拿最少的
				for x := 1; x <= m*2; x++ {
					mn = min(mn, f[i+x][max(m, x)])
				}
				// 剩下的就是自己最多的
				f[i][m] = s - mn
			}
		}
	}
	return f[0][1]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
