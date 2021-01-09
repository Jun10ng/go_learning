package interview

/*
	划分一个整形数组，把正数放在左边，零放在中间，负数放在右边
    -1 2 0 -3 0 4 -5
*/

func moveNum(nums []int) []int {
	posPtr, zeroPtr, negPtr := 0, 0, len(nums)-1
	for zeroPtr < negPtr {
		if nums[zeroPtr] > 0 {
			nums[zeroPtr], nums[posPtr] = nums[posPtr], nums[zeroPtr]
			zeroPtr++
		}
		if nums[zeroPtr] == 0 {
			zeroPtr++
		}
		if nums[zeroPtr] < 0 {
			nums[zeroPtr], nums[negPtr] = nums[negPtr], nums[zeroPtr]
			negPtr--
		}
	}
	return nums
}
