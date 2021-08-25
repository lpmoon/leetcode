package main

import (
	"fmt"
)

/*
示例 1：

输入：s = "12"
输出：2
解释：它可以解码为 "AB"（1 2）或者 "L"（12）。
示例 2：

输入：s = "226"
输出：3
解释：它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
示例 3：

输入：s = "0"
输出：0
解释：没有字符映射到以 0 开头的数字。
含有 0 的有效映射是 'J' -> "10" 和 'T'-> "20" 。
由于没有字符，因此没有有效的方法对此进行解码，因为所有数字都需要映射。
示例 4：

输入：s = "06"
输出：0
解释："06" 不能映射到 "F" ，因为字符串含有前导 0（"6" 和 "06" 在映射中并不等价）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-ways
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func numDecodings(s string) int {
	bytes := []byte(s)

	if len(s) == 1 {
		if bytes[0] == '0' {
			return 0
		}

		return 1
	}

	if len(s) == 2 {
		if bytes[0] == '0' {
			return 0
		}

		if bytes[0] >= '2' && bytes[1] >= '7' || bytes[0] >= '3' {
			if bytes[1] == '0' {
				return 0
			}
			return 1
		}

		if bytes[1] == '0' {
			return 1
		}

		return 2
	}

	dp := make([]int, len(s), len(s))

	if bytes[0] == '0' {
		return 0
	}

	dp[0] = 1

	if bytes[0] >= '2' && bytes[1] >= '7' || bytes[0] >= '3' {
		dp[1] = 1
		if bytes[1] == '0' {
			return 0
		}
	} else {
		if bytes[1] == '0' {
			dp[1] = 1
		} else {
			dp[1] = 2
		}
	}

	for i := 2; i < len(s); i++ {
		if bytes[i-1] == '0' && bytes[i] == '0' {
			return 0
		}

		if bytes[i-1] == '0' {
			dp[i] = dp[i-1]
		} else {
			if bytes[i-1] >= '2' && bytes[i] >= '7' || bytes[i-1] >= '3' {
				if bytes[i] == '0' {
					return 0
				}
				dp[i] = dp[i-1]
			} else {
				if bytes[i] == '0' {
					dp[i] = dp[i-2]
				} else {
					dp[i] = dp[i-1] + dp[i-2]
				}
			}
		}
	}

	return dp[len(s)-1]
}

func main() {
	fmt.Println(numDecodings(""))
}
