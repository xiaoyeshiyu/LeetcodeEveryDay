//给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
//
//
//
// 示例1：
//
//
//
//
//输入: root = [1,3,2,5,3,null,9]
//输出: [1,3,9]
//
//
// 示例2：
//
//
//输入: root = [1,2,3]
//输出: [1,3]
//
//
//
//
// 提示：
//
//
// 二叉树的节点个数的范围是 [0,10⁴]
//
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 291 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		newQueue := make([]*TreeNode, 0)
		max := queue[0].Val
		for i := 0; i < len(queue); i++ {
			if queue[i].Val > max {
				max = queue[i].Val
			}
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue,  queue[i].Right)
			}
		}
		res = append(res, max)
		queue = newQueue
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
