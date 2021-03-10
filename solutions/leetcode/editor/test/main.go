package main

import "fmt"

func main() {
	k := 2
	nums := []int{0}
	vals := []int{-1, 1, -2, -4, 3}
	obj := Constructor(k, nums)
	for _, val := range vals {
		fmt.Print(obj.Add(val))
	}
}

type KthLargest struct {
	K   int
	Kth []int
}

func Constructor(k int, nums []int) KthLargest {
	this := KthLargest{
		k,
		[]int{},
	}
	for _, num := range nums {
		this.Add(num)
	}
	return this
}

func (this *KthLargest) Add(val int) int {
	this.Insert(val)
	if len(this.Kth) >= this.K {
		return this.Kth[len(this.Kth)-1]
	}
	return this.Kth[len(this.Kth)-1]
}

func (this *KthLargest) Insert(val int) {
	if len(this.Kth) == 0 {
		this.Kth = append(this.Kth, val)
		return
	}
	if this.Kth[len(this.Kth)-1] > val {
		if len(this.Kth) < this.K {
			this.Kth = append(this.Kth, val)
		}
		return
	}
	i, e := 0, 0
	for i, e = range this.Kth {
		if e <= val {
			break
		}
	}
	new := make([]int, 0, len(this.Kth)+1)
	new = append(new, this.Kth[:i]...)
	new = append(new, val)
	new = append(new, this.Kth[i:]...)
	if len(new) > this.K {
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
