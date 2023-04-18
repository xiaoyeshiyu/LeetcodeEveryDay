// 给定二叉树的根节点 root，找出存在于 不同 节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。
//
// （如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）
//
// 示例 1：
//
// 输入：root = [8,3,10,1,6,null,14,null,null,4,7,13]
// 输出：7
// 解释：
// 我们有大量的节点与其祖先的差值，其中一些如下：
// |8 - 3| = 5
// |3 - 7| = 4
// |8 - 1| = 7
// |10 - 13| = 3
// 在所有可能的差值中，最大值 7 由 |8 - 1| = 7 得出。
//
// 示例 2：
//
// 输入：root = [1,null,2,null,0,3]
// 输出：3
//
// 提示：
//
// 树中的节点数在 2 到 5000 之间。
// 0 <= Node.val <= 10⁵
//
// Related Topics 树 深度优先搜索 二叉树 👍 150 👎 0
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(i, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

func min(i, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dfs(node *TreeNode, ma int, mi int) int {
	if node == nil {
		return 0
	}
	ma, mi = max(node.Val, ma), min(node.Val, mi)
	res := max(abs(node.Val-ma), abs(node.Val-mi))
	res = max(res,dfs(node.Left,ma,mi))
	res = max(res,dfs(node.Right,ma,mi))
	return res
}

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

//leetcode submit region end(Prohibit modification and deletion)
