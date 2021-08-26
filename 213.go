package main

import "fmt"

/*

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

示例 1：

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2：

输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func rob1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// 用第一个
	dp := make([]int, len(nums), len(nums))
	dp[0] = nums[0]
	dp[1] = nums[0]

	// 不用第一个
	dp1 := make([]int, len(nums), len(nums))
	dp1[0] = 0
	dp1[1] = nums[1]

	for i := 2; i < len(nums); i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
	}

	for i := 2; i < len(nums); i++ {
		if i == len(nums)-1 {
			dp[i] = max(dp[i-1], dp[i-2])
		} else {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}
	}

	return max(dp[len(nums)-1], dp1[len(nums)-1])
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}


