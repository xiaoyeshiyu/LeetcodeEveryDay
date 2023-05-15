//在一个仓库里，有一排条形码，其中第 i 个条形码为 barcodes[i]。 
//
// 请你重新排列这些条形码，使其中任意两个相邻的条形码不能相等。 你可以返回任何满足该要求的答案，此题保证存在答案。 
//
// 
//
// 示例 1： 
//
// 
//输入：barcodes = [1,1,1,2,2,2]
//输出：[2,1,2,1,2,1]
// 
//
// 示例 2： 
//
// 
//输入：barcodes = [1,1,1,1,2,2,3,3]
//输出：[1,3,1,3,2,1,2,1] 
//
// 
//
// 提示： 
//
// 
// 1 <= barcodes.length <= 10000 
// 1 <= barcodes[i] <= 10000 
// 
//
// Related Topics 贪心 数组 哈希表 计数 排序 堆（优先队列） 👍 171 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
type PriorityQueue [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][0] > pq[j][0]
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.([]int)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func rearrangeBarcodes(barcodes []int) []int {
	count := make(map[int]int)
	for _, b := range barcodes {
		count[b]++
	}
	q := &PriorityQueue{}
	heap.Init(q)
	for k, v := range count {
		heap.Push(q, []int{v, k})
	}
	n := len(barcodes)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		p := heap.Pop(q).([]int)
		cx, x := p[0], p[1]
		if i == 0 || res[i-1] != x {
			res[i] = x
			if cx > 1 {
				heap.Push(q, []int{cx - 1, x})
			}
		} else {
			p2 := heap.Pop(q).([]int)
			cy, y := p2[0], p2[1]
			res[i] = y
			if cy > 1 {
				heap.Push(q, []int{cy - 1, y})
			}
			heap.Push(q, []int{cx, x})
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
