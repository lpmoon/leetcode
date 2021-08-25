package main

// https://leetcode-cn.com/problems/count-numbers-with-unique-digits/

// 3 739

func CountNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 10
	}

	preC := 10
	for i := 1; i < n-1; i++ {
		c := 9
		starCount := 9
		for j := 0; j < i; j++ {
			c *= starCount
			starCount--
		}

		preC += c
	}

	t := 9
	starCount := 9
	for i := 0; i < n-1; i++ {
		t *= starCount
		starCount--
	}

	return preC + t
}
