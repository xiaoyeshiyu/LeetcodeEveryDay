//按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，一个表示此过程的 转换序列 是形式上像 beginWord ->
//s1 -> s2 -> ... -> sk 这样的单词序列，并满足：
//
//
//
//
// 每对相邻的单词之间仅有单个字母不同。
// 转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList 中的单
//词。
// sk == endWord
//
//
//
//
// 给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。请你找出并返回所有从 beginWord 到 endWord 的
// 最短转换序列 ，如果不存在这样的转换序列，返回一个空列表。每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk] 的形式返回。
//
//
//
// 示例 1：
//
//
//输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot",
//"log","cog"]
//输出：[["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
//解释：存在 2 种最短的转换序列：
//"hit" -> "hot" -> "dot" -> "dog" -> "cog"
//"hit" -> "hot" -> "lot" -> "log" -> "cog"
//
//
// 示例 2：
//
//
//输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot",
//"log"]
//输出：[]
//解释：endWord "cog" 不在字典 wordList 中，所以不存在符合要求的转换序列。
//
//
//
//
// 提示：
//
//
// 1 <= beginWord.length <= 5
// endWord.length == beginWord.length
// 1 <= wordList.length <= 500
// wordList[i].length == beginWord.length
// beginWord、endWord 和 wordList[i] 由小写英文字母组成
// beginWord != endWord
// wordList 中的所有单词 互不相同
//
//
// Related Topics 广度优先搜索 哈希表 字符串 回溯 👍 629 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
type ThisNode struct {
	Val      string
	Children []*ThisNode
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	length := 20
	all := make(map[string]bool, length)
	for i := 0; i < len(wordList); i++ {
		all[wordList[i]] = true
	}
	if !all[endWord] {
		return nil
	}
	root := &ThisNode{
		Val:      beginWord,
		Children: make([]*ThisNode, 0,length),
	}
	queue := []*ThisNode{root}
	level := 0
	find := false
	for len(queue) > 0 {
		newQueue := make([]*ThisNode, 0,length)
		for i := 0; i < len(queue); i++ {
			tmp := queue[i]
			for j := range all {
				tag := 0
				for k := 0; k < len(tmp.Val); k++ {
					if j[k] != tmp.Val[k] {
						tag++
					}
					if tag == 2 {
						break
					}
				}
				if tag == 1 {
					myNode := &ThisNode{
						Val:      j,
						Children: make([]*ThisNode, 0,length),
					}
					tmp.Children = append(tmp.Children, myNode)
					if j == endWord {
						find = true
					}
					newQueue = append(newQueue, myNode)
				}
			}
		}
		if find {
			break
		}
		for i := 0; i < len(newQueue); i++ {
			delete(all, newQueue[i].Val)
		}
		queue = newQueue
		level++
		fmt.Println(level)
	}
	if !find {
		return nil
	}
	fmt.Println("get node")
	return DFS(root, endWord)
}

func DFS(root *ThisNode, endWord string) [][]string {
	var res [][]string
	if root != nil {
		if root.Val == endWord {
			return [][]string{{endWord}}
		}
		for i := 0; i < len(root.Children); i++ {
			child := DFS(root.Children[i], endWord)
			for c := 0; c < len(child); c++ {
				path := make([]string, len(child[c])+1)
				path[0] = root.Val
				copy(path[1:], child[c])
				res = append(res, path)
			}
		}
		return res
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
