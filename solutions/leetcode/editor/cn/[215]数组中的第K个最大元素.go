package main

import "container/heap"

//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 示例 1: 
//
// 输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
// 
//
// 示例 2: 
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4 
//
// 说明: 
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。 
// Related Topics 堆 分治算法 
// 👍 979 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type IntHeap []int

func (h *IntHeap)Len()int  {
	return len((*h))
}
func (h *IntHeap)Less(i,j int)bool  {
	return (*h)[i]<(*h)[j]
}
func (h *IntHeap)Swap(i,j int){
	(*h)[i],(*h)[j] = (*h)[j],(*h)[i]
}
func (h *IntHeap)Push(e interface{}){
	*h = append(*h,e.(int))
}
func (h *IntHeap)Pop()interface{}  {
	ret := (*h)[len(*h)-1]
	*h = (*h)[0:len(*h)-1]
	return ret
}
func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for i,e:=range nums{
		heap.Push(h,e)
		if i >= k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

//leetcode submit region end(Prohibit modification and deletion)
