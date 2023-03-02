// 二进制数转字符串。给定一个介于0和1之间的实数（如0.72），类型为double，打印它的二进制表达式。如果该数字无法精确地用32位以内的二进制表示，则打印
// “ERROR”。
//
// 示例1:
//
// 输入：0.625
// 输出："0.101"
//
// 示例2:
//
// 输入：0.1
// 输出："ERROR"
// 提示：0.1无法被二进制准确表示
//
// 提示：
//
// 32位包括输出中的 "0." 这两位。
// 题目保证输入用例的小数位数最多只有 6 位
//
// Related Topics 位运算 数学 字符串 👍 63 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func printBin(num float64) string {
	s := strings.Builder{}
	s.WriteString("0.")
	var size int

	for size < 32 && num != 0 {
		num *= 2
		s.WriteString(strconv.Itoa(int(num)))
		num = num - float64(int(num))
		size++
	}
	if size < 32 {
		return s.String()
	}
	return "ERROR"
}

//leetcode submit region end(Prohibit modification and deletion)
