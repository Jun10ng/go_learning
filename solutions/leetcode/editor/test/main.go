package main

import "fmt"

func main() {
	nums := []int{1, -1}
	fmt.Print(maxSlidingWindow(nums, 1))
	fmt.Print(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Print(maxSlidingWindow([]int{3, -2, -3, -1, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	// 维护一个双端队列，滑动窗口内最大的元素始终是最左边的元素
	// 向右滑动时，如果push进了一个新的最大值，就把新最大值左端的元素全部出队
	// 向右滑动时，如果push进了一个非最大值，则把小于该值的元素都出队
	ret := make([]int, 0)
	d := &deque{
		[]int{},
	}
	for i, e := range nums {
		if d.len() >= k {
			d.pop()
		}
		for d.len() > 0 && e > d.D[d.len()-1] {
			d.popTail(i)
		}
		d.push(e)
		ret = append(ret, d.max())
	}
	return ret
}

type deque struct {
	D []int
}

func (d *deque) len() int {
	return len(d.D)
}

func (d *deque) pop() {
	d.D = d.D[1:]
}
func (d *deque) push(e int) {
	d.D = append(d.D, e)
}
func (d *deque) max() int {
	return d.D[0]
}
func (d *deque) popTail(i int) {
	d.D = d.D[:len(d.D)-1]
	//if i == d.len() {
	//	d.D = []int{}
	//}
	//d.D =append(d.D[:i],d.D[i+1:]...)
}
