package main
//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。 
//
// 返回滑动窗口中的最大值。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
// 
//
// 示例 2： 
//
// 
//输入：nums = [1], k = 1
//输出：[1]
// 
//
// 示例 3： 
//
// 
//输入：nums = [1,-1], k = 1
//输出：[1,-1]
// 
//
// 示例 4： 
//
// 
//输入：nums = [9,11], k = 2
//输出：[11]
// 
//
// 示例 5： 
//
// 
//输入：nums = [4,-2], k = 2
//输出：[4] 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 105 
// -104 <= nums[i] <= 104 
// 1 <= k <= nums.length 
// 
// Related Topics 堆 Sliding Window 
// 👍 883 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

func maxSlidingWindow(nums []int, k int) []int {
	// 维护一个双端队列，滑动窗口内最大的元素始终是最左边的元素
	// 向右滑动时，如果push进了一个新的最大值，就把新最大值左端的元素全部出队
	// 向右滑动时，如果push进了一个非最大值，则把小于该值的元素都出队
	// 坑：这里不能直接使用nums放到数组里，而是要用num的index代替，这样是为了方便计算队列的长度

	ret := make([]int, 0)
	d := &deque{
		[]int{},
		0,
	}
	for i,e :=range nums{
		if i >= k && i - d.D[0] >= k {
			d.pop()
		}
		for len(d.D) >0 && nums[d.D[len(d.D)-1]]<e{
			d.popleft()
		}
		d.push(i)
		if i+1 >= k {
			ret = append(ret,nums[d.peek()])
		}
	}

	return ret
}

type deque struct {
	D []int
	l int
}

func (d *deque) pop() {
	d.D = d.D[1:]
}
func (d *deque) push(e int) {
	d.D = append(d.D, e)
}
func (d *deque) peek() int {
	return d.D[0]
}
func (d *deque) popleft() {
	d.D = d.D[:len(d.D)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
