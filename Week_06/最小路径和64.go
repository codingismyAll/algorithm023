/**
  @author: dingfeng
  @date: 2021/2/5
  @note:
**/
package leetcode

/*
动态规划
状态转移方程
f(x,y) = min(f(x-1,y),f(x,y-1)+grid(x,y)
复杂度分析

时间复杂度：O(mn)，其中 m 和 n 分别是网格的行数和列数。需要对整个网格遍历一次，计算dp 的每个元素的值。

空间复杂度：O(mn)，其中 m 和 n 分别是网格的行数和列数。创建一个二维数组 dp，和网格大小相同。
空间复杂度可以优化，例如每次只存储上一行的}dp 值，则可以将空间复杂度优化到 O(n)。

*/
func MinPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := grid
	for i := 1; i < len(grid); i++ {
		dp[i][0] += dp[i-1][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] += dp[0][i-1]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[n-1][m-1]
}
