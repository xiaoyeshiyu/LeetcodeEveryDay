// 给你一棵树，树上有 n 个节点，按从 0 到 n-1 编号。树以父节点数组的形式给出，其中 parent[i] 是节点 i 的父节点。树的根节点是编号为 0
// 的节点。
//
// 树节点的第 k 个祖先节点是从该节点到根节点路径上的第 k 个节点。
//
// 实现 TreeAncestor 类：
//
// TreeAncestor（int n， int[] parent） 对树和父数组中的节点数初始化对象。
// getKthAncestor(int node, int k) 返回节点 node 的第 k 个祖先节点。如果不存在这样的祖先节点，返回 -1 。
//
// 示例 1：
//
// 输入：
// ["TreeAncestor","getKthAncestor","getKthAncestor","getKthAncestor"]
// [[7,[-1,0,0,1,1,2,2]],[3,1],[5,2],[6,3]]
//
// 输出：
// [null,1,0,-1]
//
// 解释：
// TreeAncestor treeAncestor = new TreeAncestor(7, [-1, 0, 0, 1, 1, 2, 2]);
//
// treeAncestor.getKthAncestor(3, 1);  // 返回 1 ，它是 3 的父节点
// treeAncestor.getKthAncestor(5, 2);  // 返回 0 ，它是 5 的祖父节点
// treeAncestor.getKthAncestor(6, 3);  // 返回 -1 因为不存在满足要求的祖先节点
//
// 提示：
//
// 1 <= k <= n <= 5 * 10⁴
// parent[0] == -1 表示编号为 0 的节点是根节点。
// 对于所有的 0 < i < n ，0 <= parent[i] < n 总成立
// 0 <= node < n
// 至多查询 5 * 10⁴ 次
//
// Related Topics 树 深度优先搜索 广度优先搜索 设计 二分查找 动态规划 👍 139 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
const kLog = 16

type TreeAncestor struct {
	ancestors [][]int
}

func Constructor(n int, parent []int) TreeAncestor {
	var this TreeAncestor
	this.ancestors = make([][]int, n)
	for i := 0; i < n; i++ {
		this.ancestors[i] = make([]int, kLog)
		for j := 0; j < kLog; j++ {
			this.ancestors[i][j] = -1
		}
		this.ancestors[i][0] = parent[i]
	}
	for j := 1; j < kLog; j++ {
		for i := 0; i < n; i++ {
			if this.ancestors[i][j-1] != -1 {
				this.ancestors[i][j] = this.ancestors[this.ancestors[i][j-1]][j-1]
			}
		}
	}
	return this
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	for j := 0; j < kLog; j++ {
		if (k>>j)&1 != 0 {
			node = this.ancestors[node][j]
			if node == -1 {
				return -1
			}
		}
	}
	return node
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
//leetcode submit region end(Prohibit modification and deletion)
