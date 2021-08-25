package leetcode

// https://leetcode-cn.com/problems/house-robber/

func rob(nums []int) int {
	//输入：[2,7,9,3,1]
	//输出：12
	//解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
	//偷窃到的最高金额 = 2 + 9 + 1 = 12 。

	// dp[n] = dp[n-1] + (dp[n-2] + nums[n])
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums), len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
