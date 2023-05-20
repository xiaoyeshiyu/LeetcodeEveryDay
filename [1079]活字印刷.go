// 你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。
//
// 注意：本题中，每个活字字模只能使用一次。
//
// 示例 1：
//
// 输入："AAB"
// 输出：8
// 解释：可能的序列为 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"。
//
// 示例 2：
//
// 输入："AAABBC"
// 输出：188
//
// 示例 3：
//
// 输入："V"
// 输出：1
//
// 提示：
//
// 1 <= tiles.length <= 7
// tiles 由大写英文字母组成
//
// Related Topics 哈希表 字符串 回溯 计数 👍 205 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func numTilePossibilities(tiles string) int {
	allArray := []string{""}
	all := make(map[string]bool)
	for i := 0; i < len(tiles); i++ {
		for _,tmp := range allArray {
			// 中间
			for j := 0; j < len(tmp); j++ {
				newS := tmp[:j]+tiles[i:i+1]+tmp[j:]
				if !all[newS] {
					allArray = append(allArray, newS)
					all[newS] = true
				}
			}
			// 尾
			newS := tmp+tiles[i:i+1]
			if !all[newS] {
				allArray = append(allArray, newS)
				all[newS] = true
			}
		}
	}
	return len(all)
}

//leetcode submit region end(Prohibit modification and deletion)
