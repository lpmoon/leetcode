package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	// dp[i]以i作为最后一个元素的子序列最大值
	dp := nums[0]
	max := dp
	for i := 1; i < len(nums); i++ {
		if dp <= 0 {
			dp = nums[i]
		} else {
			dp += nums[i]
		}

		if max <= dp {
			max = dp
		}
	}

	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
