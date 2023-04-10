// 给定一个长度为 n 的链表 head
//
// 对于列表中的每个节点，查找下一个 更大节点 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 严格大于 它的值。
//
// 返回一个整数数组 answer ，其中 answer[i] 是第 i 个节点( 从1开始 )的下一个更大的节点的值。如果第 i 个节点没有下一个更大的节点
// ，设置 answer[i] = 0 。
//
// 示例 1：
//
// 输入：head = [2,1,5]
// 输出：[5,5,0]
//
// 示例 2：
//
// 输入：head = [2,7,4,3,5]
// 输出：[7,0,5,5,0]
//
// 提示：
//
// 链表中节点数为 n
// 1 <= n <= 10⁴
// 1 <= Node.val <= 10⁹
//
// Related Topics 栈 数组 链表 单调栈 👍 252 👎 0
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func nextLargerNodes(head *ListNode) []int {
	// 单调递减栈和结果
	type node struct {
		i   int
		max int
	}

	res := []int{0}
	stack := []node{{i: 0, max: head.Val}}
	start := 1
	head = head.Next
	for head != nil {
		for len(stack) > 0 && head.Val > stack[len(stack)-1].max {
			res[stack[len(stack)-1].i] = head.Val
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, node{
			i:   start,
			max: head.Val,
		})
		start++
		res = append(res, 0)
		head = head.Next
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
