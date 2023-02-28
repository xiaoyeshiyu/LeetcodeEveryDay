//给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。
//
// 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使
//用。
//
//
//
// 示例 1：
//
//
//输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f",
//"l","v"]], words = ["oath","pea","eat","rain"]
//输出：["eat","oath"]
//
//
// 示例 2：
//
//
//输入：board = [["a","b"],["c","d"]], words = ["abcb"]
//输出：[]
//
//
//
//
// 提示：
//
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 12
// board[i][j] 是一个小写英文字母
// 1 <= words.length <= 3 * 10⁴
// 1 <= words[i].length <= 10
// words[i] 由小写英文字母组成
// words 中的所有字符串互不相同
//
//
// Related Topics 字典树 数组 字符串 回溯 矩阵 👍 737 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)

const Letter = 26

var res []string
var m, n int

type Tire struct {
	word  string
	child [Letter]*Tire
}

func (t *Tire) insert(word string) {
	node := t
	for _, w := range word {
		w = w - 'a'
		if node.child[w] == nil {
			node.child[w] = new(Tire)
		}
		node = node.child[w]
	}
	node.word = word
}

func findWords(board [][]byte, words []string) []string {
	tire := new(Tire)
	for _, word := range words {
		tire.insert(word)
	}

	m = len(board)
	n = len(board[0])
	res = make([]string, 0, len(words))

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			word := board[i][j] - 'a'
			if tire.child[word] != nil {
				getLatter(i, j, board, tire)
			}
		}
	}

	return res
}

func getLatter(i int, j int, board [][]byte, tire *Tire) {
	if tire.word != "" {
		res = append(res, tire.word)
		tire.word = ""
	}

	if i < 0 || i > m-1 || j < 0 || j > n-1 || board[i][j] == '0' {
		return
	}


	word := board[i][j] - 'a'
	if tire.child[word] != nil {
		board[i][j] = '0'
		getLatter(i-1, j, board, tire.child[word])
		getLatter(i, j+1, board, tire.child[word])
		getLatter(i+1, j, board, tire.child[word])
		getLatter(i, j-1, board, tire.child[word])
		board[i][j] = word + 'a'
	}
}

//leetcode submit region end(Prohibit modification and deletion)
