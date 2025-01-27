package leetcode

import (
	"fmt"
	"math"
	"time"
)

func SolveFindMinInRotated(nums []int) int {
	l := 0
	r := len(nums)
	min := math.MaxInt

	for {
		mid := (r - l) / 2
		mid += l
		fmt.Printf("leftPointer: %v, rightPointer: %v, midPointer: %v\n", l, r, mid)
		if nums[mid] > nums[l] {
			l = mid
		} else {
			r = mid
			min = nums[mid]
		}

		if l+1 == r || l == r {
			// If min was not set
			if min == math.MaxInt {
				min = nums[0]
			}
			return min
		}
		fmt.Println("---------")
		time.Sleep(2 * time.Second)
	}
}
