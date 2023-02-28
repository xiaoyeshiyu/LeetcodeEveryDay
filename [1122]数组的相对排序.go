//给你两个数组，arr1 和 arr2，arr2 中的元素各不相同，arr2 中的每个元素都出现在 arr1 中。
//
// 对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末
//尾。
//
//
//
// 示例 1：
//
//
//输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
//输出：[2,2,2,1,4,3,3,9,6,7,19]
//
//
// 示例 2:
//
//
//输入：arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
//输出：[22,28,8,6,17,44]
//
//
//
//
// 提示：
//
//
// 1 <= arr1.length, arr2.length <= 1000
// 0 <= arr1[i], arr2[i] <= 1000
// arr2 中的元素 arr2[i] 各不相同
// arr2 中的每个元素 arr2[i] 都出现在 arr1 中
//
//
// Related Topics 数组 哈希表 计数排序 排序 👍 253 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func relativeSortArray(arr1 []int, arr2 []int) []int {
	length2 := len(arr2)
	all2 := make(map[int]int, length2)
	for i := 0; i < length2; i++ {
		all2[arr2[i]] = i
	}

	length1 := len(arr1)
	all1 := make([][]int, length2+1)
	all1[length2] = make([]int, 0, length1)
	for i := 0; i < length1; i++ {
		point, ok := all2[arr1[i]]
		if ok {
			if all1[point] == nil {
				all1[point] = make([]int, 0, length1)
			}
			all1[point] = append(all1[point], arr1[i])
		} else {
			var inside int
			var flag  bool
			for j := 0; j < len(all1[length2]); j++ {
				if all1[length2][j] > arr1[i] {
					inside = j
					flag = true
					break
				}
			}
			all1[length2] = append(all1[length2], arr1[i])
			if flag {
				copy(all1[length2][inside+1:], all1[length2][inside:])
				all1[length2][inside] = arr1[i]
			}
		}
	}

	res := make([]int, 0, length1)
	for i := 0; i < length2+1; i++ {
		tmp := all1[i]
		for j := 0; j < len(tmp); j++ {
			res = append(res, tmp[j])
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
