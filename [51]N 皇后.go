//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
// Related Topics 数组 回溯 👍 1598 👎 0
//package main

//import (
//	"fmt"
//	"strings"
//)
//
//var res [][]string
//
////leetcode submit region begin(Prohibit modification and deletion)
//func solveNQueens(n int) [][]string {
//	res = [][]string{}
//	board := clean(n)
//	for i := 0; i < n; i++ {
//		recursion(board, 0, i, 0, n)
//	}
//	return res
//}
//
//// 使用递归，将某一个格子放进去，然后递归判断是否合法，当合法，并且棋子个数满足，则记录
//func recursion(board [][]string, x, y, size, total int) {
//	//	退出条件，为该点不满足
//	if !isValid(board, x, y, total) {
//		return
//	}
//	board[x][y] = "Q"
//	size++
//	if size == total {
//		res = append(res, print(board))
//	}
//	for i := 0; i < total; i++ {
//		recursion(board, x+1, i, size, total)
//	}
//	board[x][y] = "."
//}
//
//func isValid(board [][]string, x, y, total int) bool {
//	fmt.Println(x, y, total, board)
//	//	退出条件，为该点不满足
//	var direction = [][]int{
//		{0, 1},
//		{1, 1},
//		{1, 0},
//		{1, -1},
//		{0, -1},
//		{-1, -1},
//		{-1, 0},
//		{-1, -1},
//	}
//	if x < 0 || x >= total || y < 0 || y >= total {
//		return false
//	}
//
//	for i := 0; i < 8; i++ {
//		tmpx, tmpy := x, y
//		for tmpx >= 0 && tmpx < total && tmpy >= 0 && tmpy < total {
//			if board[tmpx][tmpy] == "Q" {
//				fmt.Println("false")
//				return false
//			}
//			tmpx += direction[i][0]
//			tmpy += direction[i][1]
//		}
//	}
//	fmt.Println("true")
//	return true
//}
//
//func clean(n int) [][]string {
//	start := make([][]string, n)
//	for i := 0; i < n; i++ {
//		inside := make([]string, n)
//		for j := 0; j < n; j++ {
//			inside[j] = "."
//		}
//		start[i] = inside
//	}
//	return start
//}
//
//func print(board [][]string) []string {
//	var res []string
//	for i := 0; i < len(board); i++ {
//		res = append(res, strings.Join(board[i], ""))
//	}
//	fmt.Println(res)
//	return res
//}

import "strings"

var res [][]string

//func isValid(board [][]string, row, col int) (res bool) {
//	n := len(board)
//	for i := 0; i < row; i++ {
//		if board[i][col] == "Q" {
//			return false
//		}
//	}
//	for i := 0; i < n; i++ {
//		if board[row][i] == "Q" {
//			return false
//		}
//	}
//
//	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
//		if board[i][j] == "Q" {
//			return false
//		}
//	}
//	for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
//		if board[i][j] == "Q" {
//			return false
//		}
//	}
//	return true
//}

func isValid(board [][]string, x, y, total int) bool {
	//	退出条件，为该点不满足
	var direction = [][]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	if x < 0 || x >= total || y < 0 || y >= total {
		return false
	}

	for i := 0; i < 8; i++ {
		tmpx, tmpy := x, y
		for tmpx >= 0 && tmpx < total && tmpy >= 0 && tmpy < total {
			if board[tmpx][tmpy] == "Q" {
				return false
			}
			tmpx += direction[i][0]
			tmpy += direction[i][1]
		}
	}
	return true
}

//func backtrack(board [][]string, row int) {
//	size := len(board)
//	if row == size {
//		temp := make([]string, size)
//		for i := 0; i < size; i++ {
//			temp[i] = strings.Join(board[i], "")
//		}
//		res = append(res, temp)
//		return
//	}
//	for col := 0; col < size; col++ {
//		if !isValid(board, row, col) {
//			continue
//		}
//		board[row][col] = "Q"
//		backtrack(board, row+1)
//		board[row][col] = "."
//	}
//}

func solveNQueens(n int) [][]string {
	res = [][]string{}
	board := clean(n)
	for i := 0; i < n; i++ {
		recursion(board, 0, i, 0, n)
	}
	return res
}

func recursion(board [][]string, x, y, size, total int) {
	//	退出条件，为该点不满足
	if !isValid(board, x, y, total) {
		return
	}
	board[x][y] = "Q"
	size++
	if size == total {
		res = append(res, print(board))
	}
	for i := 0; i < total; i++ {
		recursion(board, x+1, i, size, total)
	}
	board[x][y] = "."
}

func clean(n int) [][]string {
	start := make([][]string, n)
	for i := 0; i < n; i++ {
		inside := make([]string, n)
		for j := 0; j < n; j++ {
			inside[j] = "."
		}
		start[i] = inside
	}
	return start
}

func print(board [][]string) []string {
	var res []string
	for i := 0; i < len(board); i++ {
		res = append(res, strings.Join(board[i], ""))
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
