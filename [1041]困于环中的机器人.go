// 在无限的平面上，机器人最初位于 (0, 0) 处，面朝北方。注意:
//
// 北方向 是y轴的正方向。
// 南方向 是y轴的负方向。
// 东方向 是x轴的正方向。
// 西方向 是x轴的负方向。
//
// 机器人可以接受下列三条指令之一：
//
// "G"：直走 1 个单位
// "L"：左转 90 度
// "R"：右转 90 度
//
// 机器人按顺序执行指令 instructions，并一直重复它们。
//
// 只有在平面中存在环使得机器人永远无法离开时，返回 true。否则，返回 false。
//
// 示例 1：
//
// 输入：instructions = "GGLLGG"
// 输出：true
// 解释：机器人最初在(0,0)处，面向北方。
// “G”:移动一步。位置:(0,1)方向:北。
// “G”:移动一步。位置:(0,2).方向:北。
// “L”:逆时针旋转90度。位置:(0,2).方向:西。
// “L”:逆时针旋转90度。位置:(0,2)方向:南。
// “G”:移动一步。位置:(0,1)方向:南。
// “G”:移动一步。位置:(0,0)方向:南。
// 重复指令，机器人进入循环:(0,0)——>(0,1)——>(0,2)——>(0,1)——>(0,0)。
// 在此基础上，我们返回true。
//
// 示例 2：
//
// 输入：instructions = "GG"
// 输出：false
// 解释：机器人最初在(0,0)处，面向北方。
// “G”:移动一步。位置:(0,1)方向:北。
// “G”:移动一步。位置:(0,2).方向:北。
// 重复这些指示，继续朝北前进，不会进入循环。
// 在此基础上，返回false。
//
// 示例 3：
//
// 输入：instructions = "GL"
// 输出：true
// 解释：机器人最初在(0,0)处，面向北方。
// “G”:移动一步。位置:(0,1)方向:北。
// “L”:逆时针旋转90度。位置:(0,1).方向:西。
// “G”:移动一步。位置:(- 1,1)方向:西。
// “L”:逆时针旋转90度。位置:(- 1,1)方向:南。
// “G”:移动一步。位置:(- 1,0)方向:南。
// “L”:逆时针旋转90度。位置:(- 1,0)方向:东方。
// “G”:移动一步。位置:(0,0)方向:东方。
// “L”:逆时针旋转90度。位置:(0,0)方向:北。
// 重复指令，机器人进入循环:(0,0)——>(0,1)——>(- 1,1)——>(- 1,0)——>(0,0)。
// 在此基础上，我们返回true。
//
// 提示：
//
// 1 <= instructions.length <= 100
// instructions[i] 仅包含 'G', 'L', 'R'
//
// Related Topics 数学 字符串 模拟 👍 160 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func isRobotBounded(instructions string) bool {
	pos := [2]int{}
	dir := [2]int{0, 1}

	for i := 0; i < 4; i++ {

		for _, instruction := range instructions {
			switch instruction {
			case 'G':
				pos[0] += dir[0]
				pos[1] += dir[1]
			case 'L':
				pos[0], pos[1] = pos[1]*-1, pos[0]
			case 'R':
				pos[0], pos[1] = pos[1], pos[0]*-1
			}
		}
		if pos == [2]int{0, 0} && dir == [2]int{0, 1} {
			return true
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
