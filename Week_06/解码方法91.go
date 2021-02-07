/**
  @author: dingfeng
  @date: 2021/2/6
  @note:
**/
package leetcode

import (
	"strings"
)

/*
动态规划
F[i] = F[i-1]+1
F[i] = F[i-2]+1
*/
func NumDecodings(s string) int {
	if strings.HasPrefix(s, "0") {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 1
	if len(s) >= 2 {
		if s[0] > 50 && s[1] == '0' {
			return 0
		}
		dp[1] = 1
		if s[1] != '0' {
			if s[0] == '1' || s[0] == '2' && s[1] <= 54 {
				dp[1] = 2
			}
		}
	}

	for i := 2; i < len(s); i++ {
		if s[i] == '0' && s[i-1] != '1' && s[i-1] != '2' {
			return 0
		}
		if s[i] != '0' {
			dp[i] = dp[i-1]
			if (s[i-1] == '2' && s[i] <= 54) || (s[i-1] == '1') {
				dp[i] += dp[i-2]
			}
		} else {
			dp[i] = dp[i-2]
		}
	}
	return dp[len(s)-1]

}
