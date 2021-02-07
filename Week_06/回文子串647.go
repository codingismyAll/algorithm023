/**
  @author: dingfeng
  @date: 2021/2/6
  @note:
**/
package leetcode

/*
动态规划
dp[i][j]  表示s[i,j]的子串是不是回文字符串
*/
func CountSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	count:=0
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
				count++
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
				count++
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				count++
			}
		}
	}
	return count
}
