//基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
//
// 假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
//
//
// 例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
//
//
// 另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。（变化后的基因必须位于基因库 bank 中）
//
// 给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数。如果无法完成
//此基因变化，返回 -1 。
//
// 注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。
//
//
//
// 示例 1：
//
//
//输入：start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]
//输出：1
//
//
// 示例 2：
//
//
//输入：start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA",
//"AAACGGTA"]
//输出：2
//
//
// 示例 3：
//
//
//输入：start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC",
//"AACCCCCC"]
//输出：3
//
//
//
//
// 提示：
//
//
// start.length == 8
// end.length == 8
// 0 <= bank.length <= 10
// bank[i].length == 8
// start、end 和 bank[i] 仅由字符 ['A', 'C', 'G', 'T'] 组成
//
//
// Related Topics 广度优先搜索 哈希表 字符串 👍 233 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func minMutation(startGene string, endGene string, bank []string) int {
	all := make(map[string]bool, len(bank))
	for i := 0; i < len(bank); i++ {
		all[bank[i]] = true
	}
	if !all[endGene] {
		return -1
	}
	res := 0
	queue := []string{startGene}
	thisall := map[string]bool{startGene:true}
	for len(queue) > 0 {
		var nextQueue []string
		nextAll := make(map[string]bool)
		for i := 0; i < len(queue); i++ {
			tmp := queue[i]
			for b := 0; b < len(bank); b++ {
				for j := 0; j < len(startGene); j++ {
					if tmp[j] != bank[b][j] {
						next := tmp[:j] + string(bank[b][j]) + tmp[j+1:]
						if next == endGene {
							return res + 1
						}
						if !nextAll[next] &&  all[next] && !thisall[next]{
							nextQueue = append(nextQueue, next)
							nextAll[next] = true
						}
					}
				}
			}
		}
		res++
		queue = nextQueue
		thisall =nextAll
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
