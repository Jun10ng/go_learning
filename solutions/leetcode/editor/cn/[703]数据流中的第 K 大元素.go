package main
//设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。 
//
// 请实现 KthLargest 类： 
//
// 
// KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。 
// int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。 
// 
//
// 
//
// 示例： 
//
// 
//输入：
//["KthLargest", "add", "add", "add", "add", "add"]
//[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
//输出：
//[null, 4, 5, 5, 8, 8]
//
//解释：
//KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
//kthLargest.add(3);   // return 4
//kthLargest.add(5);   // return 5
//kthLargest.add(10);  // return 5
//kthLargest.add(9);   // return 8
//kthLargest.add(4);   // return 8
// 
//
// 
//提示：
//
// 
// 1 <= k <= 104 
// 0 <= nums.length <= 104 
// -104 <= nums[i] <= 104 
// -104 <= val <= 104 
// 最多调用 add 方法 104 次 
// 题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素 
// 
// Related Topics 堆 设计 
// 👍 250 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type KthLargest struct {
	K int
	Kth []int
}


func Constructor(k int, nums []int) KthLargest {
	this := KthLargest{
		k,
		[]int{},
	}
	for _,num :=range nums{
		this.Add(num)
	}
	return this
}


func (this *KthLargest) Add(val int) int {
	this.Insert(val)
	if len(this.Kth) >= this.K {
		return this.Kth[len(this.Kth)-1]
	}
	return 0
}

func (this *KthLargest) Insert(val int) {
	if len(this.Kth)==0 {
		this.Kth = append(this.Kth,val)
		return
	}
	if this.Kth[len(this.Kth)-1] > val{
		if len(this.Kth) < this.K {
			this.Kth = append(this.Kth,val)
		}
		return
	}
	i,e := 0,0
	for i,e = range this.Kth {
		if e <= val {
			break
		}
	}
	new := make([]int,0,len(this.Kth)+1)
	new = append(new,this.Kth[:i]...)
	new = append(new,val)
	new = append(new,this.Kth[i:]...)
	if len(new)>this.K {
		new = new[:len(new)-1]
	}
	this.Kth = new
	return
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
//leetcode submit region end(Prohibit modification and deletion)
