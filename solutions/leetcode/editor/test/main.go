package main

import "fmt"

func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	start, end := 0, len(numbers)-1
	mid := (start + end) / 2
	for start < end {
		mid = (start + end) / 2
		if numbers[mid] < numbers[end] {
			//最小数在前半段
			end = mid
		} else if numbers[mid] > numbers[end] {
			//最小数在后半段
			start = mid + 1
		} else {
			end--
		}
	}
	return numbers[start]
}

func main() {
	n := []int{4, 1, 2, 3}
	fmt.Print(minArray(n))
}
