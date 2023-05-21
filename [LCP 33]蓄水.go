//给定 N 个无限容量且初始均空的水缸，每个水缸配有一个水桶用来打水，第 `i` 个水缸配备的水桶容量记作 `bucket[i]`。小扣有以下两种操作：
//- 升级水桶：选择任意一个水桶，使其容量增加为 `bucket[i]+1`
//- 蓄水：将全部水桶接满水，倒入各自对应的水缸
//
//每个水缸对应最低蓄水量记作 `vat[i]`，返回小扣至少需要多少次操作可以完成所有水缸蓄水要求。
//
//注意：实际蓄水量 **达到或超过** 最低蓄水量，即完成蓄水要求。
//
//**示例 1：**
//
//> 输入：`bucket = [1,3], vat = [6,8]`
//>
//> 输出：`4`
//>
//> 解释：
//> 第 1 次操作升级 bucket[0]；
//> 第 2 ~ 4 次操作均选择蓄水，即可完成蓄水要求。
//> ![vat1.gif](https://pic.leetcode-cn.com/1616122992-RkDxoL-vat1.gif)
//
//**示例 2：**
//
//> 输入：`bucket = [9,0,1], vat = [0,2,2]`
//>
//> 输出：`3`
//>
//> 解释：
//> 第 1 次操作均选择升级 bucket[1]
//> 第 2~3 次操作选择蓄水，即可完成蓄水要求。
//
//**提示：**
//- `1 <= bucket.length == vat.length <= 100`
//- `0 <= bucket[i], vat[i] <= 10^4`
//
// Related Topics 贪心 数组 堆（优先队列） 👍 180 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func storeWater(bucket []int, vat []int) int {
	n := len(bucket)
	maxk := 0
	for _, v := range vat {
		if v > maxk {
			maxk = v
		}
	}
	if maxk == 0 {
		return 0
	}
	res := math.MaxInt32
	for k := 1; k <= maxk && k < res; k++ {
		t := 0
		for i := 0; i < n; i++ {
			t += max(0, (vat[i] + k - 1) / k - bucket[i])
		}
		res = min(res, t+k)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
