//序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方
//式重构得到原数据。
//
// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串
//反序列化为原始的树结构。
//
// 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的
//方法解决这个问题。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3,null,null,4,5]
//输出：[1,2,3,null,null,4,5]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中结点数在范围 [0, 10⁴] 内
// -1000 <= Node.val <= 1000
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 设计 字符串 二叉树 👍 1037 👎 0

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

import (
	"strconv"
	"strings"
)

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := make([]string, 0)
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		newNodes := make([]*TreeNode, 0)
		for i := 0; i < len(nodes); i++ {
			if nodes[i] != nil {
				res = append(res, strconv.Itoa(nodes[i].Val))
				newNodes = append(newNodes, nodes[i].Left, nodes[i].Right)
			} else {
				res = append(res, "null")
			}
		}
		nodes = newNodes
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	datas := strings.Split(data, ",")
	if datas[0] == "null" {
		return nil
	}
	val, _ := strconv.Atoi(datas[0])
	root := &TreeNode{
		Val: val,
	}
	nodes := []*TreeNode{root}
	tag := 1
	for len(nodes) > 0 {
		newNode := make([]*TreeNode, 0)
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			if datas[tag] != "null" {
				val, _ := strconv.Atoi(datas[tag])
				node.Left = &TreeNode{Val: val}
				newNode = append(newNode, node.Left)
			}
			tag++
			if datas[tag] != "null" {
				val, _ := strconv.Atoi(datas[tag])
				node.Right = &TreeNode{Val: val}
				newNode = append(newNode, node.Right)
			}
			tag++
		}
		nodes = newNode
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
//leetcode submit region end(Prohibit modification and deletion)
